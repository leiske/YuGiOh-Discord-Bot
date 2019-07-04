package card

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/command"
)

func (c *Cards) SetupCommands() {
	command.AddCommand(
		command.Command{
			KeyWord:     "card",
			ArgCount:    []int{1},
			Permissions: command.SHARED,
			Handler:     c.GetCardByIDOrNameCommand,
		},
	)
}

func (c Cards) GetCardByIDOrNameCommand(s *discordgo.Session, i command.IncomingCommand) error {
	card, err := c.Datastore.GetCardByIDOrName(i.Args[0])
	if err != nil {
		return err
	}

	cardEmbed := &discordgo.MessageEmbed{
		Title: card.Name,
		Image: &discordgo.MessageEmbedImage{
			URL: card.CardImages[0].ImageURL,
		},
		Description: card.Description,
		Color:       16312092,
		Fields: [](*discordgo.MessageEmbedField){
			&discordgo.MessageEmbedField{Name: "Stats", Value: fmt.Sprintf("**ATK**: %s\n**DEF**: %s\n", card.ATK, card.DEF), Inline: true},
			&discordgo.MessageEmbedField{Name: "Misc", Value: fmt.Sprintf("**LVL**: %s\n", card.Level), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("ID: %s", card.ID),
		},
	}

	s.ChannelMessageSendEmbed(i.Message.ChannelID, cardEmbed)
	return nil
}
