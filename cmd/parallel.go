package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/daviddengcn/go-colortext"
	"github.com/korovkin/limiter"
	"github.com/korovkin/parallel"
)

type logger struct {
	ticket   int
	hostname string
	buf      *bytes.Buffer
}

var (
	loggerMutex     = new(sync.Mutex)
	loggerIndex     = int(0)
	loggerStartTime = time.Now()
	loggerHostname  = ""
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
			now := time.Now().Format("15:04:05")
			s := string(line)
			ts := time.Since(loggerStartTime).String()

			loggerMutex.Lock()
			ct.ChangeColor(loggerColors[l.ticket%len(loggerColors)], false, ct.None, false)
			fmt.Printf("[%16s %s %s %d] ", ts, l.hostname, now, l.ticket)
			ct.ResetColor()

			if l.buf != nil {
				l.buf.Write([]byte(s))
			}

			fmt.Print(s)
			loggerMutex.Unlock()

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
	return l
}

func executeCommand(p *Parallel, ticket int, cmdLine string) (*parallel.Output, error) {
	T_START := time.Now()
	var err error
	output := &parallel.Output{}
	loggerOut := newLogger(ticket, true)
	loggerErr := newLogger(ticket, true)

	defer func() {
		fmt.Fprintf(
			loggerOut,
			"execute: done: dt: "+time.Since(T_START).String()+"\n",
		)
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

		output, err = client.Execute(&cmd)
		if err != nil {
			log.Fatalln("failed to execute:", err.Error())
		}

		hostname := output.Tags["hostname"]
		loggerOut.hostname = hostname
		loggerErr.hostname = hostname

		fmt.Fprintf(loggerOut,
			"execute: remotely: host: %s stdout: [%s]\n",
			hostname,
			output.Stdout)

		fmt.Fprintf(loggerErr,
			"execute: remotely: host: %s stderr: [%s]\n",
			hostname,
			output.Stderr)

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

	fmt.Fprintf(loggerOut, "run: '"+cmdLine+"'\n")

	err = cmd.Start()
	if err != nil {
		log.Fatalln("failed to start:", err)
		return output, err
	}

	if err == nil {
		err = cmd.Wait()
	}

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
	Address string `json:"address"`
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
}

func (p *Parallel) Close() {
	p.worker.Wait()
}

func mainMaster(p *Parallel) {
	var err error

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
		ok, err := client.Ping()
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

func (p *ParallelSlaveHandler) Execute(command *parallel.Cmd) (output *parallel.Output, err error) {
	output = nil
	err = nil
	output, err = executeCommand(p.p, int(command.Ticket), command.CmdLine)

	// TODO:: recover, handle panics

	return output, err
}

func (p *ParallelSlaveHandler) Ping() (r string, err error) {
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

	loggerHostname, _ = os.Hostname()

	flag.Parse()
	fmt.Fprintf(logger, fmt.Sprintf("concurrency limit: %d", *flag_jobs))
	fmt.Fprintf(logger, fmt.Sprintf("slaves: %s", *flag_slaves))

	p := Parallel{}
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

	if *flag_slave == false {
		loggerHostname = p.slaveAddress
		logger.hostname = loggerHostname

		fmt.Fprintf(logger, "running as master\n")
		mainMaster(&p)
	} else {
		loggerHostname = p.slaveAddress
		logger.hostname = loggerHostname

		fmt.Fprintf(logger, "running as slave on: %s\n", p.slaveAddress)
		mainSlave(&p)
	}
}
