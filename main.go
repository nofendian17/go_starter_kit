package main

import (
	"github.com/nofendian17/gostarterkit/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
