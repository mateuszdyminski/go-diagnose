### Build main.go file to simulate that it's running in production

For Docker
```bash
GOOS=linux GOARCH=amd64 go build -gcflags="all=-N -l" -o app .
```

For Mac
```bash
go build -gcflags "all=-N -l" -ldflags=-compressdwarf=false -o app-mac .
```

### To run `app`

```bash
./app-mac
```

### Run Delve headless server next to application which we would like to debug

```bash
dlv attach --headless --listen=:2345 --log --api-version=2 $(pgrep app-mac)
```

### Build Docker image with app and delve
```bash
docker build -t mateuszdyminski/debug:latest .
```

### Push Docker image
```bash
docker push mateuszdyminski/debug
```

### Run Docker image with debug port open
```bash
docker run -d -p 8081:8081 -p 2345:2345 --security-opt=seccomp:unconfined mateuszdyminski/debug:latest
```

### Takeaways

* Unverified breakpoint or variables not loading when debugging a binary

    Ensure that the binary being debugged was built with no optimizations. Use the flags `-gcflags="all=-N -l"` when building the binary.

* Go 1.12 is a step toward a better debugging experience for optimized binaries and we have plans to improve it even further.

* Go 1.11 started compressing debug information to reduce binary sizes. This is natively supported by Delve, but neither LLDB nor GDB support compressed debug info on macOS. If you are using LLDB or GDB, there are two workarounds: build binaries with -ldflags=-compressdwarf=false, or use splitdwarf (go get golang.org/x/tools/cmd/splitdwarf) to decompress the debug information in an existing binary.

* To debug Golang code in Docker we need to use: `--security-opt="apparmor=unconfined"` and `--cap-add=SYS_PTRACE` flags while running docker image

* The application will wait for your IDE to connect to it before starting execution of code, that way you are able to start debugging from start.

* Application will stop in the state completed each time you stop debugging on IDE side (or loose connection)

* your remote debugging will be slower due to network latency between delve and your IDE

* Each time you want to test a new line of code you will need to rebuild the code in debug mode, re-build the Docker Image in debug mode, re-push it to your docker registry, then re-deploy the operator in debug mode.

### CREDITS

* (https://golang.org/doc/gdb)[https://golang.org/doc/gdb]
* (https://blog.golang.org/debugging-what-you-deploy)[https://blog.golang.org/debugging-what-you-deploy]
