package eventhandler

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/command"
	"github.com/colbyleiske/yugioh-bot/config"
)

func registerMessageHandlers(s *discordgo.Session) {
	s.AddHandler(messageCreate)
}

//Calls on every message sent where the bot has access
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, config.Config.BotPrefix) || m.Author.Bot {
		//ignore ourselves, not prefixed messages, or other bots
		return
	}

	incomingCommand, err := command.ParseIncomingCommand(s, m)
	if err != nil {
		//just exit if we cant parse
		log.Println(err)
		return
	}

	if err = command.RouteCommand(s, incomingCommand); err != nil {
		log.Println("Could not run the command:", incomingCommand, "with error", err)
		return
	}

}
