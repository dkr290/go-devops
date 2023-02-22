package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayCard struct {
	Suit string
	Rank string
}

type Deck[c any] struct {
	cards []c
}

type TradingCard struct {
	CollectableName string
}

func NewTradingCard(c string) *TradingCard {
	return &TradingCard{
		CollectableName: c,
	}
}

func (tc *TradingCard) String() string {
	return tc.CollectableName
}

func NewPlayCard(suit string, card string) *PlayCard {
	return &PlayCard{
		Suit: suit,
		Rank: card,
	}
}

func (pc *PlayCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

func NewPlayingCardDeck() *Deck[*PlayCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck[*PlayCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayCard(suit, rank))
		}
	}
	return deck
}

func NewTradigCardDeck() *Deck[*TradingCard] {
	collectables := []string{"Sammy", "Droplets", "Spaces", "App Platform"}

	deck := &Deck[*TradingCard]{}

	for _, c := range collectables {
		deck.AddCard(NewTradingCard(c))
	}

	return deck
}

func (d *Deck[c]) AddCard(card c) {
	d.cards = append(d.cards, card)
}

func (d *Deck[c]) RandomCard() c {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func main() {

	playingDeck := NewPlayingCardDeck()
	tradingDeck := NewTradigCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard := playingDeck.RandomCard()

	fmt.Printf("drew card: %s\n", playingCard)

	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)

	fmt.Printf("--- drawing trading card ---\n")

	tradingCard := tradingDeck.RandomCard()
	fmt.Printf("drew card: %s\n", tradingCard)
	fmt.Printf("card collectable name: %s\n", tradingCard.CollectableName)

}
