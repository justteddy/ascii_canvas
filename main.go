package main

import "canvas/canvas"

func main() {
	c := canvas.New(12, 12)
	c.DrawPoint(12, 12, "s")
	c.Draw()
}
