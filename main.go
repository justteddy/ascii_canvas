package main

import (
	"log"

	"canvas/canvas"
)

func main() {
	c := canvas.New(30, 30)
	err := c.DrawRectangle(14, 1, 7, 6, ".", "")
	assertNoErr(err)

	err = c.DrawRectangle(1, 3, 8, 4, "", "0")
	assertNoErr(err)
	err = c.DrawRectangle(5, 5, 5, 3, "X", "X")
	assertNoErr(err)

	c.Draw()
}

func assertNoErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
