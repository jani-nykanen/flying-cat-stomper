/**
 * Sprite
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import "github.com/veandco/go-sdl2/sdl"

type sprite struct {
	currentFrame     int
	currentRow       int
	changeFrameCount float32
	width            int
	height           int
}

/**
 * Create a new sprite
 *
 * Params:
 * w Width
 * h Height
 *
 * Returns:
 * A sprite
 */
func createSprite(w, h int) sprite {
	var s sprite
	s.width = w
	s.height = h
	s.currentFrame = 0
	s.currentRow = 0
	s.changeFrameCount = 0.0
	return s
}

/**
 * Animate a sprite
 *
 * Params:
 * row Row
 * start Starting frame
 * end Ending frame
 * speed Speed
 * timeMul Time multiplier
 */
func (spr *sprite) animate(row, start, end int, speed, timeMul float32) {

	if start == end {

		spr.changeFrameCount = 0
		spr.currentFrame = start
		spr.currentRow = row
		return
	}

	if spr.currentRow != row {

		spr.changeFrameCount = 0
		spr.currentFrame = start
		spr.currentRow = row
	}

	if spr.currentFrame < start {

		spr.currentFrame = start
	}

	spr.changeFrameCount += 1.0 * timeMul
	if spr.changeFrameCount > speed {

		spr.currentFrame++

		if spr.currentFrame > end {

			spr.currentFrame = start
		}
		spr.changeFrameCount -= speed
	}
}

/**
 * Draw a sprite
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 * dx Destination x
 * dy Destination y
 */
func (spr *sprite) draw(bmp bitmap, rend *sdl.Renderer, dx, dy int32) {
	x := int32(spr.currentFrame * spr.width)
	y := int32(spr.currentRow * spr.height)

	drawBitmapRegion(bmp, x, y, int32(spr.width), int32(spr.height), dx, dy)
}

/**
 * Draw a scaled sprite
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 * dx Destination x
 * dy Destination y
 * scale Scale
 */
func (spr *sprite) drawScaled(bmp bitmap, rend *sdl.Renderer, dx, dy int32, scale float32) {
	x := int32(spr.currentFrame * spr.width)
	y := int32(spr.currentRow * spr.height)

	dw := int32(scale * float32(spr.width))
	dh := int32(scale * float32(spr.height))

	drawScaledBitmapRegion(bmp, x, y, int32(spr.width), int32(spr.height), dx, dy, dw, dh)
}
