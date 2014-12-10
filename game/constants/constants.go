package constants

// Create custom events for key presses (so we don't leak lib structs)
type InputEvent int

const (
	INPUT_UP InputEvent = iota
	INPUT_RIGHT
	INPUT_DOWN
	INPUT_LEFT
	INPUT_QUIT
)

// Create enum with sliding directions
type SlidingDirection int

const (
	SLIDE_UP SlidingDirection = iota
	SLIDE_RIGHT
	SLIDE_DOWN
	SLIDE_LEFT
)

// Create array with possible sliding directions for scrambling
var POSSIBLE_SLIDING_DIRECTIONS = []SlidingDirection{
	SLIDE_UP, SLIDE_RIGHT, SLIDE_DOWN, SLIDE_LEFT,
}

// Map input events to sliding directions
var InputEventToSlidingDirection = map[InputEvent]SlidingDirection{
	INPUT_UP:    SLIDE_UP,
	INPUT_RIGHT: SLIDE_RIGHT,
	INPUT_DOWN:  SLIDE_DOWN,
	INPUT_LEFT:  SLIDE_LEFT,
}

// Messages
const WELCOME_MESSAGE = "Welcome to 15-puzzle!\n" +
	"\n" +
	"~~~~~ How to play ~~~~~\n" +
	"Use the arrow keys to slide tiles around\n" +
	"VI keys (h,j,k,l) also work\n" +
	"\n" +
	"~~~~~ How to win ~~~~~\n" +
	"Slide the tiles until all numbers are in (reading) order\n" +
	"Remember: tiles can only slide towards the empty space\n" +
	"Press Q (or Ctrl-C) if you want to give up\n" +
	"\n" +
	"~~~~~ Have fun! ~~~~~\n" +
	"(Press any key to start playing..)\n"

const VICTORY_MESSAGE = "Congratulations!\n" +
    "How about trying a larger puzzle?\n"

const EXIT_MESSAGE = "(Press any key to exit..)\n"
