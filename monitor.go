package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/qubevo/sysbot/store"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func run(t time.Time) {
	if store.BotEnable {
		fmt.Printf("%v: Hello, World!\n", t)
		memory, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		if !store.IsOnlyOverhead() {
			// fmt.Printf("memory total: %d bytes\n", memory.Total)
			// fmt.Printf("memory used: %d bytes\n", memory.Used)
			// fmt.Printf("memory cached: %d bytes\n", memory.Cached)
			fmt.Printf("memory free: %d bytes\n", memory.Free)
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage("mem", store.GetChannelID()))
		} else {
			//process overhead
		}
	}
}

func monitor() {
	fmt.Println(store.GetMonitorTime())
	mult := time.Duration(store.GetMonitorTime())
	doEvery(mult*time.Minute, run)
}

//https://socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
//https://github.com/shirou/gopsutil
//https://github.com/mackerelio/go-osstat
