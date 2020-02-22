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

func handleMessage(ev *slack.MessageEvent) {
	fmt.Println(ev.Msg.Text)
	if !store.BotEnable {
		if ev.Msg.Text == "start" {
			store.BotEnable = true
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("enabled_msg"), store.GetChannelID()))
		} else {
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("disabled_msg"), store.GetChannelID()))
		}
		return
	}
	if ev.Msg.Text == "stop" {
		store.BotEnable = false
		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("disabled_msg"), store.GetChannelID()))
		return
	}
	if strings.HasPrefix(ev.Msg.Text, "run") {
		result := strings.TrimPrefix(ev.Msg.Text, "run ")
		cmdFound := false
		for _, d := range store.GetCmds() {
			key, value := parseLine(d)
			if key == result {
				cmdFound = true
				out, _, _ := Shellout(value)
				store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(
					"```"+out+"```",
					store.GetChannelID()))
				break
			} else {
				cmdFound = false
			}
		}
		if !cmdFound {
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("cmd_not_found"), store.GetChannelID()))
		}
		return
	}
	if strings.HasPrefix(ev.Msg.Text, "exec") {
		fmt.Println(" --- exec ---")
		return
	}

	store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("i_dont_know"), store.GetChannelID()))

}

func parseLine(d string) (k string, value string) {
	sp := strings.SplitAfter(d, "}}")
	p := strings.TrimSuffix(sp[0], "}}")
	key := strings.TrimPrefix(p, "{{")
	return key, sp[1]
}
