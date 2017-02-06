package main

import "flag"
import (
	"./load_test"
	"time"
	"fmt"
	"runtime"
	"encoding/json"
	"os"
)

func main() {
	hostName := flag.String("h", "localhost", "the hostname")
	port := flag.Int("p", 8980, "The port")
	minutes := flag.Int("l", 1, "the length in minutes to run the test for")
	concurrentRequests := flag.Int("c", 500, "the number of concurrent requests")
	loopsPerReqeust := flag.Int("lpr", 5, "the number of times " +
		"each concurrent request should loop")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}


	testLength := time.Minute * time.Duration(*minutes)
	testRequest := load.LoadTestRequest{
		Hostname: *hostName,
		Port: *port,
		TestLength: testLength,
		ConcurrentRequests: *concurrentRequests,
		LoopsPerRequest: *loopsPerReqeust,
	}
	requestJson, err := json.MarshalIndent(&testRequest,"", "  ")
	if err != nil {
		fmt.Errorf("Error parsing json: %v", err)
		os.Exit(1)
	}
	fmt.Println("Test Details: \n", string(requestJson))
	fmt.Printf("Timed test.  Will execute test for %v with the below settings\n", testLength)
	fmt.Printf("%v concurrent requests with %v grpc executions per concurrent request " +
		"on %v logical processors\n",*concurrentRequests, *loopsPerReqeust, runtime.NumCPU())
	fmt.Println("Started...")

	start := time.Now()
	results := load.Loadtest(testRequest)
	duration := time.Since(start)
	fmt.Println("...Complete.")

	fmt.Printf("\nTotal Test Time took %v\nStarted at %v, ended at %v.\n",time.Since(start), start, time.Now())
	fmt.Printf("Completed %v grpc calls in %v.\n%v requests per second.\n",
		results.TotalExecutions, duration, float64(results.TotalExecutions) / duration.Seconds())
	fmt.Printf("%v successful, %v errored, \n", results.SuccessfulExecutions, results.ErrorExecutions)
	fmt.Printf("%v individual streamed messages", results.StreamedMessages)
}
