/**
 * Cat
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

/// Cat struct
type cat struct {
	pos        vec2
	starty     float32
	typeid     int
	spr        sprite
	exist      bool
	sinMod     float32
	dead       bool
	waveHeight float32
}

/**
 * Create a new cat object
 *
 * Returns:
 * A cat object
 */
func newCat() cat {
	var c cat
	c.exist = false
	c.spr = createSprite(48, 48)
	c.waveHeight = 0.0
	return c
}

/**
 * Put a cat object to the game world
 *
 * Params:
 * x X coordinate
 * y Y coordinate
 * id Type id
 */
func (c *cat) createCat(x, y float32, id int) {

	c.exist = true
	c.pos.x = x
	c.pos.y = y
	c.starty = y
	c.typeid = id
	c.sinMod = rand.Float32() * 2 * math.Pi
	c.dead = false

	if id == 1 {
		c.waveHeight = float32(math.Min(math.Abs(float64(y-32)), math.Abs(float64(y-(320-24)))))
	}
}

/**
 * Update a cat
 *
 * Params:
 * timeMul Time multiplier
 * globalSpeed Global speed multiplier
 */
func (c *cat) update(timeMul, globalSpeed float32) {

	if c.exist == false {
		if c.dead {
			c.pos.y += 0.75 * timeMul
			c.pos.x -= globalSpeed * timeMul
			c.spr.animate(c.typeid*2+1, 0, 5, 5, timeMul)
			if c.spr.currentFrame == 5 {
				c.dead = false
				return
			}
		}
		return
	}

	// Move
	c.pos.x -= globalSpeed * timeMul
	// If out of screen, stop existing
	if c.pos.x < -48 {
		c.exist = false
	}

	// Update sin modifier
	if c.typeid == 0 {
		c.sinMod += 0.05 * timeMul
		// Set y position
		c.pos.y = c.starty + float32(math.Sin(float64(c.sinMod))*8)
	} else {
		c.sinMod += 0.03 * (globalSpeed / 2.0) * timeMul
		c.pos.y = c.starty + float32(math.Floor((math.Sin(float64(c.sinMod)) * float64(c.waveHeight/2.0))))
	}

	animSpeed := float32(6.0)
	if c.typeid == 1 {
		animSpeed = 3.0
	}

	// Animate
	c.spr.animate(c.typeid*2, 0, 3, animSpeed, timeMul)

}

/**
 * Cat-player collision
 *
 * Param:
 * pl Player object
 */
func (c *cat) onPlayerCollision(pl *player) {
	if c.exist == false {
		return
	}

	if pl.speed.y > 0.0 && pl.pos.x+8 > c.pos.x-24 && pl.pos.x-8 < c.pos.x+24 && pl.pos.y > c.pos.y-12 && pl.pos.y < c.pos.y-4 {
		c.exist = false

		pl.speed.y = -4.0

		pl.catTouchTimer = 15
		pl.doubleJump = false

		c.dead = true
		c.spr.currentFrame = 0
		c.spr.changeFrameCount = 0
		c.spr.currentRow = 2*c.typeid + 1

		status.killCount++
		status.score += status.killCount * uint(c.typeid+1)

		// Create star (or "pow!")
		for i := 0; i < len(gobj.stars); i++ {
			if gobj.stars[i].exist == false {
				gobj.stars[i].create(vec2{x: pl.pos.x, y: pl.pos.y})
				break
			}
		}

		// Create a message
		for i := 0; i < len(gobj.messages); i++ {
			if gobj.messages[i].exist == false {
				gobj.messages[i].create(vec2{x: pl.pos.x, y: c.pos.y - 4}, int(status.killCount*uint(c.typeid+1)))
				break
			}
		}

		gobj.sndHurt.play(0.5)
	}
}

/**
 * Draw a cat
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (c *cat) draw(bmp bitmap, rend *sdl.Renderer) {

	if c.exist == false && !c.dead {
		return
	}

	// Draw sprite
	c.spr.draw(bmp, rend, int32(math.Floor(float64(c.pos.x-24))), int32(math.Floor(float64(c.pos.y-24))))

}
