package load

import (
	"time"
	"sync/atomic"
	"math/rand"
	"google.golang.org/grpc"
	"context"
	"sync"
	"strconv"
	"io"
	"fmt"
)

type LoadTestRequest struct {
	Hostname string
	Port int
	TestLength time.Duration
	ConcurrentRequests int
	LoopsPerRequest int
}

type ExecutionResults struct {
	TotalExecutions      uint64
	SuccessfulExecutions uint64
	ErrorExecutions    uint64
	StreamedMessages     uint64
}

var executionResults = ExecutionResults{}
func Loadtest(test LoadTestRequest) ExecutionResults {

	rand.Seed(time.Now().Unix())
	endLoadTest := time.Now().Add(test.TestLength)

	conn, err := grpc.Dial(test.Hostname+":"+strconv.Itoa(test.Port), grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go printElapsed(time.Now(), endLoadTest)
	for time.Now().Before(endLoadTest) {
		paralellLoadTest(test.ConcurrentRequests, test.LoopsPerRequest, conn)
	}

	return executionResults
}

func printElapsed(start, end time.Time) {
	for time.Now().Before(end) {
		time.Sleep(time.Minute)
		fmt.Printf("%v minute(s) elapsed...\n", int(time.Since(start).Minutes()))
	}
}

func paralellLoadTest(numGoRoutines, numLoopsPerRoutine int, conn *grpc.ClientConn ) {
	var wg sync.WaitGroup
	//start := time.Now()
	wg.Add(numGoRoutines)
	for i := 1; i <= numGoRoutines; i++ {
		go asyncClientExecution(&wg, conn, numLoopsPerRoutine)
	}
	wg.Wait()
}

func asyncClientExecution ( wg *sync.WaitGroup, conn *grpc.ClientConn, executionsPer int  ) {
	executeClient(executionsPer, conn)
	wg.Done()

}


func executeClient(numCalls int, conn *grpc.ClientConn  ) {
	for i := 1; i <= numCalls; i ++ {
		listFeatures(conn)
		atomic.AddUint64(&executionResults.TotalExecutions, 1)
	}

}


func listFeatures( conn *grpc.ClientConn ) {

	routeGuideClient := NewRouteGuideClient(conn)

	stream, err := routeGuideClient.ListFeatures(context.Background(), &Rectangle{
		Hi: &Point{
			Longitude: 800000000,
			Latitude: 800000000,
		},
		Lo: &Point{
			Longitude: -800000000,
			Latitude: -800000000,
		},


	})

	if err != nil {
		atomic.AddUint64(&executionResults.ErrorExecutions, 1)
		return
	}
	streamedCount := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			atomic.AddUint64(&executionResults.ErrorExecutions, 1)
			return
		}
		streamedCount++
	}
	atomic.AddUint64(&executionResults.StreamedMessages, uint64(streamedCount))
	atomic.AddUint64(&executionResults.SuccessfulExecutions, 1)
}