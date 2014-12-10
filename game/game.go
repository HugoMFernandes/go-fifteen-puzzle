package game

import (
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/constants"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/io"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/puzzle"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/stats"
)

// Game structure
type Game struct {
	puzzle   *puzzle.Puzzle
	renderer *io.Renderer
	input    *io.InputHandler
	stats    *stats.StatsMonitor
}

// Getters for printing the victory message
func (g *Game) Puzzle() *puzzle.Puzzle {
	return g.puzzle
}

func (g *Game) Stats() stats.Stats {
	return g.stats.Stats()
}

// Game constructor
func CreateGame(puzzleWidth int, puzzleHeight int, renderer *io.Renderer, input *io.InputHandler) *Game {
	// Create puzzle and stats
	puzzle := puzzle.CreatePuzzle(puzzleWidth, puzzleHeight)
	stats := stats.CreateStatsMonitor()

	return &Game{
		puzzle:   puzzle,
		renderer: renderer,
		input:    input,
		stats:    stats,
	}
}

// Game loop
func (g *Game) Run() {

	puzzle := g.puzzle
	renderer := g.renderer
	input := g.input
	stats := g.stats

	// Start game timer
	stats.StartTimer()

	// Render initial puzzle state on screen
	renderPuzzle(renderer, puzzle)

	// Play until puzzle is solved
	for !puzzle.IsSolved() {

		// Read input from user
		playerInput := input.ReadInputEvent()

		if playerInput == constants.INPUT_QUIT {
			// Quit key was pressed
			break
		}

		// Input corresponds to a sliding direction
		slidingDirection, ok := constants.InputEventToSlidingDirection[playerInput]

		// This should never happen unless we add new keys and forget to handle them
		if !ok {
			continue
		}

		// Try to slide piece in the requested direction
		if puzzle.Move(slidingDirection) {
			// Count move
			stats.CountMove()

			// Clear screen and render new puzzle state
			renderer.ClearScreen()
			renderPuzzle(renderer, puzzle)
		}
	}

	// Stop game timer
	stats.StopTimer()
}

func renderPuzzle(r *io.Renderer, p *puzzle.Puzzle) {
	r.RenderPuzzle(p)
}
