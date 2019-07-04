package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/colbyleiske/yugioh-bot/card"
	"github.com/colbyleiske/yugioh-bot/config"
	"github.com/colbyleiske/yugioh-bot/eventhandler"
)

//not great code, but I really just want to focus on getting this to work. I always tend to overcomplicate everything way to early and get overwhelmed. Trying to K.I.S.S. this one
func main() {

	config.ReadConfig()

	memStore := &card.MemoryDatastore{}
	if err := memStore.ReadCards(); err != nil {
		log.Fatal(err) //don't bother running without cards
	}
	cards := card.NewCards(memStore)
	cards.SetupCommands()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Config.BotToken)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	eventhandler.RegisterHandlers(dg)

	// Open a websocket connection to Discord and begin listening.
	if err = dg.Open(); err != nil {
		log.Fatal("error opening connection,", err)
		return
	}
	defer dg.Close()
	dg.UpdateStatus(0, config.Config.GameName)

	// Wait here until CTRL-C or other term signal is received.
	log.Println("It's time to D-D-D-D-D-D-DUEL")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
