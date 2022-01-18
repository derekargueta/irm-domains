package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scannedtext, err := os.Open("scannedtext.txt")
	fixedtext, err2 := os.OpenFile("fixedtext.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
	if err != nil {
		fmt.Println("djhfdksjfh")
	}
	if err2 != nil {
		fmt.Println("djhfdksjfh")
	}
	reader := bufio.NewScanner(scannedtext)
	fixedtextwriter := bufio.NewWriter(fixedtext)

	for reader.Scan() {
		if strings.Count(reader.Text()[8:len(reader.Text())], "/") == 0 {
			fixedtextwriter.WriteString(reader.Text()[8:len(reader.Text())] + "\n")
			fmt.Println(reader.Text()[8:len(reader.Text())])

		}

		fmt.Println(strings.Count(reader.Text()[8:len(reader.Text())], "/"))

	}
	fixedtextwriter.Flush()
}

/*
 fileEntry(filepath string, workers int, cdn_fast probes.Fastlyprobe, cdn_cloud probes.Cloudflareprobe, cdn_max probes.MaxCDN) TotalTestResult {
	domains, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}

	domain := bufio.NewScanner(domains)

	jobs := make(chan string, 300)
	results := make(chan ProbeResult, 1000000)

	log.Printf("Running with %d goroutine workers\n", workers)

	for x := 0; x < workers; x++ {
		go func() {
			worker(jobs, results, cdn_fast, cdn_cloud, cdn_max)
		}()
	}

	log.Println("workers started")
	count := 0
	for domain.Scan() {
		jobs <- domain.Text()
		count++
	}
	close(jobs)

*/
