// +build !appengine,!appenginevm

package main

import (
	"github.com/ynozue/hellow-gae-go/gae_echo"
)

func main() {
	gae_echo.Start()
}

