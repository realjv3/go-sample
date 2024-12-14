//go:build failtidy

package main

import "https://github.com/something/that-does-not-exist-at-all"

func getInt() int {

	var i int = 24

	return i
}
