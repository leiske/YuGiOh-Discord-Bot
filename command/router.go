package command

import (
	"github.com/bwmarrin/discordgo"
)

//RouteCommand will call the appropriate command. This assumes it has passed parsing. Will run the command and validate before.
//This will get reworked as I continue to refine the command layout and structure. Already on the third rewrite :)
func RouteCommand(s *discordgo.Session, i IncomingCommand) (err error) {
	if err = validateIncomingCommand(i); err != nil {
		return
	}

	c, err := getCommandFromKeyword(i.KeyWord)
	if err != nil {
		return
	}

	return c.Handler(s, i)
}
