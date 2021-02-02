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

	if fill != "" {
		c.fulfillRectangle(x, y, width, height, fill)
	}

	if outline != "" {
		c.boundRectangle(x, y, width, height, outline)
	}

	return nil
}

func (c *Canvas) boundRectangle(x, y, width, height int, outline string) {
	xEnd, yEnd := x+width-1, y+height-1
	xCurrent := x
	for xCurrent <= xEnd {
		c.drawPoint(xCurrent, y, outline)
		c.drawPoint(xCurrent, yEnd, outline)
		xCurrent++
	}

	yCurrent := y
	for yCurrent <= yEnd {
		c.drawPoint(x, yCurrent, outline)
		c.drawPoint(xEnd, yCurrent, outline)
		yCurrent++
	}
}

func (c *Canvas) fulfillRectangle(x, y, width, height int, fill string) {
	yCurrent := y
	xEnd, yEnd := x+width, y+height

	for yCurrent < yEnd {
		xCurrent := x
		for xCurrent < xEnd {
			c.drawPoint(xCurrent, yCurrent, fill)
			xCurrent++
		}
		yCurrent++
	}
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
