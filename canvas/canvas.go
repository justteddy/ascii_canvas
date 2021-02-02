package canvas

import (
	"fmt"

	"github.com/pkg/errors"
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

func (c *Canvas) DrawRectangle(x, y, width, height int, fill, outline string) error {
	if err := c.validateRectanglePosition(x, y, width, height); err != nil {
		return errors.Wrap(err, "invalid rectangle position")
	}

	if err := c.validateRectangleFilling(fill, outline); err != nil {
		return errors.Wrap(err, "invalid rectangle filling")
	}

	yCurrent := y
	xEnd, yEnd := x+width, y+height

	//yFrom := start.Y - 1
	//yTo := yFrom + height
	//if yTo > sizeY-1 {
	//	yTo = sizeY - 1
	//}
	//
	//xFrom := start.X - 1
	//xTo := xFrom + width
	//if xTo > sizeX-1 {
	//	xTo = sizeX - 1
	//}

	for yCurrent < yEnd {
		xCurrent := x
		for xCurrent < xEnd {
			c.drawPoint(xCurrent, yCurrent, fill)
			xCurrent++
		}
		yCurrent++
	}

	return nil
}

func (c *Canvas) drawPoint(x, y int, symbol string) {
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
