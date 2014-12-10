package puzzle

import (
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/constants"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/math"
	"errors"
	"math/rand"
	"time"
)

// Init random generator
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Puzzle structure
type Puzzle struct {
	tiles  [][]int
	width  int
	height int
}

// Getters
func (p *Puzzle) Width() int {
	return p.width
}

func (p *Puzzle) Height() int {
	return p.height
}

// Constructor
func CreatePuzzle(width int, height int) *Puzzle {
	// Initialize slices for tiles. This could obviously be a 1D slice,
	// but it doesn't really matter
	tiles := make([][]int, height)
	for line := range tiles {
		tiles[line] = make([]int, width)
	}

	// Create puzzle
	puzzle := &Puzzle{
		tiles:  tiles,
		width:  width,
		height: height,
	}

	defer puzzle.init()

	return puzzle
}

func (p *Puzzle) init() {
	// Init solved puzzle first to ensure it's solvable
	numNumbers := p.height * p.width
	numbers := make([]int, numNumbers)

	// Start at 1 (0 is last)
	for i := 0; i < len(numbers)-1; i++ {
		numbers[i] = i + 1
	}

	// Initialize the puzzle tiles with the generated numbers. This could be much
	// smarter, but it's kind of irrelevant here
	for y := range p.tiles {
		for x := range p.tiles[y] {
			p.tiles[y][x] = numbers[y*p.width+x]
		}
	}

	// Scramble puzzle
	p.scramble()
}

func (p *Puzzle) scramble() {
	// Do some trivial scrambling (brute force is fine). God's number for 15-puzzle
	// is ~80 I think, so 20 ratio seems ok for any dimension. Naturally, the puzzle could
	// come in a solved state if we're really lucky (easier on smaller versions)
	numNumbers := p.height * p.width
	numberToScrambleLengthRatio := 20

	scrambleLen := numNumbers * numberToScrambleLengthRatio

	for i := 0; i < scrambleLen; i++ {
		p.slideRandomTile()
	}
}

func (p *Puzzle) slideRandomTile() {
	for {
		// Ugly hack but couldn't find any real enum support
		var randomSlide = constants.SlidingDirection(random.Intn(len(constants.POSSIBLE_SLIDING_DIRECTIONS)))

		if p.Move(randomSlide) {
			return
		}
	}
}

func (p *Puzzle) Tiles() [][]int {
	return p.tiles
}

func (p *Puzzle) IsSolved() bool {
	// Start at 1 and increment, always expecting (1, 2, ..., height*width - 1)
	expectedNumber := 1

	finalHeight := p.height - 1
	finalWidth := p.width - 1

	for y := range p.tiles {
		for x := range p.tiles[y] {
			// Last piece is a 0
			if y == finalHeight && x == finalWidth {
				return true
			}

			// Break if a number is out of order
			if p.tiles[y][x] != expectedNumber {
				return false
			}

			expectedNumber++
		}
	}

	return true
}

func (p *Puzzle) Move(input constants.SlidingDirection) bool {
	var direction math.Vertex2

	// SlidingDirection -> direction vertex
	// Note: The directions are reversed, because a slide is actually seen here
	// as a "free tile" slide (i.e. it's the free tile that moves into the actual tile,
	// and not the other way around
	switch input {
	case constants.SLIDE_UP:
		direction.X = 0
		direction.Y = 1
	case constants.SLIDE_RIGHT:
		direction.X = -1
		direction.Y = 0
	case constants.SLIDE_DOWN:
		direction.X = 0
		direction.Y = -1
	case constants.SLIDE_LEFT:
		direction.X = 1
		direction.Y = 0
	}

	return p.slidePieces(direction)
}

func (p *Puzzle) slidePieces(direction math.Vertex2) bool {
	// This could obviously be initialized on puzzle creation, cached and updated,
	// but I'll keep it simple as performance is irrelevant here
	emptyTile, err := p.findEmptyTile()

	if err != nil {
		// In a normal flow this will never happen, so let's assume whatever caused it
		// will handle itself from here (and return false, so we don't break)
		return false
	}

	targetTile := emptyTile.Add(direction)

	if p.tileWithinBounds(targetTile) {
		p.swapTiles(emptyTile, targetTile)

		return true
	}

	// No pieces were moved
	return false
}

func (p *Puzzle) tileWithinBounds(pos math.Vertex2) bool {
	w := p.width
	h := p.height

	// This should obviously be abstracted in a more complex game, in order to prevent
	// code duplication. This is literally the only place we need this in this game, though
	return pos.X >= 0 && pos.Y >= 0 && pos.X < w && pos.Y < h
}

func (p *Puzzle) swapTiles(pos1 math.Vertex2, pos2 math.Vertex2) {
	// Multiple return values are pretty cool :)
	p.tiles[pos1.Y][pos1.X], p.tiles[pos2.Y][pos2.X] = p.tiles[pos2.Y][pos2.X], p.tiles[pos1.Y][pos1.X]
}

func (p *Puzzle) findEmptyTile() (math.Vertex2, error) {
	for y := range p.tiles {
		for x, tile := range p.tiles[y] {
			if tile == 0 {
				return math.Vertex2{
					X: x,
					Y: y,
				}, nil
			}
		}
	}

	// This should never happen
	return math.Vertex2{}, errors.New("empty tile not found")
}
