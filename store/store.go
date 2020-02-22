package store

import "gopkg.in/ini.v1"

var Config *ini.File

func GetCmds() []string {
	return Config.Section("cmds").Key("cmd").ValueWithShadows()
}

func GetSlackToken() string {
	return Config.Section("").Key("slack_token").String()
}

func GetChannelID() string {
	return Config.Section("").Key("channel_id").String()
}
