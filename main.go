package main

import (
	"log"
	"os"

	"canvas/canvas"
)

func main() {
	c := canvas.New()
	err := c.DrawRectangle(5, 5, 7, 3, ".", "")
	assertNoErr(err)

	//err = c.DrawRectangle(0, 3, 8, 4, "", "0")
	//assertNoErr(err)
	//err = c.DrawRectangle(5, 5, 5, 3, "X", "X")
	//assertNoErr(err)
	//
	//err = c.FloodFill(0, 0, "*")
	//assertNoErr(err)

	err = c.Render(os.Stdout)
	assertNoErr(err)
}

func assertNoErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
