package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

/// Is application running
var isRunning bool

/// SDL Window
var window *sdl.Window

/// SDL Renderer
var rend *sdl.Renderer

/// Ticks, old
var oldTicks uint32

/// Ticks, new
var newTicks uint32

/// Delta time
var deltaTime uint32

/// Error
var err error

/// Event
var event sdl.Event

/// Is fullscreen
var isFullscreen bool

/// FRAMEWAIT Frame wait constant (= 1.0 / FRAMELIMIT)
const FRAMEWAIT = 17

/**
 * Initialize SDL
 *
 * Returns:
 * 0 on success
 */
func initSDL() int {

	// Init everything (todo: remove things we truly do not need)
	sdl.Init(sdl.INIT_EVERYTHING)

	var err error

	// Create a window
	window, err = sdl.CreateWindow("Game", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 640, 480, sdl.WINDOW_RESIZABLE)
	if err != nil {
		return 1
	}

	// Not yet full screen
	isFullscreen = false

	// Hide cursor
	sdl.ShowCursor(0)

	// Create renderer
	rend, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return 1
	}

	// Clear the screen to black before initializing anything else
	rend.SetDrawColor(0, 0, 0, 255)
	rend.Clear()
	rend.Present()

	return 0
}

/**
 * Initializes application
 *
 * Returns:
 * 0 on success, 1 otherwise
 */
func initialize() int {

	// Init SDL
	if initSDL() == 1 {
		return 1
	}

	// Init controls
	initControls()

	// Set to running
	isRunning = true

	return 0
}

/**
 * Check SDL events
 */
func events() {
	// Go through every event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch t := event.(type) {

		case *sdl.QuitEvent:
			isRunning = false

		case *sdl.KeyDownEvent:
			onKeyDown(uint(t.Keysym.Scancode))

		case *sdl.KeyUpEvent:
			onKeyUp(uint(t.Keysym.Scancode))
		}
	}

}

/**
 * Update frame
 *
 * Params:
 * deltaTime Time passed since the previous frame
 */
func update(deltaTime float64) {

	// Leave if Ctrl+Q pressed (no dialogue)
	if getKeyState(sdl.SCANCODE_LCTRL) == STATE_DOWN &&
		getKeyState(sdl.SCANCODE_Q) == STATE_PRESSED {
		isRunning = false
	}

	// If Alt+Enter pressed, go to the full screen mode
	if getKeyState(sdl.SCANCODE_LALT) == STATE_DOWN &&
		getKeyState(sdl.SCANCODE_RETURN) == STATE_PRESSED {
		toggleFullscreen()
	}

	// Update controls in the end of the frame
	updateControls()
}

/**
 * Toggle fullscreen mode
 */
func toggleFullscreen() {
	flag := uint32(0)
	if isFullscreen {
		flag = 0
	} else {
		flag = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	window.SetFullscreen(flag)
	isFullscreen = !isFullscreen
}

/**
 * Draw a frame
 *
 * Params:
 * null
 */
func draw() {

	// Clear screen
	rend.SetDrawColor(170, 170, 170, 255)
	rend.Clear()

	// Refresh screen
	rend.Present()

}

/**
 * Destroys the application
 */
func destroy() {
	window.Destroy()
	rend.Destroy()
}

/**
 * Runs application
 *
 * Returns:
 * 0 on success, 1 otherwise
 */
func run() int {

	// Initialize
	if initialize() == 1 {
		sdl.ShowSimpleMessageBox(sdl.MESSAGEBOX_ERROR, "A fatal error occurred", err.Error(), window)
	}

	// Loop if running
	for isRunning {

		// Time before the frame events
		oldTicks = sdl.GetTicks()

		// Do frame events
		events()
		update(0.0)
		draw()

		// Time after the frame events
		newTicks = sdl.GetTicks()

		// Get rest time
		deltaMilliseconds := (newTicks - oldTicks)
		restTime := (FRAMEWAIT - 1) - deltaMilliseconds

		// Rest if needed
		if restTime > 0 {
			sdl.Delay(restTime)
		}

		// Set delta time
		deltaTime = sdl.GetTicks() - oldTicks
	}

	return 0
}
