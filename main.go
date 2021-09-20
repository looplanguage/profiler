package main

import (
	"github.com/looplanguage/profiler/benchmark"
	"os"
)

func main() {
	benchmark.StartBenchmark(os.Args[1])
}
