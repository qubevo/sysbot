package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/qubevo/sysbot/store"
	"github.com/slack-go/slack"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func handleMessage(ev *slack.MessageEvent, rtm *slack.RTM) {
	fmt.Println(ev.Msg.Text)
	if strings.HasPrefix(ev.Msg.Text, "run") {
		result := strings.TrimPrefix(ev.Msg.Text, "run ")
		cmdFound := false
		for _, d := range store.GetCmds() {
			sp := strings.SplitAfter(d, "}}")
			p := strings.TrimSuffix(sp[0], "}}")
			key := strings.TrimPrefix(p, "{{")
			if key == result {
				cmdFound = true
				out, _, _ := Shellout(sp[1])
				rtm.SendMessage(rtm.NewOutgoingMessage(
					"```"+out+"```",
					store.GetChannelID()))
				break
			} else {
				cmdFound = false
			}
		}
		if !cmdFound {
			rtm.SendMessage(rtm.NewOutgoingMessage("Command not found !", store.GetChannelID()))
		}
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage(ev.Msg.Text, store.GetChannelID()))
	}

}
