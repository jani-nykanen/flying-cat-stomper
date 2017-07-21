/**
* Assets
 *
* Author:
* (c) 2017 Jani Nykänen
*
* Todo:
* If I have time, I'll make the assets be loaded from
* an external text file
*/

package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

/// String pair
type stringPair struct {
	path string
	name string
}

/// Bitmap struct
type bitmap struct {
	texture *sdl.Texture
	width   int32
	height  int32
	name    string
}

/// A struct that contains all the assets
type assets struct {
	bitmaps []bitmap
	sounds  []sound
	music   *mix.Music
}

/**
 * Load a bitmap
 *
 * Params:
 * rend Renderer
 * path Bitmap path
 *
 * Returns:
 * A new bitmap
 * Error code, 0 on success
 */
func loadBitmap(rend *sdl.Renderer, path string, name string) (bitmap, int) {

	var bmp bitmap
	var err error

	// Load surface
	surf, err := img.Load(path)
	if err != nil {
		setErr("Failed to load a bitmap in " + path)
		return bmp, 1
	}

	// Create texture
	bmp.texture, err = rend.CreateTextureFromSurface(surf)
	if err != nil {
		setErr("Failed to create a bitmap from a surface loaded in " + path)
		return bmp, 1
	}

	// Set data
	bmp.width = surf.W
	bmp.height = surf.H
	bmp.name = name

	// Destroy surface
	surf.Free()

	return bmp, 0
}

/**
 * Load bitmaps from a list
 *
 * Params:
 * rend Renderer
 *
 * Returns:
 * 0 on success, 1 on error
 */
func loadBitmaps(rend *sdl.Renderer) (int, []bitmap) {

	// Let's temporary put list of bitmaps here
	bitmaps := [...]stringPair{
		stringPair{path: "assets/bitmaps/font.png", name: "font"},
		stringPair{path: "assets/bitmaps/font_big.png", name: "fontBig"},
		stringPair{path: "assets/bitmaps/sky.png", name: "sky"},
		stringPair{path: "assets/bitmaps/mountains.png", name: "mountains"},
		stringPair{path: "assets/bitmaps/forest.png", name: "forest"},
		stringPair{path: "assets/bitmaps/cursor.png", name: "cursor"},
		stringPair{path: "assets/bitmaps/hud.png", name: "hud"},
		stringPair{path: "assets/bitmaps/fontMedium.png", name: "fontMedium"},
		stringPair{path: "assets/bitmaps/cats.png", name: "cats"},
		stringPair{path: "assets/bitmaps/bunny.png", name: "bunny"},
		stringPair{path: "assets/bitmaps/star.png", name: "star"},
		stringPair{path: "assets/bitmaps/gas.png", name: "gas"},
	}

	var errCode int

	// Create a slice for bitmaps and load bitmaps there
	bmps := make([]bitmap, len(bitmaps))
	for i := 0; i < len(bitmaps); i++ {
		bmps[i], errCode = loadBitmap(rend, bitmaps[i].path, bitmaps[i].name)
		if errCode == 1 {
			return 0, bmps
		}
	}

	return 0, bmps

}

/**
 * Load sounds
 *
 * Returns:
 * Error code, 0 on success
 * An array of sounds
 */
func loadSounds() (int, []sound) {

	// Let's temporary put list of bitmaps here
	soundsArr := [...]stringPair{
		stringPair{path: "assets/sounds/hurt.wav", name: "hurt"},
		stringPair{path: "assets/sounds/jump.wav", name: "jump"},
		stringPair{path: "assets/sounds/destroy.wav", name: "destroy"},
	}

	var errCode int

	// Create a slice for bitmaps and load bitmaps there
	sounds := make([]sound, len(soundsArr))
	for i := 0; i < len(soundsArr); i++ {
		sounds[i], errCode = loadSound(soundsArr[i].path, soundsArr[i].name)
		if errCode == 1 {
			return 0, sounds
		}
	}

	return 0, sounds
}

/**
 * Load assets
 *
 * Return:
 * An assets object
 * Error code, 0 on success
 */
func loadAssets(rend *sdl.Renderer) (assets, int) {
	var ass assets

	var errCode int
	var err error

	// Load bitmaps
	errCode, ass.bitmaps = loadBitmaps(rend)
	if errCode == 1 {
		return ass, 1
	}

	// Load sounds
	errCode, ass.sounds = loadSounds()
	if errCode == 1 {
		return ass, 1
	}

	// Load music
	path := "assets/music/theme.ogg"
	ass.music, err = mix.LoadMUS("assets/music/theme.ogg")
	if err != nil {
		setErr("Failed to load a bitmap in " + path)
		return ass, 1
	}

	return ass, 0
}

/**
 * Get a bitmap
 *
 * Params:
 * name Bitmap name
 *
 * Returns:
 * A bitmap
 */
func (ass *assets) getBitmap(name string) bitmap {
	for i := 0; i < len(ass.bitmaps); i++ {
		if ass.bitmaps[i].name == name {
			return ass.bitmaps[i]
		}
	}

	return bitmap{texture: nil, width: 0, height: 0, name: ""}
}

/**
 * Get a sound
 *
 * Params:
 * name Sound name
 *
 * Returns:
 * A sound
 */
func (ass *assets) getSound(name string) sound {
	for i := 0; i < len(ass.sounds); i++ {
		if ass.sounds[i].name == name {
			return ass.sounds[i]
		}
	}

	return sound{chunk: nil, played: false, channel: 0, name: ""}
}

/**
 * Destroy assets
 * Params:
 * ass Assets object
 */
func destroyAssets(ass assets) {
	for i := 0; i < len(ass.bitmaps); i++ {
		ass.bitmaps[i].texture.Destroy()
	}
}
