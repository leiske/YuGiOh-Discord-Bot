package command

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	KeyWord     string
	ArgCount    []int //is an array but only want to use at MOST 2 spots... [1,5] -> 1 to 5 arguments.
	Permissions Permission
	Handler     func(*discordgo.Session, IncomingCommand) error
}

//IncomingCommand is a small wrapper on top of the message struct that has convience things such as the keyword and the args in their own array already. No parsing inside of the command needed
type IncomingCommand struct {
	KeyWord string
	Args    []string
	Source  Permission //just include whether or not channel is a DM or a public one
	Message *discordgo.MessageCreate
}

type Permission int

const (
	SHARED = iota
	PUBLIC
	PRIVATE
)
