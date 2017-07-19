/**
 * Background
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 */

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	/// Sky bitmap
	bmpSky bitmap

	/// Mountains bitmap
	bmpMountains bitmap

	/// Forest bitmap
	bmpForest bitmap

	/// Mountain position
	mountainPos float32

	/// Forest position
	forestPost float32
}

/// Background object
var bg background

/**
 * Initialize background
 *
 * Params:
 * ass Assets object
 */
func (bg *background) init(ass assets) {

	bg.bmpSky = ass.getBitmap("sky")
	bg.bmpMountains = ass.getBitmap("mountains")
	bg.bmpForest = ass.getBitmap("forest")

	bg.mountainPos = 0.0
	bg.forestPost = 0.0

}

/**
 * Update background
 *
 * Params:
 * timeMul Time multiplier
 * globalSpeed Global speed multiplier
 */
func (bg *background) update(timeMul float32, globalSpeed float32) {
	// Move mountains
	bg.mountainPos -= globalSpeed * 0.2 * timeMul
	if bg.mountainPos <= -float32(bg.bmpMountains.width) {
		bg.mountainPos += float32(bg.bmpMountains.width)
	}

	// Move forest
	bg.forestPost -= globalSpeed * 0.5 * timeMul
	if bg.forestPost <= -float32(bg.bmpForest.width) {
		bg.forestPost += float32(bg.bmpForest.width)
	}
}

/**
 * Draw background
 *
 * Params:
 * rend Renderer
 */
func (bg *background) draw(rend *sdl.Renderer) {
	drawBitmap(bg.bmpSky, 0, -16)

	for i := int32(0); i < 2; i++ {
		drawBitmap(bg.bmpMountains, int32(bg.mountainPos)+i*480, 40)
	}

	for i := int32(0); i < 2; i++ {
		drawBitmap(bg.bmpForest, int32(bg.forestPost)+i*320, 240-96)
	}
}
