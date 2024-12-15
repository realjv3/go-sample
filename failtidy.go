//go:build failtidy

package main

import "github.com/google/go-cmp/cmp"

func compare() {

	cmp.Diff("this", "that")
}
