package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var printOutput bool

func request(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	if printOutput {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}

	defer resp.Body.Close()
}

func main() {
	var url string
	flag.StringVar(&url, "url", "", "URL of the file to fetch")

	var numReqs int
	flag.IntVar(&numReqs, "nreqs", 0, "Number of requests")

	flag.BoolVar(&printOutput, "print", false, "Should the HTTP response be printed")

	flag.Parse()

	var wg sync.WaitGroup

	start := time.Now()

	for i := 1; i <= numReqs; i++ {
		wg.Add(1)
		go request(url, &wg)
	}

	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Total time taken: ", elapsed)
}
