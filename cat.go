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

/// Gas
type gas struct {
	pos   vec2
	timer float32
	exist bool
}

/**
 * Create new, non-existent gas
 *
 * Returns:
 * A new gas
 */
func newGas() gas {
	var g gas

	g.exist = false

	return g
}

/**
 * Create gas
 *
 * Params:
 * pos Position
 */
func (g *gas) create(pos vec2) {
	g.pos = pos
	g.exist = true
	g.timer = 60
}

/**
 * Update gas
 *
 * Params:
 * timeMul Time multiplier
 */
func (g *gas) update(timeMul float32) {
	if g.exist {
		g.timer -= 2.0 * timeMul
		if g.timer <= 0.0 {
			g.exist = false
		}
	}
}

/**
 * Draw gas
 *
 * Params:
 * bmp Bitmap
 * rend Renderer
 */
func (g *gas) draw(bmp bitmap, rend *sdl.Renderer) {

	if g.exist {
		scale := g.timer / 60.0
		dx := g.pos.x - (scale)*12
		dy := g.pos.y - (scale)*12
		drawScaledBitmapRegion(bmp, 0, 0, 32, 32, int32(dx), int32(dy), int32(24*scale), int32(24*scale))
	}

}

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
	speed      float32
	gases      [8]gas
	gasTimer   float32
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

	for i := 0; i < 8; i++ {
		c.gases[i] = newGas()
	}

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
	c.speed = 1.0
	c.gasTimer = 5.0

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

	// Update gas
	for i := 0; i < 8; i++ {
		c.gases[i].update(timeMul)
	}

	if c.exist == false {
		if c.dead {
			if c.typeid != 3 {
				c.pos.y += 0.75 * timeMul
			}
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
	c.pos.x -= c.speed * globalSpeed * timeMul
	// If out of screen, stop existing
	if c.pos.x < -48 {
		c.exist = false
	}

	// Update sin modifier
	if c.typeid == 0 {
		c.sinMod += 0.05 * timeMul
		// Set y position
		c.pos.y = c.starty + float32(math.Sin(float64(c.sinMod))*8)
	} else if c.typeid == 1 {
		c.sinMod += 0.03 * (globalSpeed / 2.0) * timeMul
		c.pos.y = c.starty + float32(math.Floor((math.Sin(float64(c.sinMod)) * float64(c.waveHeight/2.0))))
	} else if c.typeid == 2 {
		c.speed += 0.015 * timeMul
	} else if c.typeid == 3 {
		c.sinMod += 0.06 * timeMul
		c.pos.y = c.starty + float32(math.Sin(float64(c.sinMod))*16)
	}

	if c.typeid == 2 {
		c.gasTimer -= 1.0 * timeMul
		if c.gasTimer <= 0.0 {
			// Create gas
			for i := 0; i < len(c.gases); i++ {
				if c.gases[i].exist == false {
					c.gases[i].create(vec2{x: c.pos.x + 22, y: c.pos.y})
					break
				}
			}
			c.gasTimer += 5.0
		}
	}

	// Set anim speed
	animSpeed := float32(6.0)
	if c.typeid == 1 {
		animSpeed = 3.0
	} else if c.typeid == 2 {
		animSpeed = 4.0
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
	if c.exist == false || pl.isPassive {
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

		status.score += uint(c.typeid + 1)

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
				gobj.messages[i].create(vec2{x: pl.pos.x, y: c.pos.y - 4}, int(uint(c.typeid+1)))
				break
			}
		}

		if c.typeid == 3 {
			gobj.sndDestroy.play(0.5)

			for i := 0; i < len(gobj.cats); i++ {
				if gobj.cats[i].exist == true {
					gobj.cats[i].dead = true
					gobj.cats[i].exist = false
					gobj.cats[i].spr.currentFrame = 0
					gobj.cats[i].spr.changeFrameCount = 0
					gobj.cats[i].spr.currentRow = 2*c.typeid + 1
				}
			}
			bg.flashTimer = 60.0

		} else {
			gobj.sndHurt.play(0.35)
		}
	}
}

/**
 * Draw a cat
 *
 * Params:
 * bmp Bitmap
 * bmpGas Gas bitmap
 * rend Renderer
 */
func (c *cat) draw(bmp bitmap, bmpGas bitmap, rend *sdl.Renderer) {

	// Draw gas
	for i := 0; i < 8; i++ {
		c.gases[i].draw(bmpGas, rend)
	}

	if c.exist == false && !c.dead {
		return
	}

	// Draw sprite
	c.spr.draw(bmp, rend, int32(math.Floor(float64(c.pos.x-24))), int32(math.Floor(float64(c.pos.y-24))))

}
