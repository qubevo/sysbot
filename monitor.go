package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/memory"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func run(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	// fmt.Printf("memory total: %d bytes\n", memory.Total)
	// fmt.Printf("memory used: %d bytes\n", memory.Used)
	// fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n", memory.Free)

}

func monitor() {
	doEvery(10*time.Second, run)
}

//https://socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
//https://github.com/shirou/gopsutil
//https://github.com/mackerelio/go-osstat
