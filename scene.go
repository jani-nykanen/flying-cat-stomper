/**
 * Scene structure
 *
 * Author:
 * (c) 2017 Jani Nykänen
 *
 * Todo:
 * I heard that Go supports "interfaces", but since I'm a C programmer,
 * I'll do things in a C way!
 */

package main

import "github.com/veandco/go-sdl2/sdl"

/// Function type, no arguments
type Function func() int

/// Function type, float32 argument
type FunctionFloat32 func(float32)

/// Function type, renderer argument
type FunctionRend func(*sdl.Renderer)

/// Scene structure
type scene struct {
	onInit    Function
	onUpdate  FunctionFloat32
	onDraw    FunctionRend
	onDestroy Function
}
