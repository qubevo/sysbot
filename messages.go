package main

import (
	"bytes"
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
	// fmt.Println(ev.Msg.Text)
	if ev.Msg.Text == MSG_HELP {

		help := "```" + `
Commands:

set channel - will set the channel(with user or slack channel) where sysbot can communicate.
reset channel - will reset the channel, which allow sysbot to communicate with any channel and user.
run <cmd label> - to run a predifined shell command or script
exec <shell cmd> - will execute any shell command if the feature is enabled.(Can be enabled ONLY from config file)
enable monitor - will enable sysbot to monitor system CPU, MEMORY and DISK. ** Works ONLY if a channel is set
disable monitor - will disable sysbot system monitoring.

More configurable features are vailable in the config file !
- watch files
- cron scripts and commands
- configure, enable or disable remote commands
- configure monitoring 
and more...
		` + "```"

		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(help, ev.Channel))

		return
	}
	if (store.ChannelID != ev.Channel) && (store.ChannelID != "*") {
		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage("Sorry, I am not allowed to speak with you !", ev.Channel))
		return
	}
	if ev.Msg.Text == MSG_RESET_CHANNEL {
		store.SaveChannel("*")
		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage("This channel was unset !", store.GetChannelID(ev.Channel)))
		return
	}
	if ev.Msg.Text == MSG_SET_CHANNEL {
		store.SaveChannel(ev.Channel)
		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage("Channel was saved !", store.GetChannelID(ev.Channel)))
		return
	}
	// if !store.BotEnable {
	// 	if ev.Msg.Text == MSG_START {
	// 		store.BotEnable = true
	// 		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("enabled_msg"), store.GetChannelID(ev.Channel)))
	// 	} else {
	// 		store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("disabled_msg"), store.GetChannelID(ev.Channel)))
	// 	}
	// 	return
	// }
	// if ev.Msg.Text == MSG_STOP {
	// 	store.BotEnable = false
	// 	store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("disabled_msg"), store.GetChannelID(ev.Channel)))
	// 	return
	// }
	if strings.HasPrefix(ev.Msg.Text, MSG_RUN) {
		result := strings.TrimPrefix(ev.Msg.Text, MSG_RUN+" ")
		cmdFound := false
		for _, d := range store.GetCmds() {
			key, value := parseLine(d)
			if key == result {
				cmdFound = true
				out, _, _ := Shellout(value)
				store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(
					"```"+out+"```",
					store.GetChannelID(ev.Channel)))
				break
			} else {
				cmdFound = false
			}
		}
		if !cmdFound {
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("cmd_not_found"), store.GetChannelID(ev.Channel)))
		}
		return
	}
	if strings.HasPrefix(ev.Msg.Text, MSG_EXEC) {
		if store.FreeShell() {
			result := strings.TrimPrefix(ev.Msg.Text, MSG_EXEC+" ")
			out, _, _ := Shellout(result)
			store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(
				"```"+out+"```",
				store.GetChannelID(ev.Channel)))
			return
		}
	}

	store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("i_dont_know"), store.GetChannelID(ev.Channel)))

}

func parseLine(d string) (k string, value string) {
	sp := strings.SplitAfter(d, "}}")
	p := strings.TrimSuffix(sp[0], "}}")
	key := strings.TrimPrefix(p, "{{")
	return key, sp[1]
}
