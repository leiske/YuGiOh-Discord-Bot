package model

import "fmt"

//This is all fine for Vol. 1 cards. Later on it gets complicated, but I will add that as we go. Definitely want to start small here
// Honestly not sure how I want to store all the cards. For now it seems easiest to get some JSON setup then try and pull from some API somewhere :shrug:
//I really don't want to rely on other APIs when the bot is running. This seems like something that could package up all the info in here.... I think the biggest thing id call out for is images

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
