/**
 * Game scene
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

/// Global speed
var globalSpeed float32

/// Speed up timer
var speedUpTimer float32

// Is paused
var paused bool

/**
 * Initialize the game scene
 *
 * Returns:
 * 0 on success, 1 on error
 */
func gameInit(ass assets) int {

	// Set global speed
	globalSpeed = 1.5

	// Initialize background
	bg.init(ass)

	// Init hud
	hud.init(ass)

	// Init game objects
	gobj.init(ass)

	// Init status
	initStatus()

	// Not paused
	paused = false

	return 0
}

/**
 * Updates the game scene
 *
 * Params:
 * timeMul Time multiplier
 */
func gameUpdate(timeMul float32) {

	// Update speed
	speedUpTimer += 1.0 * timeMul
	if speedUpTimer >= 60.0*7.5 {
		speedUpTimer -= 60.0 * 7.5
		globalSpeed += 0.15
	}

	if !hud.showTitle && hud.titleGoaway <= 0.0 && getKeyState(sdl.SCANCODE_RETURN) == StatePressed {
		paused = !paused
	}

	// If scene is paused, do not update objects & events
	if paused {
		return
	}

	// Update background
	bg.update(timeMul, globalSpeed)

	if !hud.showTitle {
		// Update game objects
		gobj.update(timeMul, globalSpeed)
	}

	// Update HUD
	hud.update(timeMul)

	// Update best score if needed
	if status.score > status.best {
		status.best = status.score
	}
}

/**
 * Draws the game scene
 *
 * Params:
 * rend Renderer
 */
func gameDraw(rend *sdl.Renderer) {

	rend.SetDrawColor(170, 170, 170, 255)
	rend.Clear()

	// Draw background
	bg.draw(rend)

	// Draw game objects
	gobj.draw(rend)

	// Draw HUD
	hud.draw(rend)

	// Draw "Paused" if paused
	if paused {
		drawCenteredText(hud.bmpFontBig, "Game Paused", 160, 96)
	}
}

/**
 * Destroy the game scene
 *
 * Returns:
 * 0 always
 */
func gameDestroy() {

}

/**
 * Return a "list" of scene functions
 *
 * Returns:
 * Scene object
 */
func gameGetScene() scene {

	// Put scenes to a local scene object and return it
	s := scene{
		onInit:    gameInit,
		onUpdate:  gameUpdate,
		onDraw:    gameDraw,
		onDestroy: gameDestroy,
	}

	return s

}
