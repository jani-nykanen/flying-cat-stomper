/**
 * Player object
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// Player struct
type player struct {
	pos           vec2
	spr           sprite
	speed         vec2
	target        vec2
	isPassive     bool
	catTouchTimer float32
	doubleJump    bool
}

/**
 * Initialize a player object
 *
 * Params:
 * pos New position
 */
func (pl *player) init(pos vec2) {
	pl.pos = pos
	pl.spr = createSprite(48, 48)
	pl.spr.currentFrame = 1
	pl.speed = vec2{x: 0, y: 0}
	pl.target = vec2{x: 0, y: 0}
	pl.isPassive = true
	pl.doubleJump = false
	pl.catTouchTimer = 0.0
}

/**
 * Controls
 */
func (pl *player) control() {

	if pl.isPassive {
		if getMouseButtonState(1) == StatePressed {
			pl.speed.y = -4.0
			pl.isPassive = false
			pl.doubleJump = false
		} else {
			return
		}
	}

	// Horizontal movement
	deltaX := (pl.pos.x - float32(getCursorPos().x)) / 64.0
	if deltaX < 0.0 {
		deltaX -= 1.0
	} else if deltaX > 0.0 {
		deltaX += 1.0
	}
	pl.target.x = -deltaX

	// Vertical movement
	pl.target.y = 2.0

	if getMouseButtonState(1) == StatePressed {

		if !pl.doubleJump {
			pl.doubleJump = true
			pl.speed.y = -4.0
		}
	}

	if getMouseButtonState(1) == StateReleased {

		if pl.speed.y < 0.0 {
			pl.speed.y /= 1.5
		}
	}
}

/**
 * Move
 *
 * Params:
 * timeMul Time multiplier
 */
func (pl *player) move(timeMul float32) {

	// Horizontal speed
	if pl.target.x > pl.speed.x {
		pl.speed.x += 0.125 * timeMul
		if pl.speed.x > pl.target.x {
			pl.speed.x = pl.target.x
		}
	} else if pl.target.x < pl.speed.x {
		pl.speed.x -= 0.125 * timeMul
		if pl.speed.x < pl.target.x {
			pl.speed.x = pl.target.x
		}
	}

	// Vertical speed
	if pl.target.y > pl.speed.y {
		pl.speed.y += 0.125 * timeMul
		if pl.speed.y > pl.target.y {
			pl.speed.y = pl.target.y
		}
	} else if pl.target.y < pl.speed.y {
		pl.speed.y -= 0.125 * timeMul
		if pl.speed.y < pl.target.y {
			pl.speed.y = pl.target.y
		}
	}

	// Update position
	pl.pos.x += pl.speed.x * timeMul
	pl.pos.y += pl.speed.y * timeMul

	// Make sure the player is not outside the game area
	if pl.pos.x < 16 {
		pl.pos.x = 16
		pl.speed.x = 0.0
	} else if pl.pos.x > 320-16 {
		pl.pos.x = 320 - 16
		pl.speed.x = 0.0
	}

	// TODO: Death here
	if pl.pos.y-44 > 240.0 {
		pl.pos.x = 64
		pl.pos.y = 80
		pl.isPassive = true
	}
}

/**
 * Update player
 *
 * Params:
 * timeMul Time multiplier
 */
func (pl *player) update(timeMul float32) {

	pl.control()
	if !pl.isPassive {
		pl.move(timeMul)
	}

	// Animate
	if pl.doubleJump && pl.speed.y < 0.0 {
		pl.spr.animate(1, 0, 3, 3, timeMul)

	} else {
		pl.spr.currentRow = 0
		if math.Abs(float64(pl.speed.y)) < 0.75 {
			pl.spr.currentFrame = 1
		} else {
			if pl.speed.y < 0.0 {
				pl.spr.currentFrame = 2
			} else {
				pl.spr.currentFrame = 0
			}
		}
	}

	// Cat touch timer
	if pl.catTouchTimer > 0.0 {
		if pl.doubleJump {
			pl.doubleJump = false
			pl.speed.y *= 2.75 / 2.0
			pl.spr.currentRow = 0
		}
		pl.catTouchTimer -= 1.0 * timeMul
	}

}

/**
 * Draw player
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (pl *player) draw(bmp bitmap, rend *sdl.Renderer) {
	pl.spr.draw(bmp, rend, int32(math.Floor(float64(pl.pos.x)-24)),
		int32(math.Floor(float64(pl.pos.y)-44)))
}
