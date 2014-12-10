package main

import (
	"github.com/HugoMFernandes/go-fifteen-puzzle/game"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/io"
	"flag"
	"github.com/nsf/termbox-go"
)

const defaultPuzzleWidth = 4
const defaultPuzzleHeight = 4

func main() {

	// Parse dimension arguments from command args
	puzzleWidth, puzzleHeight := parseDimensionArguments()

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// Initialize input handler
	input := io.CreateInputHandler()

	// Initialize game renderer
	renderer := io.CreateRenderer()

	// Render welcome message
	renderer.RenderWelcomeMessage()
	input.ReadKey()
	renderer.ClearScreen()

	// Start game
	game := game.CreateGame(puzzleWidth, puzzleHeight, renderer, input)
	game.Run()

	if game.Puzzle().IsSolved() {
		// Render victory message
		renderer.RenderVictoryMessage(game.Puzzle(), game.Stats())
		input.ReadKey()
	}

	// Clear the screen before we quit
	renderer.ClearScreen()
}

func parseDimensionArguments() (int, int) {
	puzzleWidthPtr := flag.Int("w", defaultPuzzleWidth, "a puzzle width")
	puzzleHeightPtr := flag.Int("h", defaultPuzzleHeight, "a puzzle height")

	flag.Parse()

	return *puzzleWidthPtr, *puzzleHeightPtr
}
