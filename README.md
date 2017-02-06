# grpc_load_test
Small golang utility to load test the grpc route guide example application.

Runs a timed load test for a configurable amount of concurrent requests (via goroutines) and then reports on the statistics it received. Switches are detailed below.  

Currently only exercises the "ListFeatures" example streaming api.

Example output:
```
Timed test.  Will execute test for 1m0s with the below settings
1000 concurrent requests with 5 grpc executions per concurrent request on 8 logical processors
Started...
1 minute(s) elapsed...
...Complete.

Total Test Time took 1m0.420228914s
Started at 2017-02-05 22:36:29.63100387 -0800 PST, ended at 2017-02-05 22:37:30.051232823 -0800 PST.
Completed 165000 grpc calls in 1m0.420220136s.
2730.8738635609266 requests per second.
165000 successful, 0 errored, 
10560000 individual streamed messages
```

## Build
go build -o grpc-load

## Usage

```
-c int
        the number of concurrent requests (default 500)
  -h string
        the hostname (default "localhost")
  -l int
        the length in minutes to run the test for (default 1)
  -lpr int
        the number of times each concurrent request should loop (default 5)
  -p int
        The port (default 8980)
```
