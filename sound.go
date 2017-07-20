/**
 * Sound effects
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"github.com/veandco/go-sdl2/mix"
)

/// Sound object
type sound struct {
	chunk   *mix.Chunk
	channel int
	played  bool
	name    string
}

/**
 * Load sound
 *
 * Params:
 * path Path
 *
 * Returns:
 * A new sound
 * Error code, 0 on success
 */
func loadSound(path string, name string) (sound, int) {

	var errorObj error
	var s sound

	s.chunk, errorObj = mix.LoadWAV(path)
	if errorObj != nil {
		setErr("Failed to load a wav file in " + path)
		return s, 1
	}

	s.name = name

	return s, 0
}

/**
 * Play a sound effect
 *
 * Params:
 * vol Volume
 */
func (s *sound) play(vol float32) {

	if !s.played {
		mix.HaltChannel(s.channel)
		s.played = true
	}

	var errObj error
	s.channel, errObj = s.chunk.Play(-1, 0)
	_ = errObj

	s.chunk.Volume(int(vol * 128))

}
