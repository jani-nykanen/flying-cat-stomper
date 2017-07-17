/**
 * Graphics utility functions for faster drawing
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import "github.com/veandco/go-sdl2/sdl"

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

	rend.Copy(bmp.texture, &sdl.Rect{X: 0, Y: 0, W: bmpFont.width, H: bmpFont.height},
		&sdl.Rect{X: x, Y: y, W: bmpFont.width, H: bmpFont.height})
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
