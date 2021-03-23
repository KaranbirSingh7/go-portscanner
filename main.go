package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	wg             sync.WaitGroup
	hostFlag       *string
	startPortFlag  *int
	endPortFlag    *int
	timeoutFlag    *string
	pauseFlag      *string
	listClosedFlag *bool
)

func init() {
	// max number of ports available to scan
	const max = 65535

	// Command Line Flags
	hostFlag = flag.String("host", "localhost", "the host to scan")
	startPortFlag = flag.Int("start", 80, "the lower end to scan")
	endPortFlag = flag.Int("end", -1, "the upper end to scan")
	timeoutFlag = flag.String("timeout", "1000ms", "timeout (e.g. 50ms or 1s)")
	pauseFlag = flag.String("pause", "1ms", "pause after each scanned port (e.g. 5ms)")
	listClosedFlag = flag.Bool("closed", false, "list closed ports (true/false)")

	// parse flags
	flag.Parse()

	if *endPortFlag == -1 {
		endPortFlag = startPortFlag
	}
	if *startPortFlag < 1 || *startPortFlag > max {
		log.Fatalf("starting port out of range (should be between 1 and %d)\n", max)
	}
	if *endPortFlag < 1 || *endPortFlag > max {
		log.Fatalf("ending port out of range (should be between 1 and %d)\n", max)
	}
	if *endPortFlag < *startPortFlag {
		log.Fatalln("ending port must be greater than starting port")
	}
}

func main() {
	fmt.Println("Go port scanner starting")
	fmt.Printf("Host: %v", *hostFlag)
	fmt.Printf(" | Start Port: %v", *startPortFlag)
	fmt.Printf(" | End Port: %v", *endPortFlag)
	fmt.Printf(" | Timeout: %v\n\n", *timeoutFlag)

	startTime := time.Now()

	pauseDuration, err := time.ParseDuration(*pauseFlag)
	if err != nil {
		log.Println(err)
	}

	// SCAN ports from here
	for port := *startPortFlag; port <= *endPortFlag; port++ {
		// increment work group of goroutine by 1
		wg.Add(1)

		// scan ports with goroutine
		go scan(*hostFlag, port, *timeoutFlag, *listClosedFlag)

		// wait between another port scan
		time.Sleep(pauseDuration)
	}

	// wait for add goroutine to complete
	wg.Wait()
	scanDuration := time.Since(startTime)
	fmt.Printf("scan completed in %v\n", scanDuration)
}

func scan(host string, port int, timeout string, listClosed bool) {
	// decrease goroutine count by 1 when function completes
	defer wg.Done()

	// parse timeout string
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Println(err)
		return
	}

	// convert port to STRING
	portStr := strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", host+":"+portStr, timeoutDuration)
	if err != nil {
		// if flag is true the print closed ports as well
		if listClosed {
			log.Println(err)
		}
		return
	}
	conn.Close()
	fmt.Printf("open %v\n", port)
}
