/**
 * Application routines
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

package main

import (
	"errors"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

/**
 * TODO:
 * Add type Application struct { ... }
 * and var app Application and this stuff
 * inside it
 */

/// Is application running
var isRunning bool

/// SDL Window
var window *sdl.Window

/// SDL Renderer
var rend *sdl.Renderer

/// Canvas
var canvas *sdl.Texture

/// Canvas drawing size
var canvasSize vec2int

/// Canvas position
var canvasPos vec2int

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

// Game scene
var sceneGame scene

// Current scene
var currentScene scene

// Assets
var ass assets

// Cursor
var bmpCursor bitmap

/// FRAMEWAIT Frame wait constant (= 1.0 / FRAMELIMIT)
const FRAMEWAIT = 17

/**
 * Set an error message
 *
 * Params:
 * msg Error message
 */
func setErr(msg string) {
	err = errors.New(msg)
}

/**
 * Get renderer
 *
 * Returns:
 * Renderer
 */
func getRend() *sdl.Renderer {
	return rend
}

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

	// Create render target texture aka canvas
	canvas, err = rend.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 320, 240)
	if err != nil {
		return 1
	}
	windowWidth, windowHeight := window.GetSize()
	calcCanvasPosAndSize(windowWidth, windowHeight)

	// Init IMG addon
	img.Init(img.INIT_PNG)

	// Init audio
	mix.Init(mix.INIT_OGG)
	err = mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 1024)
	if err != nil {
		return 1
	}

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

	// Load assets
	// (Normally this would happen in external thread, but no time for that)
	var errCode int
	ass, errCode = loadAssets(rend)
	if errCode == 1 {
		return 1
	}

	bmpCursor = ass.getBitmap("cursor")

	// Create game scene and initilize it
	sceneGame = gameGetScene()
	if sceneGame.onInit != nil {
		if sceneGame.onInit(ass) == 1 {
			return 1
		}
	}

	// Set game the current scene
	currentScene = sceneGame

	// Set music playing
	mix.VolumeMusic(72)
	ass.music.FadeIn(-1, 1000)

	// Read score
	readBest("hi.score")

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

		case *sdl.MouseButtonEvent:
			if t.Type == sdl.MOUSEBUTTONDOWN {
				onMouseDown(uint(t.Button))
			} else if t.Type == sdl.MOUSEBUTTONUP {
				onMouseUp(uint(t.Button))
			}

		case *sdl.MouseMotionEvent:
			onMouseMove(t.X, t.Y)

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_RESIZED {
				calcCanvasPosAndSize(int(t.Data1), int(t.Data2))
			}
		}
	}

}

/**
 * Calculate canvas rendering position and size
 */
func calcCanvasPosAndSize(windowWidth, windowHeight int) {

	/*
	 * Determine canvas position and size by the window size
	 * and the aspect ratio
	 */
	if float32(windowWidth)/float32(windowHeight) >= 320.0/240.0 {
		canvasSize.x = int(float32(windowHeight) / 240.0 * 320.0)
		canvasPos.x = windowWidth/2.0 - canvasSize.x/2.0
		canvasPos.y = 0.0

		canvasSize.y = windowHeight
	} else {
		canvasSize.y = int(float32(windowWidth) / 320.0 * 240.0)
		canvasPos.y = windowHeight/2.0 - canvasSize.y/2.0
		canvasPos.x = 0.0

		canvasSize.x = windowWidth
	}
}

/**
 * Update frame
 *
 * Params:
 * deltaTime Time passed since the previous frame
 */
func update(deltaTime uint32) {

	// Set cursor vpos before anything else
	w, h := window.GetSize()
	setCursorVpos(vec2int{x: int(w), y: int(h)}, vec2int{x: 320, y: 240}, canvasPos)

	// Calculate time multiplier (1.0 if fps = 60, 2.0 if fps = 30 etc)
	timeMul := float32(deltaTime) / 1000.0 / (1.0 / 60.0)
	// Limit time multiplier to avoid unexcepted jumps
	if timeMul <= 0.0 {
		timeMul = 1.0
	} else if timeMul > 5.0 {
		timeMul = 5.0
	}

	// Leave if Ctrl+Q pressed (no dialogue)
	if getKeyState(sdl.SCANCODE_LCTRL) == StateDown &&
		getKeyState(sdl.SCANCODE_Q) == StatePressed {
		isRunning = false
	}

	// If F4 is pressed, go to the full screen mode
	if getKeyState(sdl.SCANCODE_F4) == StatePressed {
		toggleFullscreen()
	}

	// Update current scene, if update function defined
	if currentScene.onUpdate != nil {
		currentScene.onUpdate(timeMul)
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
 * Draw to the canvas
 */
func drawToCanvas() {

	rend.SetRenderTarget(canvas)

	// Draw current scene, if drawing function defined
	if currentScene.onDraw != nil {
		currentScene.onDraw(rend)
	}

	// Draw cursor (always visible)
	cpos := getCursorPos()
	drawBitmap(bmpCursor, int32(cpos.x), int32(cpos.y))

	rend.SetRenderTarget(nil)
}

/**
 * Draw a frame
 *
 * Params:
 * null
 */
func draw() {

	// Clear screen
	rend.SetDrawColor(0, 0, 0, 255)
	rend.Clear()

	// Draw to canvas
	drawToCanvas()

	// Draw canvas
	rend.Copy(canvas, &sdl.Rect{X: 0, Y: 0, W: 320, H: 240},
		&sdl.Rect{X: int32(canvasPos.x), Y: int32(canvasPos.y), W: int32(canvasSize.x), H: int32(canvasSize.y)})

	// Refresh screen
	rend.Present()

}

/**
 * Destroys the application
 */
func destroy() {

	// Write best score
	storeBest("hi.score")

	// Destroy all the scenes
	if sceneGame.onDestroy != nil {
		sceneGame.onDestroy()
	}

	// Destroy assets
	destroyAssets(ass)

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
		update(deltaTime)
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

	// Destroy
	destroy()

	return 0
}
