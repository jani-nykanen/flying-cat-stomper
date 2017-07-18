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

/**
 * Initialize the game scene
 *
 * Returns:
 * 0 on success, 1 on error
 */
func gameInit(ass assets) int {

	// Initialize background
	bg.init(ass)

	// Init hud
	hud.init(ass)

	return 0
}

/**
 * Updates the game scene
 *
 * Params:
 * timeMul Time multiplier
 */
func gameUpdate(timeMul float32) {

	// Update background
	bg.update(timeMul, 1.0)

	// Update HUD
	hud.update(timeMul)

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

	// Draw HUD
	hud.draw(rend)
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
