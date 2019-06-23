package commands

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/cards"
)

var BotPrefix = "y!"
var Commands map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (err error)

var cardListEnabled = false

//Honestly, not too sure what im doing here... I am not a fan of the global vars on everything so far... ^^^^^^
//Gonna get some stuff set up then refactor  from there I guess. Make it work -> Make it good -> Make it fast ????

func SetupCommands(prefix string, newCardListEnabled bool) {
	BotPrefix = prefix                   // just to allow the config to let this be custom
	cardListEnabled = newCardListEnabled // incase in the future some places don't like being able to list all cards (good for lots of cards/customs)

	Commands = make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (err error))
	Commands["cards"] = func(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
		var b bytes.Buffer
		for _, v := range cards.Cards.Cards {
			b.WriteString(fmt.Sprintf("%v\n", v))
		}
		s.ChannelMessageSend(m.ChannelID, b.String())
		return nil
	}
}

//Not very well done for now - will refactor once I start working on the game portion of the bot
func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	command := m.Content[len(BotPrefix):]
	return Commands[command](s, m)
}
