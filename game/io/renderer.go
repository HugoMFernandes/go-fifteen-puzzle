package io

import (
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/constants"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/math"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/puzzle"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/stats"
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
	"strings"
)

// Constants
const SEPARATOR = " "
const EMPTY_SPACE = " "
const LINE_CHAR = "-"
const VERTICAL_BAR = "|"

const TILE_PADDING = 1

const DEFAULT_COLOR = termbox.ColorDefault

// Actual renderer class
// Note: Since termbox is being used for this, we could actually not store this and just
// obtain it from termbox on every redraw. We'll assume no one will really be resizing this
// while playing though, as even resizing on every redraw would not be good enough (i.e. it would
// still require a redraw call
type Renderer struct {
	screenWidth  int
	screenHeight int
}

func CreateRenderer() *Renderer {

	// Get terminal dimensions from termbox
	terminalWidth, terminalHeight := termbox.Size()

	return &Renderer{
		screenWidth:  terminalWidth,
		screenHeight: terminalHeight,
	}
}

func (r *Renderer) RenderWelcomeMessage() {
	r.ClearScreen()
	r.renderTextBlock(constants.WELCOME_MESSAGE)
	r.FlushScreen()
}

func (r *Renderer) RenderVictoryMessage(puzzle *puzzle.Puzzle, stats stats.Stats) {
	r.ClearScreen()

	// Create the actual message
	statsRepresentation := generateStatsRepresentation(stats)
	puzzleRepresentation := generatePuzzleRepresentation(puzzle)
	finalMessage := fmt.Sprintf("%s\n%s\n%s\n%s", constants.VICTORY_MESSAGE, puzzleRepresentation, statsRepresentation, constants.EXIT_MESSAGE)

	r.renderTextBlock(finalMessage)
	r.FlushScreen()
}

func generateStatsRepresentation(stats stats.Stats) (ret string) {
	return fmt.Sprintf("Solve time: %s\nMove count: %d",
		stats.ElapsedTime(), stats.MoveCount())
}

// Note: Using a single function that renders text blocks is a peculiar way to do this,
// but since we just want to draw text anyway, it's actually a cool way (it enables us to
// use one single function to print welcome messages, as well as rendering a board
func (r *Renderer) renderTextBlock(textBlock string) {
	// Break block into lines
	textLines := strings.Split(textBlock, "\n")

	// Calculate paddings
	textBlockHeight := len(textLines)
	totalVerticalPadding := r.screenHeight - textBlockHeight
	verticalPadding := totalVerticalPadding / 2

	// Actual text
	renderCenteredLines(verticalPadding, textLines, r.screenWidth)
}

func renderCenteredLines(verticalPadding int, lines []string, screenWidth int) {
	for i, line := range lines {
		// Calculate horizontal padding for this line (based on length)
		totalHorizontalPadding := screenWidth - len(line)
		horizontalPadding := totalHorizontalPadding / 2

		// Calculate vertical padding for this line (based on index)
		lineVerticalPadding := verticalPadding + i
		renderText(lineVerticalPadding, horizontalPadding, line)
	}
}

// Note: This only renders a single line (i.e. \n's will not push y further down)
func renderText(y int, x int, text string) {
	for i, r := range text {
		// Calculate rune position from index and start of line
		runeX := x + i
		renderRune(y, runeX, r)
	}
}

func renderRune(y int, x int, r rune) {
	termbox.SetCell(x, y, r, DEFAULT_COLOR, DEFAULT_COLOR)
}

func (r *Renderer) ClearScreen() {
	termbox.Clear(DEFAULT_COLOR, DEFAULT_COLOR)
	r.FlushScreen()
}

func (r *Renderer) FlushScreen() {
	termbox.Flush()
}

// Puzzle stuff
func (r *Renderer) RenderPuzzle(p *puzzle.Puzzle) {
	r.ClearScreen()

	r.renderTextBlock(generatePuzzleRepresentation(p))
	r.FlushScreen()
}

func generatePuzzleRepresentation(p *puzzle.Puzzle) (ret string) {

	// Calculate spacings and paddings
	puzzleWidth := p.Width()
	puzzleHeight := p.Height()

	// Calculate largest number found on the puzzle
	largestNumber := puzzleHeight*puzzleWidth - 1

	// Calculate how many digits this number occupies
	maxNumLen := math.NumDigits(largestNumber)

	// Calculate how much padding is in a tile
	tilePadding := maxNumLen + TILE_PADDING*2
	lineLen := (tilePadding+len(VERTICAL_BAR))*puzzleWidth + 1

	// Top line
	ret += generateHorizontalLine(lineLen)

	for y := range p.Tiles() {
		for x := range p.Tiles()[y] {
			ret += VERTICAL_BAR
			ret += generateNumber(p.Tiles()[y][x], tilePadding)
		}

		// Vertical separator
		ret += VERTICAL_BAR
		ret += "\n"

		// Interim line
		ret += generateHorizontalLine(lineLen)
	}

	return
}

func generateHorizontalLine(lineLen int) string {
	return generateHorizontalLineOfChar(LINE_CHAR, lineLen)
}

func generateHorizontalLineOfChar(ch string, lineLen int) string {
	return strings.Repeat(ch, lineLen) + "\n"
}

func generateNumber(num int, tilePadding int) (ret string) {
	// Calculate how many digits this number occupies
	numLen := math.NumDigits(num)

	// Calculate and generate padding for this tile
	totalPadding := tilePadding - numLen
	padding := strings.Repeat(SEPARATOR, totalPadding/2)

	// Add additional padding on uneven cases
	if totalPadding%2 != 0 {
		ret += SEPARATOR
	}

	// Print left padding
	ret += padding

	// Print number or empty space, if the empty tile was found
	if num != 0 {
		ret += strconv.Itoa(num)
	} else {
		ret += EMPTY_SPACE
	}

	// Print right padding
	ret += padding

	return
}
