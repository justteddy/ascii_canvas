package canvas

import (
	"fmt"
)

type Canvas struct {
	height int
	width  int
	field  [][]string
}

func New(width, height int) *Canvas {
	field := make([][]string, 0, height)
	for i := 0; i < width; i++ {
		columns := make([]string, width)
		field = append(field, columns)
	}

	return &Canvas{
		height: height,
		width:  width,
		field:  field,
	}
}

func (c *Canvas) DrawPoint(x, y int, symbol string) {
	c.field[y-1][x-1] = symbol
}

func (c *Canvas) Draw() {
	result := ""
	for _, rows := range c.field {
		result += "|"
		for x, sym := range rows {
			if sym == "" {
				result += " "
			} else {
				result += sym
			}
			if x != c.width-1 {
				result += "|"
			}
		}
		result += "|\n"
	}
	fmt.Print(result)
}
