package io

import (
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/constants"
	"github.com/HugoMFernandes/go-fifteen-puzzle/game/settings"
	"github.com/nsf/termbox-go"
)

type InputHandler struct{}

func CreateInputHandler() *InputHandler {
	return &InputHandler{}
}

// We are technically leaking a termbox struct out of this package if someone uses it,
// but it doesn't really matter in this case (it won't even be used apart from "continues"
func (in *InputHandler) ReadKey() termbox.Event {
	for {
		keyEvent := termbox.PollEvent()

		// Only return if an actual key was found (ignore resizing events, etc)
		// We could capture other events and handle them (e.g. resize the terminal in
		// real-time, but that's too out of the scope of this program
		if keyEvent.Type == termbox.EventKey {
			return keyEvent
		}
	}
}

func (in *InputHandler) ReadInputEvent() constants.InputEvent {

	for {
		keyEvent := in.ReadKey()

		// This could be extensible, but a Vim keymap + normal arrows is fine
		switch {
		case keyEvent.Key == termbox.KeyCtrlC || keyEvent.Ch == settings.QUIT_KEY_CHAR:
			return constants.INPUT_QUIT
		case keyEvent.Key == termbox.KeyArrowUp || keyEvent.Ch == settings.UP_KEY_CHAR:
			return constants.INPUT_UP
		case keyEvent.Key == termbox.KeyArrowDown || keyEvent.Ch == settings.DOWN_KEY_CHAR:
			return constants.INPUT_DOWN
		case keyEvent.Key == termbox.KeyArrowRight || keyEvent.Ch == settings.RIGHT_KEY_CHAR:
			return constants.INPUT_RIGHT
		case keyEvent.Key == termbox.KeyArrowLeft || keyEvent.Ch == settings.LEFT_KEY_CHAR:
			return constants.INPUT_LEFT
		default:
			// Do nothing (i.e. loop again)
		}
	}
}
