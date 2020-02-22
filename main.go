package main

import (
	"fmt"
	"os"
	"time"

	"github.com/qubevo/sysbot/store"

	"github.com/slack-go/slack"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.ShadowLoad("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	store.Config = cfg

	slackClient := slack.New(store.GetSlackToken(), slack.OptionDebug(false))
	store.Rtm = slackClient.NewRTM()
	go store.Rtm.ManageConnection()

	if store.WatchEnabled() {
		go watchFiles()
	}

	go monitor()
	time.Sleep(3 * time.Second)
	store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(
		prepareMessage(store.GetIntlStrings("init_msg")),
		store.GetChannelID()))

	for msg := range store.Rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev)
		}
	}
}
