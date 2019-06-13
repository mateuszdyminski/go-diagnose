### To run benchmark with CPU profile
```sh
go test -bench . -benchmem -cpuprofile prof.cpu
```

### To run benchmark with MEM profile
```sh
go test -bench . -benchmem -memprofile prof.mem
```

### To open pprof analysis in terminal
```sh
go tool pprof prof.cpu
```

### To open pprof analysis from terminal in interactive web 
```sh
go tool pprof -http :8091 prof.cpu
```

go test . -bench . -memprofile prof.mem