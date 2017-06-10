package main

import (
	"bufio"
	"flag"
	"io"
	_ "io/ioutil"
	"log"
	"os"
	_ "os/exec"
	"strings"
	"time"

	"github.com/korovkin/limiter"
	_ "github.com/korovkin/worker"
)

func main() {
	T_START := time.Now()
	defer func() {
		log.Println("done dt:", time.Since(T_START))
	}()

	flag_jobs := flag.Int(
		"j",
		1,
		"num of concurrent jobs")

	flag.Parse()
	worker := limiter.NewConcurrencyLimiter(*flag_jobs)

	r := bufio.NewReaderSize(os.Stdin, 1*1024*1024)
	log.Println("reading from stdin...")
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			log.Println("eof")
			break
		}

		line = strings.TrimSpace(line)
		log.Println("execute: ", line)
	}

	log.Println("wait")
	worker.Wait()
}
