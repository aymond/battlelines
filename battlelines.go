package main

//import "github.com/aymond/gocards"
import (
	"fmt"

	"github.com/aymond/gocards"
)

func main() {
	// Initialise game decks.
	// gameDeck is the set of unused shuffled cards.
	// discardDeck are cards removed from the game.
	// playerOneDeck and playerTwoDeck are the player hands.
	gameDeck, _ := (gocards.NewDeck("BattleLines"))
	tacticsDeck, _ := (gocards.NewDeck("Tactics"))
	discardDeck, _ := gocards.NewDeck("Discard")
	playerOneDeck, _ := gocards.NewDeck("PlayerOne")
	playerTwoDeck, _ := gocards.NewDeck("PlayerTwo")
	fmt.Println(gameDeck, tacticsDeck, playerOneDeck, playerTwoDeck, discardDeck)

	// Load the two decks.
	gameDeck.Initialize("data/battlelines.json")
	tacticsDeck.Initialize("data/tactics.json")

	// One for the Players, one for Tactics Cards
	// And initialise a Discard deck.

}

func loadDeck() {

}
