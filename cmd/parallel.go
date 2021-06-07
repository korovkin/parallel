package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/daviddengcn/go-colortext"
	"github.com/korovkin/limiter"
	"github.com/korovkin/parallel"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type logger struct {
	ticket   int
	hostname string
	isError  bool
	buf      *bytes.Buffer
	print    bool
}

var (
	loggerMutex     = new(sync.Mutex)
	loggerIndex     = int(0)
	loggerStartTime = time.Now()
	loggerHostname  = ""

	flag_verbose = flag.Bool(
		"v",
		false,
		"verbose level: 1")
)

var loggerColors = []ct.Color{
	ct.Green,
	ct.Cyan,
	ct.Magenta,
	ct.Yellow,
	ct.Blue,
	ct.Red,
}

func (l *logger) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	wrote := 0
	for {
		line, err := buf.ReadBytes('\n')
		if len(line) > 1 {
			now := time.Now().Format("15:01:02")
			s := string(line)
			ts := time.Since(loggerStartTime).String()
			e := "I"
			if l.isError {
				e = "E"
			}

			{
				loggerMutex.Lock()
				if l.print {
					ct.ChangeColor(loggerColors[l.ticket%len(loggerColors)], false, ct.None, false)
					fmt.Printf("[%-14s %s %s %03d %s] ", ts, l.hostname, now, l.ticket, e)
					ct.ResetColor()
					fmt.Print(s)
				}
				if l.buf != nil {
					l.buf.Write([]byte(s))
				}
				loggerMutex.Unlock()
			}

			wrote += len(line)
		}
		if err != nil {
			break
		}
	}
	if len(p) > 0 && p[len(p)-1] != '\n' {
		fmt.Println()
	}

	return len(p), nil
}

func newLogger(ticket int, collectLines bool) *logger {
	l := &logger{ticket: ticket, buf: nil}
	if collectLines {
		l.buf = &bytes.Buffer{}
	}
	l.print = true
	return l
}

func CheckFatal(e error) error {
	if e != nil {
		debug.PrintStack()
		log.Println("CHECK: ERROR:", e)
		panic(e)
	}
	return e
}

func CheckNotFatal(e error) error {
	if e != nil {
		debug.PrintStack()
		log.Println("CHECK: ERROR:", e, e.Error())
	}
	return e
}

func executeCommand(p *Parallel, ticket int, cmdLine string) (*parallel.Output, error) {
	p.StatNumCommandsStart.Inc()
	T_START := time.Now()
	var err error
	output := &parallel.Output{}
	loggerOut := newLogger(ticket, true)
	loggerOut.isError = false
	loggerErr := newLogger(ticket, true)
	loggerErr.isError = true

	defer func() {
		dt := time.Since(T_START)
		fmt.Fprintf(
			loggerOut,
			"execute: done: dt: "+dt.String()+"\n",
		)
		if err == nil {
			p.StatNumCommandsDone.Inc()
			p.StatCommandLatency.Observe(dt.Seconds())
		}
	}()

	// execute remotely:
	if len(p.Slaves) > 0 {
		slave := p.Slaves[ticket%len(p.Slaves)]

		var transport thrift.TTransport
		transport, err = thrift.NewTSocket(slave.Address)
		if err != nil {
			log.Fatalln("failed to dial slave:", err.Error())
		}

		transport, err = p.transportFactory.GetTransport(transport)
		if err != nil {
			log.Fatalln("failed to GetTransport:", err.Error())
		}

		err = transport.Open()
		if err != nil {
			log.Fatalln("failed to open:", err.Error())
		}

		defer transport.Close()
		client := parallel.NewParallelClientFactory(transport, p.protocolFactory)

		cmd := parallel.Cmd{
			CmdLine: cmdLine,
			Ticket:  int64(ticket),
		}

		loggerOut.hostname = slave.Address
		fmt.Fprintf(loggerOut, "start: '"+cmdLine+"'\n")

		output, err = client.Execute(context.Background(), &cmd)
		if err != nil {
			log.Fatalln("failed to execute:", err.Error())
		}

		hostname := output.Tags["hostname"]
		loggerOut.hostname = hostname
		loggerErr.hostname = hostname

		if *flag_verbose {
			fmt.Fprintf(loggerOut,
				"execute: remote: host: %s stdout: [%s]\n",
				hostname,
				output.Stdout)

			fmt.Fprintf(loggerErr,
				"execute: remote: host: %s stderr: [%s]\n",
				hostname,
				output.Stderr)
		}

		return output, err
	}

	// execute locally:
	cs := []string{"/bin/sh", "-c", cmdLine}
	cmd := exec.Command(cs[0], cs[1:]...)
	cmd.Stdin = nil
	cmd.Stdout = loggerOut
	cmd.Stderr = loggerErr
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("PARALLEL_TICKER=%d", ticket),
	)

	fmt.Fprintf(loggerOut, "start: '"+cmdLine+"'\n")

	loggerOut.print = *flag_verbose
	loggerErr.print = *flag_verbose

	err = cmd.Start()
	if err != nil {
		log.Fatalln("failed to start:", err)
		return output, err
	}

	if err == nil {
		err = cmd.Wait()
	}

	loggerOut.print = true
	loggerErr.print = true

	output.Tags = map[string]string{"hostname": loggerHostname}

	if loggerOut.buf != nil {
		output.Stdout = string(loggerOut.buf.Bytes())
	}

	if loggerErr.buf != nil {
		output.Stderr = string(loggerErr.buf.Bytes())
	}

	return output, err
}

type Slave struct {
	Address string
}

type Parallel struct {
	jobs   int
	logger *logger
	worker *limiter.ConcurrencyLimiter

	// master / slave
	protocolFactory  thrift.TProtocolFactory
	transportFactory thrift.TTransportFactory

	// master:
	Slaves []*Slave

	// slave:
	handler         *ParallelSlaveHandler
	serverTransport thrift.TServerTransport
	slaveAddress    string

	// stats:
	StatNumCommandsStart prometheus.Counter
	StatNumCommandsDone  prometheus.Counter
	StatCommandLatency   prometheus.Summary
}

func (p *Parallel) Close() {
	p.worker.Wait()
}

func mainMaster(p *Parallel) {
	var err error
	log.SetFlags(log.Lmicroseconds | log.Ldate | log.Lshortfile)

	// connect to slaves:
	for _, slave := range p.Slaves {
		var transport thrift.TTransport
		transport, err = thrift.NewTSocket(slave.Address)
		if err != nil {
			log.Fatalln("failed to dial slave:", err.Error())
		}

		transport, err = p.transportFactory.GetTransport(transport)
		if err != nil {
			log.Fatalln("failed to open:", err.Error())
		}

		err = transport.Open()
		if err != nil {
			log.Fatalln("failed to open:", err.Error())
		}

		defer transport.Close()

		client := parallel.NewParallelClientFactory(transport, p.protocolFactory)
		ok, err := client.Ping(context.Background())
		if err != nil {
			log.Fatalln("failed to ping client:", err.Error())
		}

		fmt.Fprintf(
			p.logger,
			"adding slave: %s ok: %s",
			slave.Address,
			ok)
	}

	r := bufio.NewReaderSize(os.Stdin, 1*1024*1024)
	fmt.Fprintf(p.logger, "reading from stdin...\n")
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		p.worker.ExecuteWithTicket(func(ticket int) {
			executeCommand(p, ticket, line)
		})
	}
}

type ParallelSlaveHandler struct {
	p *Parallel
}

func NewParallelSlaveHandler() *ParallelSlaveHandler {
	return &ParallelSlaveHandler{}
}

func (p *ParallelSlaveHandler) Execute(context context.Context, command *parallel.Cmd) (output *parallel.Output, err error) {
	output = nil
	err = nil
	output, err = executeCommand(p.p, int(command.Ticket), command.CmdLine)

	// TODO:: recover, handle panics

	return output, err
}

func (p *ParallelSlaveHandler) Ping(context context.Context) (r string, err error) {
	return "ping:ok", nil
}

func mainSlave(p *Parallel) {
	var err error

	p.serverTransport, err = thrift.NewTServerSocket(p.slaveAddress)

	if err != nil {
		log.Fatalln("failed to start server:", err.Error())
		return
	}

	p.handler = NewParallelSlaveHandler()
	p.handler.p = p

	server := thrift.NewTSimpleServer4(
		parallel.NewParallelProcessor(p.handler),
		p.serverTransport,
		p.transportFactory,
		p.protocolFactory)

	err = server.Serve()
	if err != nil {
		log.Fatalln("failed to run slave:", err.Error())
	}
}

func metricsServer(p *Parallel, serverAddress string) {
	metricsHandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})
	http.HandleFunc("/metrics", func(c http.ResponseWriter, req *http.Request) {
		metricsHandler.ServeHTTP(c, req)
	})

	http.HandleFunc("/",
		func(c http.ResponseWriter, req *http.Request) {
			io.WriteString(c,
				fmt.Sprintf(
					"go parallel time: %d slave_address: %s jobs: %d slaves: %d",
					time.Now().Unix(),
					serverAddress,
					p.jobs,
					len(p.Slaves)),
			)
		})

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Println("WARNING: failed to start the metrics server on:", serverAddress, err.Error)
	}
}

func main() {
	T_START := time.Now()
	logger := newLogger(0, false)
	defer func() {
		fmt.Fprintf(logger, "all done: dt: "+time.Since(T_START).String()+"\n")
	}()

	flag_jobs := flag.Int(
		"j",
		2,
		"num of concurrent jobs")

	flag_slave := flag.Bool(
		"slave",
		false,
		"run as slave")

	flag_slaves := flag.String(
		"slaves",
		"",
		"CSV list of slave addresses")

	flag_slave_address := flag.String(
		"address",
		"localhost:9010",
		"slave address")

	flag_slave_metrics_address := flag.String(
		"metrics_address",
		"localhost:9011",
		"slave metric address")

	loggerHostname, _ = os.Hostname()

	flag.Parse()
	fmt.Fprintf(logger, "concurrency limit: %d", *flag_jobs)
	fmt.Fprintf(logger, "slaves: %s", *flag_slaves)

	p := &Parallel{}
	p.jobs = *flag_jobs
	p.logger = logger
	p.worker = limiter.NewConcurrencyLimiter(p.jobs)
	p.slaveAddress = *flag_slave_address
	p.Slaves = []*Slave{}
	for _, slaveAddr := range strings.Split(*flag_slaves, ",") {
		if slaveAddr != "" {
			slave := Slave{Address: slaveAddr}
			p.Slaves = append(p.Slaves, &slave)
		}
	}

	defer p.Close()

	// thrift protocol
	p.protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// thrift transport
	p.transportFactory = thrift.NewTTransportFactory()
	p.transportFactory = thrift.NewTFramedTransportFactory(p.transportFactory)

	// stats:
	p.StatNumCommandsStart = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "commands_num_start",
			Help: "num received"})
	err := prometheus.Register(p.StatNumCommandsStart)
	CheckFatal(err)

	p.StatNumCommandsDone = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "commands_num_done",
			Help: "num received"})
	err = prometheus.Register(p.StatNumCommandsDone)
	CheckFatal(err)

	p.StatCommandLatency = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "commands_latency",
		Help: "commands latency stat",
	})
	err = prometheus.Register(p.StatCommandLatency)
	CheckFatal(err)

	// run the metrics server:
	go metricsServer(p, *flag_slave_metrics_address)

	if *flag_slave == false {
		loggerHostname = p.slaveAddress
		logger.hostname = loggerHostname

		fmt.Fprintf(logger, "running as master\n")
		mainMaster(p)
	} else {
		loggerHostname = p.slaveAddress
		logger.hostname = loggerHostname

		fmt.Fprintf(logger, "running as slave on: %s\n", p.slaveAddress)
		mainSlave(p)

		// stop the metrics thread as well:
		os.Exit(0)
	}
}
