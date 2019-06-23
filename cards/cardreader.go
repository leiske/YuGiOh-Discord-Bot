package cards

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/colbyleiske/yugioh-bot/model"
)

type CardsList struct {
	Cards []model.Card `json:"cards"`
}

var Cards CardsList //Global card list because we don't really care right now... If this gets built enough for random people
//Perhaps I can implement a system for custom cards but that is so far into the future.

//ReadCards will take a path to a cardlist. This opens up the possibility of custom cards if hosted locally.
func ReadCards(cardsPath string) error {
	cardsJson, err := os.Open(cardsPath)
	if err != nil {
		return err
	}
	byteValue, err := ioutil.ReadAll(cardsJson)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteValue, &Cards); err != nil {
		return err
	}

	return nil
}
