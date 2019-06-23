package eventhandler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/commands"
)

func registerMessageHandlers(s *discordgo.Session) {
	s.AddHandler(messageCreate)
}

//Calls on every message sent where the bot has access
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isSelf(m.Author.ID, s.State.User.ID) {
		return
	}

	if m.Content[0:len(commands.BotPrefix)] != commands.BotPrefix {
		return // not our message
	}

	//is our message
	err := commands.HandleCommand(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID,"Internal Bot Error - Sorry!")
	}
}

//Determines whether or not the message author is the bot itself. Helps to prevent potential recursive calling
func isSelf(authorID, botID string) bool {
	return authorID == botID
}
