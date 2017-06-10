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
	ticket int
}

var (
	loggerMutex     = new(sync.Mutex)
	loggerIndex     = int(0)
	loggerStartTime = time.Now()
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
			fmt.Printf("[%14s %s %d] ", ts, now, l.ticket)
			ct.ResetColor()
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

func newLogger(ticket int) *logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	l := &logger{ticket}
	return l
}

func executeCommand(ticket int, cmdLine string) bool {
	T_START := time.Now()
	logger := newLogger(ticket)

	defer func() {
		fmt.Fprintf(logger, "done: dt: "+time.Since(T_START).String()+"\n")
	}()

	cs := []string{"/bin/sh", "-c", cmdLine}
	cmd := exec.Command(cs[0], cs[1:]...)
	cmd.Stdin = nil
	cmd.Stdout = logger
	cmd.Stderr = logger
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("PARALLEL_TICKER=%d", ticket),
	)

	fmt.Fprintf(logger, "run: '"+cmdLine+"'\n")

	err := cmd.Start()
	if err != nil {
		log.Fatalln("failed to start:", err)
		return true
	}

	err = cmd.Wait()
	return true
}

type Parallel struct {
	jobs    int
	logger  *logger
	worker  *limiter.ConcurrencyLimiter
	address string
}

func mainMaster(p *Parallel) {
	{
		var transport thrift.TTransport
		var err error
		transport, err = thrift.NewTSocket(p.address)

		if transport == nil {
			log.Fatalln("failed allocate transport:")
		}

		if err != nil {
			log.Fatalln("failed to dial slave:", err.Error())
		}

		var transportFactory thrift.TTransportFactory
		transportFactory = thrift.NewTTransportFactory()
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

		transport = transportFactory.GetTransport(transport)
		defer transport.Close()

		err = transport.Open()
		if err != nil {
			log.Fatalln("failed to open:", err.Error())
		}

		var protocolFactory thrift.TProtocolFactory
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

		client := parallel.NewParallelClientFactory(transport, protocolFactory)
		client.Ping()
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
			executeCommand(ticket, line)
		})
	}
}

type ParallelSlaveHandler struct {
}

func NewParallelSlaveHandler() *ParallelSlaveHandler {
	return &ParallelSlaveHandler{}
}

func (p *ParallelSlaveHandler) Execute(command *parallel.Cmd) (r string, err error) {
	log.Println("ParallelSlaveHandler: Execute: ", command.CmdLine)
	return "ok", nil
}

func (p *ParallelSlaveHandler) Ping() (r string, err error) {
	log.Println("ParallelSlaveHandler: Ping")
	return "ok", nil
}

func mainSlave(p *Parallel) {
	var err error

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(p.address)

	if err != nil {
		log.Fatalln("failed to start server:", err.Error())
		return
	}

	handler := NewParallelSlaveHandler()
	processor := parallel.NewParallelProcessor(handler)
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory)

	err = server.Serve()
	if err != nil {
		log.Fatalln("failed to run slave:", err.Error())
	}
}

func main() {
	T_START := time.Now()
	logger := newLogger(0)
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

	flag.Parse()
	fmt.Fprintf(logger, fmt.Sprintf("concurrency limit: %d", *flag_jobs))

	p := Parallel{}
	p.jobs = *flag_jobs
	p.logger = logger
	p.worker = limiter.NewConcurrencyLimiter(p.jobs)
	p.address = "localhost:9010"

	if *flag_slave == false {
		fmt.Fprintf(logger, fmt.Sprintf("running as master\n"))
		mainMaster(&p)
	} else {
		fmt.Fprintf(logger, fmt.Sprintf("running as slave\n"))
		mainSlave(&p)
	}

	p.worker.Wait()
}
