package main

import (
	"os"
	"runtime/pprof"

	"github.com/timetravel-1010/indexer/cmd"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	cmd.Execute()
}
