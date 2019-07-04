package command

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/config"
	shell "github.com/kballard/go-shellquote"
)

/*
func SetupCommands(prefix string, duelStore duel.DuelStore) {
	Commands["duel"] = func(s *discordgo.Session, i IncomingCommand) (err error) {
		if len(i.args) < 1 || len(i.args) > 1 {
			s.ChannelMessageSend(i.m.ChannelID, "No other player specified or too many players specified!")
			return
		}
		_, err = duelStore.DAL.CreateNewDuel(i.m.Author.ID, i.m.Mentions[0].ID)
		if err != nil {
			s.ChannelMessageSend(i.m.ChannelID, err.Error())
			return
		}
		s.ChannelMessageSend(i.m.ChannelID, fmt.Sprintf("%s, you have been challenged to a duel by %s! Type %saccept to accept this duel!", i.m.Mentions[0].Mention(), i.m.Author.Username, BotPrefix))

		return nil
	}

	Commands["accept"] = func(s *discordgo.Session, i IncomingCommand) (err error) {
		accepted, err := duelStore.DAL.AcceptDuel(i.m.Author.ID)
		if err == nil && accepted {
			s.ChannelMessageSend(i.m.ChannelID, "The duel has been accepted. Check your DM's for more information. (WIP)")
			pmChannel, err := s.UserChannelCreate(i.m.Author.ID)
			s.ChannelMessageSend(pmChannel.ID, "boy if you don't")
			return err
		}
		s.ChannelMessageSend(i.m.ChannelID, err.Error())
		return
	}

}
*/
//Not a fan of these globals but I think it makes sense here. Up for evaluation later on..... As long as the commands themselves are testable via DAL structure etc, then I think its fine
var (
	sharedCommands  = make(map[string]Command)
	publicCommands  = make(map[string]Command)
	privateCommands = make(map[string]Command)
)

//ParseIncomingCommand will simply take whatever input and extract out the Keyword (Command name) , Arguments, and the actual message being passed from discordgo.
//This does not do any sort of real validation and is a pure function.
func ParseIncomingCommand(s *discordgo.Session, m *discordgo.MessageCreate) (i IncomingCommand, err error) {
	commandArgs, err := shell.Split(strings.TrimPrefix(m.Content, config.Config.BotPrefix))
	if err != nil {
		return
	}

	if len(commandArgs) < 1 { //must have a keyword
		return i, errors.New("Could not parse command")
	}
	i.KeyWord = commandArgs[0]

	if len(commandArgs) > 1 { //has arguments passed
		i.Args = commandArgs[1:]
	}

	i.Message = m

	channel, err := s.State.Channel(m.ChannelID) // check if private channel or public
	if err != nil {
		return i, errors.New("Cannot find source of message")
	}

	i.Source = PUBLIC
	if channel.Type == discordgo.ChannelTypeDM {
		i.Source = PRIVATE
	}

	//temp for now??? - ideally in validate but i want to avoid calling the API _again_ to get the channel for now
	if channel.Type == discordgo.ChannelTypeGroupDM {
		return i, errors.New("Group DMs are not supported yet")
	}

	return
}

//validateIncomingCommand will handle checking argument count, permissions, and ensure the command we are trying to use exists
func validateIncomingCommand(i IncomingCommand) error {
	c, err := getCommandFromKeyword(i.KeyWord)
	if err != nil {
		return err
	}

	if c.Permissions != SHARED && (c.Permissions != i.Source) { //prevent public / private commands from being executed outside of their areas
		return errors.New("Command not allowed in this context")
	}

	switch len(c.ArgCount) {
	case 1:
		if c.ArgCount[0] != len(i.Args) {
			return errors.New("Argument Mismatch") //eventually will need to move all of these errors into a standardized thing
		}
	case 2:
		if len(i.Args) < c.ArgCount[0] || len(i.Args) > c.ArgCount[1] {
			return errors.New("Argument Mismatch") //eventually will need to move all of these errors into a standardized thing
		}
	}

	//maybe validate it isn't coming from a banned account later on? Is a bot allowed to ban someone from using it? Idk.

	return nil
}

func getCommandFromKeyword(keyword string) (c Command, err error) {
	if command, ok := sharedCommands[keyword]; ok {
		return command, nil
	} else if command, ok := publicCommands[keyword]; ok {
		return command, nil
	} else if command, ok := privateCommands[keyword]; ok {
		return command, nil
	}

	return Command{}, errors.New("Command does not exist")
}

//AddCommand will add the command to the appropriate map that is defined.
func AddCommand(c Command) error {
	errTaken := errors.New("Keyword already taken")
	switch c.Permissions {
	case PUBLIC:
		if _, ok := publicCommands[c.KeyWord]; ok {
			return errTaken
		}
		publicCommands[c.KeyWord] = c
	case PRIVATE:
		if _, ok := privateCommands[c.KeyWord]; ok {
			return errTaken
		}
		privateCommands[c.KeyWord] = c
	case SHARED:
		if _, ok := sharedCommands[c.KeyWord]; ok {
			return errTaken
		}
		sharedCommands[c.KeyWord] = c
	}

	return nil
}
