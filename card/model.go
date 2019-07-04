package card

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Cards struct {
	DAL DAL
}

//Use a datastore that holds all of the card data in memory. Later on we can implement a DBDatastore that calls from a DB or something
type MemoryDatastore struct {
	CardsList map[string]Card //ID -> card
	NameToID  map[string]string
}

func NewCards(dal DAL) Cards {
	return Cards{DAL: dal}
}

//This is all fine for Vol. 1 cards. Later on it gets complicated, but I will add that as we go. Definitely want to start small here
// Honestly not sure how I want to store all the cards. For now it seems easiest to get some JSON setup then try and pull from some API somewhere :shrug:

type Card struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	CardType    CardType     `json:"type"`
	Attribute   Attribute    `json:"attribute"`
	Race        Type         `json:"race"`
	ATK         string       `json:"atk"`
	DEF         string       `json:"def"`
	Level       string       `json:"level"`
	Description string       `json:"desc"`
	CardImages  []CardImages `json:"card_images"`
}

func (c Card) String() string {
	return fmt.Sprintf("%s (%v)", c.Name, c.ID)
}

//The Yu Gi Oh API returns a lot of different pics
type CardImages struct {
	ID            string `json:"id"`
	ImageURL      string `json:"image_url"`
	ImageURLSmall string `json:"image_url_small"`
}

//ENUMS - kind of a bummer that this style of enum does not _just work TM_ with JSON without adding maps back and forth with a new marshal and unmarshal......

//CardType is the largest abstraction of the card. Boils down to the three major cards of vol. 1
type CardType int

const (
	MONSTER CardType = iota
	SPELL
	TRAP
)

//Attribute of the "being" , a lot of vol. 1 is DARK or EARTH
type Attribute int

const (
	EARTH Attribute = iota
	DARK
	WIND
	LIGHT
	WATER
)

//Type of the actual "being" in the card. IE a T-Rex is a dinosaur
type Type int

const (
	NORMAL Type = iota
	AQUA
	BEAST
	BEAST_WARRIOR
	CREATOR_GOD
	CYBERSE
	DINOSAUR
	DIVINE_BEAST
	DRAGON
	EQUIP
	FAIRY
	FIEND
	FISH
	INSECT
	MACHINE
	PLANT
	PSYCHIC
	PYRO
	REPTILE
	ROCK
	SEA_SERPENT
	SPELLCASTER
	THUNDER
	WARRIOR
	WINGED_BEAST
	WYRM
	ZOMBIE
	UNKNOWN // seen as ??? on cards
)

//This is taken from https://code.tutsplus.com/tutorials/json-serialization-with-golang--cms-30209
//Not pretty and kind of obnoxious to write for each enum, but its fine for now - until I figure out how I wanna pull the card data

//I think actually pulling from an API will be a good idea in the future, but definitely not now since I wanna mutate the data around and see
//	what works best

func (ct *CardType) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	value, ok := map[string]CardType{"Normal Monster": MONSTER, "Spell Card": SPELL, "Trap Card": TRAP}[s]
	if !ok {
		return errors.New("Invalid CardType value")
	}
	*ct = value
	return nil
}

func (atr *Attribute) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	value, ok := map[string]Attribute{"EARTH": EARTH, "DARK": DARK, "WIND": WIND, "LIGHT": LIGHT, "WATER": WATER}[s]
	if !ok {
		return errors.New("Invalid EnumType value")
	}
	*atr = value
	return nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	//This is pretty nasty feeling - all this for some "type safety" with these "enums" lol
	value, ok := map[string]Type{
		"Normal":        NORMAL,
		"Aqua":          AQUA,
		"Beast":         BEAST,
		"Beast-Warrior": BEAST_WARRIOR,
		"Creator-God":   CREATOR_GOD,
		"Cyberse":       CYBERSE,
		"Dinosaur":      DINOSAUR,
		"Divine Beast":  DIVINE_BEAST,
		"Dragon":        DRAGON,
		"Equip":         EQUIP,
		"Fairy":         FAIRY,
		"Fiend":         FIEND,
		"Fish":          FISH,
		"Insect":        INSECT,
		"Machine":       MACHINE,
		"Plant":         PLANT,
		"Psychic":       PSYCHIC,
		"Pyro":          PYRO,
		"Reptile":       REPTILE,
		"Rock":          ROCK,
		"Sea Serpent":   SEA_SERPENT,
		"Spellcaster":   SPELLCASTER,
		"Thunder":       THUNDER,
		"Warrior":       WARRIOR,
		"Winged Beast":  WINGED_BEAST,
		"Wyrm":          WYRM,
		"Zombie":        ZOMBIE,
		"Unknown":       UNKNOWN,
	}[s]
	if !ok {
		return errors.New("Invalid EnumType value")
	}
	*t = value
	return nil
}
