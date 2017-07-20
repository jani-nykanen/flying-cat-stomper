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

/// "Whitening"
type whitening struct {
	pos   vec2
	timer float32
	exist bool
	spr   sprite
}

/**
 * Create new, non-existent whitening
 *
 * Returns:
 * A new whitening
 */
func newWhitening() whitening {
	var w whitening

	w.exist = false
	w.spr = createSprite(48, 48)

	return w
}

/**
 * Create "whitening"
 *
 * Params:
 * pos Position
 */
func (wh *whitening) create(pos vec2, column, row int) {
	wh.pos = pos
	wh.exist = true
	wh.timer = 60
	wh.spr.currentFrame = column
	wh.spr.currentRow = row
}

/**
 * Update whitening
 *
 * Params:
 * timeMul Time multiplier
 */
func (wh *whitening) update(timeMul float32) {
	if wh.exist {
		wh.timer -= 1.75 * timeMul
		if wh.timer <= 0.0 {
			wh.exist = false
		}
	}
}

/**
 * Draw whitening
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (wh *whitening) draw(bmp bitmap, rend *sdl.Renderer) {

	if wh.exist {
		scale := wh.timer / 60.0
		dx := wh.pos.x - (scale)*24
		dy := wh.pos.y - (scale)*24
		wh.spr.drawScaled(bmp, rend, int32(dx), int32(dy), scale)
	}

}

/// Player struct
type player struct {
	pos           vec2
	spr           sprite
	speed         vec2
	target        vec2
	isPassive     bool
	catTouchTimer float32
	doubleJump    bool
	white         [16]whitening
	whiteTimer    float32
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

	for i := 0; i < 16; i++ {
		pl.white[i] = newWhitening()
	}
	pl.whiteTimer = 10.0
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
	pl.target.y = 2.5

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
		pl.speed.y += 0.15 * timeMul
		if pl.speed.y > pl.target.y {
			pl.speed.y = pl.target.y
		}
	} else if pl.target.y < pl.speed.y {
		pl.speed.y -= 0.15 * timeMul
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
 * Create white
 *
 * Params:
 * timeMul Time multiplier
 */
func (pl *player) createWhite(timeMul float32) {

	// Create white
	pl.whiteTimer -= 1.0 * timeMul
	if pl.whiteTimer <= 0.0 {

		for i := 0; i < len(pl.white); i++ {
			if pl.white[i].exist == false {
				pl.white[i].create(vec2{x: pl.pos.x, y: pl.pos.y - 20}, pl.spr.currentFrame, pl.spr.currentRow+2)
				break
			}
		}

		pl.whiteTimer += 5.0
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

	// Create white
	pl.createWhite(timeMul)

	// Update white
	for i := 0; i < len(pl.white); i++ {
		pl.white[i].update(timeMul)
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

	// Draw white
	for i := 0; i < len(pl.white); i++ {
		pl.white[i].draw(bmp, rend)
	}

	// Draw player
	pl.spr.draw(bmp, rend, int32(math.Floor(float64(pl.pos.x)-24)),
		int32(math.Floor(float64(pl.pos.y)-44)))
}
