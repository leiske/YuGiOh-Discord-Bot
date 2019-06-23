package model

import (
	"encoding/json"
	"errors"
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
