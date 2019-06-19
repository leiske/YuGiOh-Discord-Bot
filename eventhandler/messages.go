package eventhandler

import "github.com/bwmarrin/discordgo"

func registerMessageHandlers(s *discordgo.Session) {
	s.AddHandler(messageCreate)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isSelf(m.Author.ID, s.State.User.ID) {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "test")
	}
}

func isSelf(authorID, botID string) bool {
	return authorID == botID
}
