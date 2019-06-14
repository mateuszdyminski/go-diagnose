
# DEBUG Local
cd debug
code .
<Run server from VS Code>
curl http://localhost:8081/add?vals=1,2,4
<Set breakpoint>
<Show in VS Code that we hit the breakpoint>


# DEBUG Remote
cd debug
code .
go build -gcflags "all=-N -l" -ldflags=-compressdwarf=false -o app-mac .
<Run app>
./app-mac
<Split terminal horizontally>
<Check if it's working>
curl http://localhost:8081/add?vals=1,2,4
dlv attach --headless --listen=:2345 --log --api-version=2 --accept-multiclient=true $(pgrep app-mac)
<Open VS Code and show the launch.json file>
<Set breakpoint>
<Run debugger>
curl http://localhost:8081/add?vals=1,2,4
<Show in VS Code that we hit the breakpoint>


# DEBUG Docker
cd debug
code .
docker build -t mateuszdyminski/debug:latest .
docker run -d -p 8081:8081 -p 40000:40000 --security-opt=seccomp:unconfined mateuszdyminski/debug:latest
<Split terminal horizontally>
<Check if it's working>
curl http://localhost:8081/add?vals=1,2,4
<Open VS Code and show the launch.json file>
<Set breakpoint>
<Run debugger>
curl http://localhost:8081/add?vals=1,2,4
<Show in VS Code that we hit the breakpoint>

# PROFILING
cd profile
go run main.go
<Split terminal horizontally>
<Check if it's working>
curl http://localhost:8090/hello
curl http://localhost:8090/statsHello
go run main.go --printStats
<Describe what's on the screen - paths>
go run main.go
<Run wrk to generate load>
wrk -d 5s http://localhost:8090/hello
wrk -d 5s http://localhost:8090/statsHello
<Describe results>
<Run load generator with high value>
wrk -d 30s http://localhost:8090/statsHello
<Run profiler in new split terminal>
go tool pprof -seconds 5 http://localhost:8090/debug/pprof/profile
<In pprof tool call>
top
top25
list WithStats
list getStatsTags
go tool pprof -http :8091 -seconds 5 http://localhost:8090/debug/pprof/profile
<Show graph>
<Show flamegraph>
<Refactor code>
<Run wrk once again>
wrk -d 5s http://localhost:8090/statsHello

# EXECUTION TRACER
cd trace
export GOMAXPROCS=1
go run main.go
<Split terminal>
wrk -d 3s http://localhost:8090/hello
<Describe results>
<Generate load for longer time>
wrk -d 30s http://localhost:8090/hello
curl http://localhost:8090/debug/pprof/trace?seconds=3 > 1.trace
<Open trace>
go tool trace 1.trace
<Show tracer>
<Fix GOMAXPROCS>
export GOMAXPROCS=8
<Rerun server>
<Show tracer>
<Refactor code - sync.Mutex-> atomic.AddInt32(&counter, 1)>
<Rerun server>
<Show tracer>
<Refactor code - remove sleep>
<Rerun server>
<Show tracer>


# INSTRUMENTATION WITH PROMETHEUS
cd instrument
code .
<Describe logs in terminal>
<Open /metrics endpoint>
http://localhost:8095/metrics
<Open VS Code and talk a bit about custom metrics>
http://localhost:8095/register?name=gogoconf
http://localhost:8095/register?name=golang
http://localhost:8095/metrics
