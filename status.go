/**
 * Status
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

/// Status data, who would have guessed!
type statusdata struct {
	score     uint
	best      uint
	killCount uint
}

// Global status object
var status statusdata

/**
 * Init status
 */
func initStatus() {
	status.score = 0
	status.best = 0
	status.killCount = 0
}
