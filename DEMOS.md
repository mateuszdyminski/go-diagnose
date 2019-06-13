
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