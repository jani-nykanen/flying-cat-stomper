/**
 * Controls (keyboard + mouse)
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

///
/// TODO: Add 'type controls struct' !
///

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

/// Key states
const (
	StateUp       = 0
	StateDown     = 1
	StatePressed  = 2
	StateReleased = 3

	MaxKey    = 256
	MaxButton = 16
)

/// Keys
var keystate [MaxKey]int

/// Mouse buttons
var buttonstate [MaxButton]int

/// Cursor pos
var cursorPos vec2int

/// Virtual cursor pos
var vpos vec2int

/**
 * Init controls
 */
func initControls() {
	for i := 0; i < len(keystate); i++ {
		keystate[i] = StateUp
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
	if key >= MaxKey || keystate[key] == StateDown {
		return
	}
	// Set key value to pressed
	keystate[key] = StatePressed
}

/**
 * On key up
 *
 * Params:
 * key Key scancode
 */
func onKeyUp(key uint) {
	// If key is out of range or already up, ignore
	if key >= MaxKey || keystate[key] == StateUp {
		return
	}
	// Set key value to released
	keystate[key] = StateReleased
}

/**
 * Triggers when mouse button is pressed down
 *
 * Params:
 * button Mouse button
 */
func onMouseDown(button uint) {
	if button >= MaxButton {
		return
	}

	buttonstate[button] = StatePressed
}

/**
 * Triggers when mouse button is released
 *
 * Params:
 * button Mouse button
 */
func onMouseUp(button uint) {
	if button >= MaxButton {
		return
	}

	buttonstate[button] = StateReleased
}

/**
 * Triggers when mouse is moved
 *
 * Params:
 * x Mouse X
 * y Mouse Y
 */
func onMouseMove(x, y int32) {
	cursorPos.x = int(x)
	cursorPos.y = int(y)
}

/**
 * Update controls
 */
func updateControls() {
	for i := 0; i < MaxKey; i++ {
		// Check keyboard keys
		if keystate[i] == StatePressed {
			keystate[i] = StateDown

		} else if keystate[i] == StateReleased {
			keystate[i] = StateUp
		}

		// Check mouse buttons
		if i < MaxButton {
			if buttonstate[i] == StatePressed {
				buttonstate[i] = StateDown

			} else if buttonstate[i] == StateReleased {
				buttonstate[i] = StateUp
			}
		}
	}
}

/**
 * Set virtual cursor position, so it's always in the canvas
 *
 * Params:
 * winSize Window size
 * canvasSize Canvas size (actual size, not scaled)
 * canvasPos Canvas position on screen
 */
func setCursorVpos(winSize vec2int, canvasSize vec2int, canvasPos vec2int) {
	vpos.x = int(float32(cursorPos.x-canvasPos.x) / float32(winSize.x-canvasPos.x*2) * float32(canvasSize.x))
	vpos.y = int(float32(cursorPos.y-canvasPos.y) / float32(winSize.y-canvasPos.y*2) * float32(canvasSize.y))
}

/**
 * Get the key state
 *
 * Params:
 * scancode Scancode of the desired key
 */
func getKeyState(scancode sdl.Scancode) int {
	key := uint(scancode)
	if key >= MaxKey {
		return StateUp
	}

	return keystate[key]
}

/**
 * Get the mouse button state
 *
 * Params:
 * scancode Button index
 */
func getMouseButtonState(button uint) int {
	if button >= MaxButton {
		return StateUp
	}

	return buttonstate[button]
}

/**
 * Get the virtual cursor position
 *
 * Returns:
 * Cursor virtual position
 */
func getCursorPos() vec2int {
	return vpos
}
