package eventhandler

import "github.com/bwmarrin/discordgo"

func RegisterHandlers(s *discordgo.Session) {
	registerMessageHandlers(s) //add more as we go
}
