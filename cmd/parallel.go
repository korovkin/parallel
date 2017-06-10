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

func executeCommand(p *Parallel, ticket int, cmdLine string) bool {
	T_START := time.Now()
	logger := newLogger(ticket)

	defer func() {
		fmt.Fprintf(logger, "done: dt: "+time.Since(T_START).String()+"\n")
	}()

	output, err := p.client.Execute(&parallel.Cmd{
		CmdLine: cmdLine,
		Ticket:  int64(ticket),
	})
	if err != nil {
		log.Fatalln("failed to execute:", err.Error())
	}
	fmt.Fprintf(logger, "execute: output: '"+output+"'\n")

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

	err = cmd.Start()
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

	// master / slave
	protocolFactory  thrift.TProtocolFactory
	transportFactory thrift.TTransportFactory

	// master:
	transport thrift.TTransport
	client    *parallel.ParallelClient

	// slave:
	handler *ParallelSlaveHandler
}

func (p *Parallel) Close() {
	if p.transport != nil {
		p.transport.Close()
	}
}

func mainMaster(p *Parallel) {
	var err error
	p.transport, err = thrift.NewTSocket(p.address)

	if p.transport == nil {
		log.Fatalln("failed allocate transport")
	}

	if err != nil {
		log.Fatalln("failed to dial slave:", err.Error())
	}

	p.transport = p.transportFactory.GetTransport(p.transport)

	err = p.transport.Open()
	if err != nil {
		log.Fatalln("failed to open:", err.Error())
	}

	p.client = parallel.NewParallelClientFactory(p.transport, p.protocolFactory)
	p.client.Ping()

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

	p.handler = NewParallelSlaveHandler()

	server := thrift.NewTSimpleServer4(
		parallel.NewParallelProcessor(p.handler),
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

	// thrift protocol
	p.protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// thrift transport
	p.transportFactory = thrift.NewTTransportFactory()
	p.transportFactory = thrift.NewTFramedTransportFactory(p.transportFactory)

	if *flag_slave == false {
		fmt.Fprintf(logger, fmt.Sprintf("running as master\n"))
		mainMaster(&p)
	} else {
		fmt.Fprintf(logger, fmt.Sprintf("running as slave\n"))
		mainSlave(&p)
	}

	p.worker.Wait()
	p.Close()
}
