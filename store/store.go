package store

import (
	"github.com/slack-go/slack"
	"gopkg.in/ini.v1"
)

var Rtm *slack.RTM
var Config *ini.File
var BotEnable bool = true
var ChannelID string = "*"
var BotID string = ""

func CmdsEnabled() bool {
	return true
}

func GetCmds() []string {
	return Config.Section("cmds").Key("cmd").ValueWithShadows()
}

func GetSlackToken() string {
	return Config.Section("").Key("slack_token").String()
}

func GetChannelID(c string) string {
	if Config.Section("").Key("channel_id").String() == "*" {
		return c
	}
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

func MonitorEnable() bool {
	return Config.Section("monitor").Key("enable_monitor").MustBool(false)
}

func GetMonitorTime() int {
	return Config.Section("monitor").Key("run_every_minutes").MustInt(30)
}

func IsOnlyOverhead() bool {
	return Config.Section("").Key("alert_only_overhead").MustBool(false)
}

func FreeShell() bool {
	return Config.Section("").Key("enable_free_shell_commands").MustBool(false)
}

func SaveChannel(c string) {
	ChannelID = c
	Config.Section("").Key("channel_id").SetValue(c)
	Config.SaveTo("config.ini")
}
