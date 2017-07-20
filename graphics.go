/**
 * Graphics utility functions for faster drawing
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

/**
 * Draw a bitmap
 *
 * Params:
 * bmp Bitmap
 * x X
 * y Y
 */
func drawBitmap(bmp bitmap, x int32, y int32) {
	rend := getRend()

	rend.Copy(bmp.texture, &sdl.Rect{X: 0, Y: 0, W: bmp.width, H: bmp.height},
		&sdl.Rect{X: x, Y: y, W: bmp.width, H: bmp.height})
}

/**
 * Draw a bitmap region
 *
 * Params:
 * bmp Bitmap
 * sx Source X
 * sy Source Y
 * sw Source width
 * sh Source height
 * dx Destination X
 * dy Destination Y
 */
func drawBitmapRegion(bmp bitmap, sx int32, sy int32, sw int32, sh int32, dx int32, dy int32) {
	rend := getRend()

	rend.Copy(bmp.texture, &sdl.Rect{X: sx, Y: sy, W: sw, H: sh},
		&sdl.Rect{X: dx, Y: dy, W: sw, H: sh})
}

/**
 * Draw a scaledbitmap region
 *
 * Params:
 * bmp Bitmap
 * sx Source X
 * sy Source Y
 * sw Source width
 * sh Source height
 * dx Destination X
 * dy Destination Y
 * dw Destination width
 * dh Destination height
 */
func drawScaledBitmapRegion(bmp bitmap, sx int32, sy int32, sw int32, sh int32, dx int32, dy int32, dw int32, dh int32) {
	rend := getRend()

	rend.Copy(bmp.texture, &sdl.Rect{X: sx, Y: sy, W: sw, H: sh},
		&sdl.Rect{X: dx, Y: dy, W: dw, H: dh})
}

/**
 * Draw text with a bitmap font
 *
 * Params:
 * font Bitmap font
 * text Text
 * x X
 * y Y
 */
func drawText(font bitmap, text string, x int32, y int32) {

	dx := x
	dy := y

	cw := int32(font.width / 16)  // Character width
	ch := int32(font.height / 16) // Character height

	var c byte
	var row int32
	var column int32

	for i := 0; i < len(text); i++ {
		c = text[i]

		if c == ("\n")[0] {
			dy += int32(ch)
			continue
		}

		row = int32(math.Floor(float64(c) / 16.0))
		column = int32(c) - row*16

		drawBitmapRegion(font, column*cw, row*ch, cw, ch, dx, dy)

		dx += int32(float32(cw) / 5.0 * 3.0)
	}

}

/**
 * Draw text with a bitmap font, centered
 *
 * Params:
 * {See drawText}
 */
func drawCenteredText(font bitmap, text string, x int32, y int32) {

	cw := int32(font.width / 16) // Character width
	x -= int32(float32(cw) * float32(len(text)) / 2.0 * (3.0 / 5.0))

	drawText(font, text, x, y)

}
