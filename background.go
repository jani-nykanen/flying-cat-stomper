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

/// Sky bitmap
var bmpSky bitmap

/// Mountains bitmap
var bmpMountains bitmap

/// Forest bitmap
var bmpForest bitmap

/// Mountain position
var mountainPos float32

/// Forest position
var forestPost float32

/**
 * Initialize background
 *
 * Params:
 * ass Assets object
 */
func bgInit(ass assets) {

	bmpSky = ass.getBitmap("sky")
	bmpMountains = ass.getBitmap("mountains")
	bmpForest = ass.getBitmap("forest")

	mountainPos = 0.0
	forestPost = 0.0

}

/**
 * Update background
 *
 * Params:
 * timeMul Time multiplier
 * globalSpeed Global speed multiplier
 */
func bgUpdate(timeMul float32, globalSpeed float32) {
	// Move mountains
	mountainPos -= globalSpeed * 0.25 * timeMul
	if mountainPos <= -float32(bmpMountains.width) {
		mountainPos += float32(bmpMountains.width)
	}

	// Move forest
	forestPost -= globalSpeed * 0.5 * timeMul
	if forestPost <= -float32(bmpForest.width) {
		forestPost += float32(bmpForest.width)
	}
}

/**
 * Draw background
 *
 * Params:
 * rend Renderer
 */
func bgDraw(rend *sdl.Renderer) {
	drawBitmap(bmpSky, 0, -16)

	for i := int32(0); i < 2; i++ {
		drawBitmap(bmpMountains, int32(mountainPos)+i*480, 40)
	}

	for i := int32(0); i < 2; i++ {
		drawBitmap(bmpForest, int32(forestPost)+i*320, 240-96)
	}
}
