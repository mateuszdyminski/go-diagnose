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
go tool pprof -http :8091 -seconds 5 http://localhost:8090/debug/pprof/profile
```

### Open browser with pprof web
```sh
http://localhost:8091/ui/
```

### Comparing profiles
```sh
go tool pprof -http=:8080 --base dumps/heap-profile-cmp-001.pb.gz dumps/heap-profile-cmp-002.pb.gz
```