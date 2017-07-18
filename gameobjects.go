/**
 * Game objects
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

/// Game objects
type gameobjects struct {
	cats    [32]cat
	bmpCats bitmap
}

/// Game objects global object
var gobj gameobjects

/**
 * Init game objects
 *
 * Param:
 * ass Assets object
 */
func (gob *gameobjects) init(ass assets) {
	// Set bitmaps
	gob.bmpCats = ass.getBitmap("cats")

	// Create cats
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i] = newCat()
	}

	// Set seed
	rand.Seed(time.Now().UTC().UnixNano())
}

/**
 * Put a living cat to the game screen
 *
 * Params:
 * x X coord
 * y Y coord
 * id Type ID
 */
func (gob *gameobjects) putCat(x, y float32, id int) {

	// Go through the array of cats
	for i := 0; i < len(gob.cats); i++ {
		// If not exist, create
		if gob.cats[i].exist == false {
			gob.cats[i].createCat(x, y, id)
			break
		}
	}
}

/**
 * Update game objects
 *
 * Params:
 * timeMul Time multiplier
 * globalSpeed Global speed multiplier
 */
func (gob *gameobjects) update(timeMul, globalSpeed float32) {
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i].update(timeMul, globalSpeed)
	}

	if getKeyState(sdl.SCANCODE_X) == STATE_PRESSED {
		gob.putCat(320+24, 48+rand.Float32()*(240.0-64), 0)
	}
}

/**
 * Draw game objects
 *
 * Params:
 * rend Renderer
 */
func (gob *gameobjects) draw(rend *sdl.Renderer) {
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i].draw(gob.bmpCats, rend)
	}
}
