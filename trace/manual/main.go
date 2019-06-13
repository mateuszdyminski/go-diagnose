package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

// The trace output will be written to the file trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	tracedFunc()
}

func tracedFunc() {
	fmt.Printf("this function will be traced\n")
}
