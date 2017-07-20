/**
 * Flying message
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

/// Message
type message struct {
	pos   vec2
	exist bool
	val   int
	speed float32
}

/**
 * Create new, non-existent message
 *
 * Returns:
 * A new message
 */
func newMessage() message {
	var m message

	m.exist = false
	m.val = 0
	m.speed = 0.0

	return m
}

/**
 * Create "message"
 *
 * Params:
 * pos Position
 * val New value
 */
func (m *message) create(pos vec2, val int) {
	m.pos = pos
	m.exist = true
	m.val = val
	m.speed = -1.0
}

/**
 * Update message
 *
 * Params:
 * timeMul Time multiplier
 */
func (m *message) update(timeMul float32) {
	if m.exist {
		m.pos.y += m.speed * timeMul

		m.speed += 0.01 * timeMul
		if m.speed > 0.0 {
			m.exist = false
		}
	}
}

/**
 * Draw message
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (m *message) draw(bmp bitmap, rend *sdl.Renderer) {

	if m.exist {
		drawCenteredText(bmp, "+"+strconv.Itoa(m.val), int32(m.pos.x)-10, int32(m.pos.y)-18)
	}

}
