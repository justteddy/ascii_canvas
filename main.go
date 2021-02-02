package main

import (
	"log"

	"canvas/canvas"
)

func main() {
	c := canvas.New(10, 10)
	err := c.DrawRectangle(1, 1, 3, 3, "@", "")
	err = c.DrawRectangle(3, 3, 5, 5, "x", "")
	err = c.DrawRectangle(9, 9, 2, 2, "o", "")
	assertNoErr(err)

	c.Draw()
}

func assertNoErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
