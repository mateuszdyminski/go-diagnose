package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	memFile, _ := os.Create("./pprof.mem")
	cpuFile, _ := os.Create("./pprof.cpu")
	defer memFile.Close()
	defer cpuFile.Close()

	pprof.StartCPUProfile(cpuFile)
	go leakyFunction()

	// wait some time for initialization of leakyFunction
	time.Sleep(500 * time.Millisecond)

	pprof.WriteHeapProfile(memFile)
	pprof.StopCPUProfile()
}

func leakyFunction() {
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pprof time")
	}
}
