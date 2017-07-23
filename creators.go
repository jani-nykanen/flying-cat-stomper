/**
 * Creators screen
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

/// Creators bitmap
var bmpCreators bitmap

/// Creators timer
var crTimer float32

/**
 * Initialize the creators scene
 *
 * Returns:
 * 0 on success, 1 on error
 */
func crInit(ass assets) int {

	bmpCreators = ass.getBitmap("creator")
	crTimer = -15.0

	return 0
}

/**
 * Updates the creators scene
 *
 * Params:
 * timeMul Time multiplier
 */
func crUpdate(timeMul float32) {

	crTimer += 1.0 * timeMul
	if crTimer > 240 {
		// Change scene
		currentScene = sceneGame

		// Set music playing
		mix.VolumeMusic(72)
		ass.music.FadeIn(-1, 1000)
	}

}

/**
 * Draws the creators scene
 *
 * Params:
 * rend Renderer
 */
func crDraw(rend *sdl.Renderer) {

	rend.SetDrawColor(0, 0, 0, 255)
	rend.Clear()

	xPos := float32(0)
	if crTimer < 30.0 {
		xPos = -320 + crTimer/30.0*320

	} else if crTimer >= 90.0 && crTimer < 120.0 {
		xPos = (crTimer - 90) / 30.0 * -320

	} else if crTimer >= 120.0 && crTimer < 150.0 {
		xPos = -320 + (crTimer-120)/30.0*320

	} else if crTimer >= 210.0 {
		xPos = (crTimer - 210) / 30.0 * -320
	}

	srcPos := int32(0)
	if crTimer >= 120 {
		srcPos = 1
	}

	drawBitmapRegion(bmpCreators, 0, srcPos*240, 320, 240, int32(xPos), 0)
}

/**
 * Return a "list" of scene functions
 *
 * Returns:
 * Scene object
 */
func crGetScene() scene {

	// Put scenes to a local scene object and return it
	s := scene{
		onInit:    crInit,
		onUpdate:  crUpdate,
		onDraw:    crDraw,
		onDestroy: nil,
	}

	return s

}
