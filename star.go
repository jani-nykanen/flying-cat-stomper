/**
 * Star
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import "github.com/veandco/go-sdl2/sdl"

/// Star
type star struct {
	pos   vec2
	exist bool
	spr   sprite
}

/**
 * Create a new star object
 *
 * Returns:
 * A new star
 */
func newStar() star {
	var st star

	st.exist = false
	st.spr = createSprite(64, 32)

	return st
}

/**
 * Create star
 *
 * Params:
 * pos Position
 */
func (st *star) create(pos vec2) {
	st.pos = pos
	st.exist = true
	st.spr.currentFrame = 0
	st.spr.currentRow = 0
}

/**
 * Update star
 *
 * Params:
 * timeMul Time multiplier
 */
func (st *star) update(timeMul float32) {
	if st.exist {
		st.spr.animate(0, 0, 6, 4, timeMul)
		if st.spr.currentFrame == 6 {
			st.exist = false
		}
	}
}

/**
 * Draw star
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (st *star) draw(bmp bitmap, rend *sdl.Renderer) {

	if st.exist {
		st.spr.draw(bmp, rend, int32(st.pos.x-32), int32(st.pos.y-16))
	}

}
