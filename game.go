/**
 * Game scene
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

package main

import "github.com/veandco/go-sdl2/sdl"

/// Font bitmap
var bmpFont bitmap

/**
 * Initialize the game scene
 *
 * Returns:
 * 0 on success, 1 on error
 */
func gameInit(ass assets) int {

	bmpFont = getBitmap(ass, "font")

	return 0
}

/**
 * Updates the game scene
 *
 * Params:
 * timeMul Time multiplier
 */
func gameUpdate(timeMul float32) {

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

	drawBitmap(bmpFont, 0, 0)
	drawBitmapRegion(bmpFont,64,64,16,16,256,128)

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
