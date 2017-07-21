/**
 * Status
 *
 * Author:
 * (c) 2017 Jani Nykänen
 */

package main

import (
	"io/ioutil"
	"strconv"
)

/// Status data, who would have guessed!
type statusdata struct {
	score uint
	best  uint
}

// Global status object
var status statusdata

/**
 * Init status
 */
func initStatus() {
	status.score = 0
	status.best = 0
}

/**
 * Read the best score
 *
 * Params:
 * path The file path
 */
func readBest(path string) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		// Ignore error, but return
		return
	}
	str := string(b)

	var e error
	var num int
	num, e = (strconv.Atoi(str))
	if e != nil {
		err = e
		return
	}
	status.best = uint(num)

}

/**
 * Store best
 *
 * Params:
 * path The file path
 */
func storeBest(path string) {
	d := []byte(strconv.Itoa(int(status.best)))
	_ = ioutil.WriteFile(path, d, 0644)
}
