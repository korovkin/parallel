package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	_ "io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/daviddengcn/go-colortext"
	"github.com/korovkin/limiter"
	_ "github.com/korovkin/worker"
)

type logger struct {
	ticket int
}

var (
	loggerMutex     = new(sync.Mutex)
	loggerIndex     = uint32(0)
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

			loggerMutex.Lock()
			ct.ChangeColor(loggerColors[l.ticket%len(loggerColors)], false, ct.None, false)
			fmt.Printf("[%14s %s %d] ", time.Since(loggerStartTime).String(), now, l.ticket)
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

	flag.Parse()
	fmt.Fprintf(logger, fmt.Sprintf("concurrency limit: %d", *flag_jobs))
	worker := limiter.NewConcurrencyLimiter(*flag_jobs)

	r := bufio.NewReaderSize(os.Stdin, 1*1024*1024)
	fmt.Fprintf(logger, "reading from stdin...\n")

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		worker.ExecuteWithTicket(func(ticket int) {
			executeCommand(ticket, line)
		})
	}

	worker.Wait()
}
