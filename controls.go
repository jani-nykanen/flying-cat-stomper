package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

/// Key states
const (
	STATE_UP        = 0
	STATE_DOWN      = 1
	STATE_PRESSED   = 2
	STATED_RELEASED = 3

	MAX_KEY = 256
)

/// Keys
var keystate [MAX_KEY]int

/**
 * Init controls
 */
func initControls() {
	for i := 0; i < len(keystate); i++ {
		keystate[i] = STATE_UP
	}
}

/**
 * On key down
 *
 * Params:
 * key Key scancode
 */
func onKeyDown(key uint) {
	// If key is out of range or already down, ignore
	if key >= MAX_KEY || keystate[key] == STATE_DOWN {
		return
	}
	// Set key value to pressed
	keystate[key] = STATE_PRESSED
}

/**
 * On key up
 *
 * Params:
 * key Key scancode
 */
func onKeyUp(key uint) {
	// If key is out of range or already up, ignore
	if key >= MAX_KEY || keystate[key] == STATE_UP {
		return
	}
	// Set key value to released
	keystate[key] = STATED_RELEASED
}

/**
 * Update controls
 */
func updateControls() {
	for i := 0; i < MAX_KEY; i++ {
		if keystate[i] == STATE_PRESSED {
			keystate[i] = STATE_DOWN

		} else if keystate[i] == STATED_RELEASED {
			keystate[i] = STATE_UP
		}
	}
}

/**
 * Get the key state
 *
 * Params:
 * scancode Scancode of the desired key
 */
func getKeyState(scancode sdl.Scancode) int {
	key := uint(scancode)
	if key >= MAX_KEY {
		return STATE_UP
	}

	return keystate[key]
}
