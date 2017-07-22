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

/// HUD object
type hudObject struct {
	bmpFont       bitmap
	bmpFontBig    bitmap
	bmpFontMedium bitmap
	bmpHud        bitmap
	bmpLogo       bitmap
	bmpGuide      bitmap
	sndStart      sound
	fps           int
	fpsSum        int
	fpsCount      int
	showTitle     bool
	overlayText   bool
	titleSinMod   float32
	titleGoaway   float32
	blackTimer    float32
}

/// Global HUD object
var hud hudObject

/**
 * Initialize HUD elements
 *
 * Params:
 * ass Assets object
 */
func (hud *hudObject) init(ass assets) {
	hud.bmpFont = ass.getBitmap("font")
	hud.bmpFontBig = ass.getBitmap("fontBig")
	hud.bmpHud = ass.getBitmap("hud")
	hud.bmpFontMedium = ass.getBitmap("fontMedium")
	hud.bmpLogo = ass.getBitmap("logo")
	hud.bmpGuide = ass.getBitmap("guide")

	hud.sndStart = ass.getSound("start")

	hud.fpsSum = 0
	hud.fpsCount = 0
	hud.fps = 60

	hud.showTitle = true
	hud.titleSinMod = 0.0
	hud.overlayText = false
	hud.titleGoaway = 0.0
	hud.blackTimer = 60.0
}

/**
 * Update title screen, when drawn
 *
 * Params:
 * timeMul Time multiplier
 */
func (hud *hudObject) updateTitle(timeMul float32) {

	if hud.titleGoaway > 0.0 {
		hud.titleGoaway -= 1.0 * timeMul
		return
	}

	hud.titleSinMod += 0.02 * timeMul

	mposi := getCursorPos()
	mpos := vec2{x: float32(mposi.x), y: float32(mposi.y)}

	hud.overlayText = mpos.x > 32 && mpos.x < 320-32 && mpos.y > 240-40 && mpos.y < 240-40+26

	if hud.overlayText && getMouseButtonState(1) == StatePressed {
		hud.showTitle = false
		hud.sndStart.play(1.0)
		hud.titleGoaway = 60.0
	}

}

/**
 * Update HUD
 *
 * Params:
 * timeMul Time multiplier
 */
func (hud *hudObject) update(timeMul float32) {

	if hud.blackTimer > 0.0 {
		hud.blackTimer -= 1.0 * timeMul
		return
	}

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

	if hud.showTitle || hud.titleGoaway > 0.0 {
		hud.updateTitle(timeMul)
	}
}

/**
 * Draw title
 *
 * Params:
 * rend Renderer
 *
 */
func (hud *hudObject) drawTitle(rend *sdl.Renderer) {

	pos := float32(0)
	if hud.titleGoaway > 0.0 {
		pos = -160 + 160*(hud.titleGoaway/60.0)
	}

	// Draw logo
	drawBitmap(hud.bmpLogo, 0, int32(pos)+int32(math.Sin(float64(hud.titleSinMod))*8))

	if hud.overlayText {
		hud.bmpFontMedium.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
		hud.bmpFontMedium.texture.SetColorMod(255, 255, 0)
	}

	// Draw start timer
	if hud.titleGoaway <= 0.0 {
		drawCenteredText(hud.bmpFontMedium, "Click here to start", 160, 240-38)
	}

	pos = 0
	if hud.titleGoaway > 0.0 {
		pos = 16 * (1.0 - hud.titleGoaway/60.0)
	}
	// Draw (c) and full screen thing
	drawCenteredText(hud.bmpFont, "(c) 2017 Jani Nyk~nen", 160, 240-14+int32(pos))
	drawCenteredText(hud.bmpFont, "Press F4 to toggle full screen", 160-int32(hud.titleSinMod*20)%320, -int32(pos))
	drawCenteredText(hud.bmpFont, "Press F4 to toggle full screen", 480-int32(hud.titleSinMod*20)%320, -int32(pos))

	if hud.overlayText {
		hud.bmpFontMedium.texture.SetColorMod(255, 255, 255)
	}

	// Draw guide
	pos = 0
	if hud.titleGoaway > 0.0 {
		pos = 320 * (1.0 - hud.titleGoaway/60.0)
	}
	drawCenteredText(hud.bmpFont, "Controls:", 160+int32(pos), 104)
	drawBitmap(hud.bmpGuide, int32(pos), 80)

	// Draw blackness
	if hud.blackTimer > 0.0 {
		pos = 120.0 * hud.blackTimer / 60.0
		rend.SetDrawColor(0, 0, 0, 255)
		recSrc = sdl.Rect{X: 0, Y: 0, W: 320, H: int32(pos)}
		rend.FillRect(&recSrc)
		recSrc = sdl.Rect{X: 0, Y: 120 + (120 - int32(pos)), W: 320, H: 120 - (120 - int32(pos))}
		rend.FillRect(&recSrc)
	}
}

/**
 * Draw HUD
 *
 * Params:
 * rend Renderer
 */
func (hud *hudObject) draw(rend *sdl.Renderer) {

	if hud.showTitle || hud.titleGoaway > 0.0 {
		hud.drawTitle(rend)
		if hud.titleGoaway <= 0.0 {
			return
		}
	}

	pos := float32(0)

	if hud.titleGoaway > 0.0 {
		pos = -64 + 64*(1.0-hud.titleGoaway/60.0)
	}

	// Draw score
	drawBitmapRegion(hud.bmpHud, 0, 0, 48, 12, 160-24, 2+int32(pos))
	drawCenteredText(hud.bmpFontBig, strconv.Itoa(int(status.score)), 160-6, 11+int32(pos))

	// Draw best
	drawBitmapRegion(hud.bmpHud, 0, 12, 48, 12, 320-24-32, 2+int32(pos))
	drawCenteredText(hud.bmpFontMedium, strconv.Itoa(int(status.best)), 320-32-4, 13+int32(pos))

	// drawText(hud.bmpFont, "FPS: "+strconv.Itoa(hud.fps), 2, 240-16)
}
