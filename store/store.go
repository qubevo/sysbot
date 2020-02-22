package store

import (
	"github.com/slack-go/slack"
	"gopkg.in/ini.v1"
)

var Rtm *slack.RTM
var Config *ini.File
var BotEnable bool = false

func CmdsEnabled() bool {
	return true
}

func GetCmds() []string {
	return Config.Section("cmds").Key("cmd").ValueWithShadows()
}

func GetSlackToken() string {
	return Config.Section("").Key("slack_token").String()
}

func GetChannelID() string {
	return Config.Section("").Key("channel_id").String()
}

func WatchEnabled() bool {
	return Config.Section("watchfiles").Key("enable_watchfiles").MustBool(false)
}

func GetWatchFiles() []string {
	return Config.Section("watchfiles").Key("file").ValueWithShadows()
}

func GetIntlStrings(k string) string {
	return Config.Section("bot_messages").Key(k).String()
}

func IsMulti() bool {
	return Config.Section("").Key("multi_agents").MustBool(false)
}

func GetAgentName() string {
	return Config.Section("").Key("agent_name").String()
}

// get enable_monitor
