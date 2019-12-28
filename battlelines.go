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
	gameDeck, _ := gocards.NewDeck("BattleLines")
	tacticsDeck, _ := gocards.NewDeck("Tactics")
	discardDeck, _ := gocards.NewDeck("Discard")
	playerOneDeck, _ := gocards.NewDeck("PlayerOne")
	playerTwoDeck, _ := gocards.NewDeck("PlayerTwo")
	fmt.Println(gameDeck, tacticsDeck, playerOneDeck, playerTwoDeck, discardDeck)

	// Load the contents of the two game decks.
	// We are passing in S if Standard Deck json struture.
	// We are passing in f if Fixed Decj json structure.
	gameDeck.Initialize("data/battlelines.json", "s") // Standard Deck
	tacticsDeck.Initialize("data/tactics.json", "f")  // Fixed Deck

	fmt.Println(gameDeck)
	fmt.Println(tacticsDeck)
}

func loadDeck() {

}
