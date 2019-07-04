package card

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/colbyleiske/yugioh-bot/config"
)

type DAL interface {
	GetCardByIDOrName(string) (Card, error)
}

//not needed by DAL but something we want to use
func (d *MemoryDatastore) ReadCards() error {

	type JSONCardsList struct {
		Cards []Card `json:"cards"`
	}

	var JSONCards JSONCardsList
	cardsJSON, err := os.Open(config.Config.CardsPath)
	if err != nil {
		return err
	}
	byteValue, err := ioutil.ReadAll(cardsJSON)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteValue, &JSONCards); err != nil {
		return err
	}

	d.CardsList = make(map[string]Card, len(JSONCards.Cards)) //eventually change the string key to ints... GC doesn't like string keys
	d.NameToID = make(map[string]string, len(JSONCards.Cards))

	for _, v := range JSONCards.Cards {
		d.CardsList[v.ID] = v
		d.NameToID[v.Name] = v.ID //so we can pass in a card as an ID or name
	}
	return nil
}

func (d *MemoryDatastore) GetCardByIDOrName(id string) (Card, error) {
	card := Card{}
	if val, exists := d.CardsList[id]; exists {
		card = val
	}
	if val, exists := d.NameToID[id]; exists {
		card = d.CardsList[val]
	}

	if (card.ID == Card{}.ID) {
		return Card{}, errors.New("Card does not exist")
	}

	return card, nil
}
