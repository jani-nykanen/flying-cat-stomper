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
	cats      [32]cat
	bmpCats   bitmap
	bmpBunny  bitmap
	bmpStar   bitmap
	sndHurt   sound
	genTimer  float32
	catPos    float32
	pl        player
	stars     [8]star
	messages  [8]message
	spcCount1 int
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
	gob.bmpBunny = ass.getBitmap("bunny")
	gob.bmpStar = ass.getBitmap("star")

	// Set sounds
	gob.sndHurt = ass.getSound("hurt")

	// Set seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Create cats
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i] = newCat()
	}

	// Init player
	gob.pl.init(vec2{x: 64, y: 80}, ass)

	// Set generator values
	gob.genTimer = 0.0
	gob.catPos = 120.0
	gob.spcCount1 = int(rand.Float32()*4 + 4)

	// Init stars
	for i := 0; i < 8; i++ {
		gob.stars[i] = newStar()
	}

	// Init messages
	for i := 0; i < 8; i++ {
		gob.messages[i] = newMessage()
	}

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
		if gob.cats[i].exist == false && gob.cats[i].dead == false {
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

	// Update cats
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i].update(timeMul, globalSpeed)
		gob.cats[i].onPlayerCollision(&gob.pl)
	}

	// Update player
	gob.pl.update(timeMul)

	if getKeyState(sdl.SCANCODE_X) == StatePressed {
		gob.putCat(320+24, 48+rand.Float32()*(240.0-64), 0)
	}

	// Update stars
	for i := 0; i < 8; i++ {
		gob.stars[i].update(timeMul)
	}

	// Update messages
	for i := 0; i < 8; i++ {
		gob.messages[i].update(timeMul)
	}

	targetID := 0

	// Generate cats when the timer hits zero
	gob.genTimer -= 1.0 * globalSpeed * timeMul
	if gob.genTimer <= 0.0 {

		posDelta := rand.Float32()*(160) - 80
		if (posDelta > 0.0 && gob.catPos+posDelta > 240-32) || (posDelta < 0.0 && gob.catPos+posDelta < 80) {
			posDelta *= -1
		}
		gob.catPos += posDelta

		if gob.catPos > 96 && gob.catPos < 192 {
			gob.spcCount1--
			if gob.spcCount1 <= 0 {
				targetID = 1
				gob.spcCount1 = int(rand.Float32() + 4)
			}
		}

		gob.putCat(320+24, gob.catPos, targetID)

		gob.genTimer += 90.0
	}
}

/**
 * Draw game objects
 *
 * Params:
 * rend Renderer
 */
func (gob *gameobjects) draw(rend *sdl.Renderer) {

	// Draw cats
	for i := 0; i < len(gob.cats); i++ {
		gob.cats[i].draw(gob.bmpCats, rend)
	}

	// Draw stars
	for i := 0; i < 8; i++ {
		gob.stars[i].draw(gob.bmpStar, rend)
	}

	// Draw player
	gob.pl.draw(gob.bmpBunny, rend)

	// Draw messages
	for i := 0; i < 8; i++ {
		gob.messages[i].draw(hud.bmpFontMedium, rend)
	}
}
