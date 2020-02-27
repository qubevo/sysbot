package main

const MSG_HELP string = "help"
const MSG_START string = "start"
const MSG_STOP string = "stop"
const MSG_RUN string = "run"
const MSG_EXEC string = "exec"
const MSG_SET_CHANNEL string = "set channel"
const MSG_RESET_CHANNEL string = "reset channel"
const MSG_ENABLE_MONITOR string = "enable monitor"
const MSG_DISABLE_MONITOR = "disable monitor"

const MSG_HELP_REPLY = "```" + `
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
