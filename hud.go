/**
 * "HUD" (kind of)
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"math"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type HUD struct {
	/// Font small
	bmpFont bitmap
	/// Font big
	bmpFontBig bitmap
	/// Font medium
	bmpFontMedium bitmap
	/// HUD elemenents bitmap
	bmpHud bitmap
	/// Framerate
	fps int
	/// FPS sum
	fpsSum int
	/// FPS count
	fpsCount int
}

/// Global HUD object
var hud HUD

/// Font small
var bmpFont bitmap

/// Font big
var bmpFontBig bitmap

/// Font medium
var bmpFontMedium bitmap

/// HUD elemenents bitmap
var bmpHud bitmap

/// Framerate
var fps int

/// FPS sum
var fpsSum int

/// FPS count
var fpsCount int

/**
 * Initialize HUD elements
 *
 * Params:
 * ass Assets object
 */
func (hud *HUD) init(ass assets) {
	hud.bmpFont = ass.getBitmap("font")
	hud.bmpFontBig = ass.getBitmap("fontBig")
	hud.bmpHud = ass.getBitmap("hud")
	hud.bmpFontMedium = ass.getBitmap("fontMedium")

	hud.fpsSum = 0
	hud.fpsCount = 0
	hud.fps = 60
}

/**
 * Update HUD
 *
 * Params:
 * timeMul Time multiplier
 */
func (hud *HUD) update(timeMul float32) {

	// Since timeMul varies from frame to frame, let's take the
	// average in one second (or actually, a time frame that is at
	// least a second )
	hud.fpsSum += int(math.Floor(60.0 / float64(timeMul)))
	hud.fpsCount++
	if hud.fpsCount >= 60 {
		hud.fps = int(float32(hud.fpsSum) / float32(hud.fpsCount))
		hud.fpsCount = 0
		hud.fpsSum = 0
	}
}

/**
 * Draw HUD
 *
 * Params:
 * rend Renderer
 */
func (hud *HUD) draw(rend *sdl.Renderer) {
	// Draw score
	drawBitmapRegion(hud.bmpHud, 0, 0, 48, 12, 160-24, 2)
	drawCenteredText(hud.bmpFontBig, strconv.Itoa(int(status.score)), 160-6, 11)

	// Draw best
	drawBitmapRegion(hud.bmpHud, 0, 12, 48, 12, 320-24-32, 2)
	drawCenteredText(hud.bmpFontMedium, strconv.Itoa(int(status.best)), 320-32-4, 13)

	drawText(hud.bmpFont, "FPS: "+strconv.Itoa(hud.fps), 2, 240-16)
}
