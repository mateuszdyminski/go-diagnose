### Run HTTP server
```sh
go run main.go
```

### Run HTTP server with printing 
```sh
go run main.go --printStats=true
```

### Endpoint hello world
(http://localhost:8090/hello)[http://localhost:8090/hello]

### Endpoint hello world with stats
(http://localhost:8090/statsHello)[http://localhost:8090/statsHello]

### Endpoint with pprof debug
(http://localhost:8090/debug/pprof/)[http://localhost:8090/debug/pprof/]

### Generate load on server for 10s
```sh
wrk -d 10s http://localhost:8090/statsHello
```

### Generate profile from running app
```sh
go tool pprof -http :8091 -seconds 5 app http://localhost:8090/debug/pprof/profile
```

### Open browser with pprof web
```sh
http://localhost:8091/ui/
```

go tool pprof -seconds 5 http://localhost:9090/debug/pprof/profile
	top10: to show the top 10 functions by time spent only in that function	top10 -cum: the top 10 functions by time spent in that function or a function it called.	list regex: to show code for functions matching regex	disasm regex: to show the disassembly for functions matching regex	
go test -bench . -benchmem -cpuprofile prof.cpu -memprofile prof.memgo tool pprof stats.test prof.cpu
go tool pprof -http :8091 -seconds 5  app http://localhost:8090/debug/pprof/profile
go tool pprof -alloc_objects stats.test prof.mem
go-torch -u http://localhost:9090 --time 5
go-torch --binaryname stats.test -b prof.cpu
go build -gcflags=-m .