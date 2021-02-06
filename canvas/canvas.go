package canvas

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

const (
	canvasWidth  = 12
	canvasHeight = 12
)

type Canvas struct {
	field [canvasHeight][canvasWidth]string
}

func New() *Canvas {
	return &Canvas{
		field: [canvasHeight][canvasWidth]string{},
	}
}

// DrawRectangle draws a rectangle
// x, y - coordinates for the upper-left corner
// width, height, an optional fill and outline characters
// one of either fill or outline should always be present
func (c *Canvas) DrawRectangle(x, y, width, height int, fill, outline string) error {
	if err := c.validatePointPosition(x, y); err != nil {
		return errors.Wrap(err, "invalid rectangle start point")
	}

	if err := c.validateRectanglePosition(x, y, width, height); err != nil {
		return errors.Wrap(err, "invalid rectangle position")
	}

	if err := c.validateRectangleFilling(fill, outline); err != nil {
		return errors.Wrap(err, "invalid rectangle filling")
	}

	if fill != "" {
		if err := c.validateIsSingleASCII(fill); err != nil {
			return errors.Wrap(err, "invalid rectangle fill symbol")
		}
		c.fulfillRectangle(x, y, width, height, fill)
	}

	if outline != "" {
		if err := c.validateIsSingleASCII(outline); err != nil {
			return errors.Wrap(err, "invalid rectangle outline symbol")
		}
		c.boundRectangle(x, y, width, height, outline)
	}

	return nil
}

// FloodFill a flood fill operation
// parameterised with x, y - start coordinates and newSym - fill character
func (c *Canvas) FloodFill(x, y int, newSym string) error {
	if err := c.validatePointPosition(x, y); err != nil {
		return errors.Wrap(err, "failed to pick symbol at start position")
	}

	if err := c.validateIsSingleASCII(newSym); err != nil {
		return errors.Wrap(err, "invalid symbol for filling")
	}

	c.floodFill(x, y, c.pickPointSymbol(x, y), newSym)
	return nil
}

func (c *Canvas) floodFill(x, y int, prevSym, newSym string) {
	if x < 0 || x >= canvasWidth || y < 0 || y >= canvasHeight {
		return
	}
	if c.field[y][x] != prevSym || c.field[y][x] == newSym {
		return
	}

	c.drawPoint(x, y, newSym)
	c.floodFill(x, y+1, prevSym, newSym)
	c.floodFill(x, y-1, prevSym, newSym)
	c.floodFill(x+1, y, prevSym, newSym)
	c.floodFill(x-1, y, prevSym, newSym)
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

func (c *Canvas) pickPointSymbol(x, y int) string {
	return c.field[y][x]
}

func (c *Canvas) drawPoint(x, y int, symbol string) {
	c.field[y][x] = symbol
}

func (c *Canvas) Render(w io.Writer) error {
	result := ""
	for _, rows := range c.field {
		result += "|"
		for x, sym := range rows {
			if sym == "" {
				result += " "
			} else {
				result += sym
			}
			if x != canvasWidth-1 {
				result += "|"
			}
		}
		result += "|\n"
	}

	_, err := fmt.Fprint(w, result)
	return errors.Wrap(err, "failed to render canvas")
}
