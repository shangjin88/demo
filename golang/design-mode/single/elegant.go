package main

import "sync"

type color struct {
}

var c *color
var once sync.Once

func GetColorIns() *color {
	once.Do(func() {
		c = &color{}
	})
	return c
}
