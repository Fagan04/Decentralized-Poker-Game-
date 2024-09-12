package deck

import (
	"fmt"
	"strconv"
)

type Suit int

const (
	Hearts Suit = iota
	Spades
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Hearts:
		return "HEARTS"
	case Spades:
		return "SPADES"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("Not a valid Suit")
	}
}

func (c Card) String() string {
	value := strconv.Itoa(c.Value)
	if c.Value == 1 {
		value = "ACE"
	} else if c.Value == 11 {
		value = "JACK"
	} else if c.Value == 12 {
		value = "QUEEN"
	} else if c.Value == 13 {
		value = "KING:"
	}

	return fmt.Sprintf("%s of %s %s", value, c.Suit, SuitToUnicode(c.Suit))
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("The value of the card cannot be higher than 13")
	}

	return Card{
		Suit:  s,
		Value: v,
	}
}

type Deck [52]Card

func NewDeck() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return d
}

type Card struct {
	Suit  Suit
	Value int
}

func SuitToUnicode(suit Suit) string {
	switch suit {
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("Not a valid Suit")
	}
}
