/* vim: set ts=4 sw=4: */

package main

import (
	"flag"
	"sync"
)

var once sync.Once

func main() {
	var newApp bool

	once.Do(func() {
		flag.BoolVar(&newApp, "new", false, "Creates the basic structure of a new app in an empty directory")
		flag.Parse()
	})

	if !newApp {
		flag.Usage()
		return
	}
}
