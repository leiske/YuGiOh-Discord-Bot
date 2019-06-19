package model

//This is all fine for Vol. 1 cards. Later on it gets complicated, but I will add that as we go. Definitely want to start small here

type Card struct {
	Name      string
	CardType  CardType
	Attribute Attribute
	Types     []Type
	Property Property //used in spell cards
	CardEffectTypes []CardEffectType
	ATK       int
	DEF       int
	Level     int
	//unsure if I should include descriptions here, I think for now on this set of ~40 cards I will just get the core going and reeval later.
}

//ENUMS

type CardType int

const (
	MONSTER CardType = iota
	SPELL
	TRAP
)

type Property int

const (
	EQUIP Property = iota
	NORMAL_PROPERTY
)

type CardEffectType int

const (
	CONDITION Property = iota
	CONTINUOUS_LIKE
	EFFECT
	ACTIVATION_REQUIREMENT
)

type Attribute int

const (
	EARTH Attribute = iota
	DARK
	WIND
	LIGHT
	WATER
)

type Type int

const (
	NORMAL_TYPE Type = iota
	AQUA
	BEAST
	BEAST_WARRIOR
	CREATOR_GOD
	CYBERSE
	DINOSAUR
	DIVINE_BEAST
	DRAGON
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
