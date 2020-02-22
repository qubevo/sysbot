package main

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
	"github.com/qubevo/sysbot/store"
)

func watchFiles() {
	for _, f := range store.GetWatchFiles() {
		_, value := parseLine(f)
		t, e := tail.TailFile(value, tail.Config{Follow: true})
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		for line := range t.Lines {
			if store.BotEnable {
				fmt.Println(line.Text)
			}
		}
	}
}
