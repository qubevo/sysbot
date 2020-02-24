package main

import (
	"fmt"
	"os"
	"strings"

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

	if store.BotID == "" {
		resp, _ := slackClient.AuthTest()
		user, _ := slackClient.GetUserInfo(resp.UserID)
		store.BotID = user.ID
	}
	store.Rtm = slackClient.NewRTM()
	go store.Rtm.ManageConnection()

	if store.WatchEnabled() {
		go watchFiles()
	}

	if store.MonitorEnable() {
		go monitor()
	}

	// time.Sleep(3 * time.Second)
	// store.Rtm.SendMessage(store.Rtm.NewOutgoingMessage(store.GetIntlStrings("init_msg"), store.GetChannelID()))

	for msg := range store.Rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			pref := "<@" + store.BotID + ">"
			if strings.HasPrefix(ev.Msg.Text, pref) {
				cleanMsg := strings.TrimPrefix(ev.Msg.Text, pref+" ")
				ev.Msg.Text = cleanMsg
				go handleMessage(ev)
			}
		}
	}
}
