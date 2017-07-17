/**
 * Game scene
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

package main

import (
	"math"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

/// Font bitmap
var bmpFont bitmap

/// Framerate
var fps int

/// FPS sum
var fpsSum int

/// FPS count
var fpsCount int

/**
 * Initialize the game scene
 *
 * Returns:
 * 0 on success, 1 on error
 */
func gameInit(ass assets) int {

	bmpFont = ass.getBitmap("font")

	fpsCount = 0
	fpsSum = 0
	fps = 60

	// Initialize background
	bgInit(ass)

	return 0
}

/**
 * Updates the game scene
 *
 * Params:
 * timeMul Time multiplier
 */
func gameUpdate(timeMul float32) {

	// Since timeMul varies from frame to frame, let's take the
	// average in one second (or actually, a time frame that is at
	// least a second )
	fpsSum += int(math.Floor(60.0 / float64(timeMul)))
	fpsCount++
	if fpsCount >= 60 {
		fps = int(float32(fpsSum) / float32(fpsCount))
		fpsCount = 0
		fpsSum = 0
	}

	// Update background
	bgUpdate(timeMul, 1.0)

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
	bgDraw(rend)

	drawText(bmpFont, "Hello world!", 2, 2)
	drawText(bmpFont, "FPS: "+strconv.Itoa(fps), 2, 18)
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
