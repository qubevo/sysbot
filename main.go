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
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	time.Sleep(3 * time.Second)
	rtm.SendMessage(rtm.NewOutgoingMessage("salut ce faci ?", store.GetChannelID()))

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev, rtm)
		}
	}
}
