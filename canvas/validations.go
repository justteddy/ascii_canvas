package canvas

import (
	"unicode"

	"github.com/pkg/errors"
)

func (c *Canvas) validateIsSingleASCII(s string) error {
	if len(s) != 1 || s[0] > unicode.MaxASCII {
		return errors.New("only single ASCII symbols are usable")
	}
	return nil
}

func (c *Canvas) validatePointPosition(x, y int) error {
	if x < 0 || x >= canvasWidth {
		return errors.New("invalid x coordinate")
	}
	if y < 0 || y >= canvasHeight {
		return errors.New("invalid y coordinate")
	}
	return nil
}

func (c *Canvas) validateRectanglePosition(x, y, width, height int) error {
	if width <= 1 {
		return errors.New("invalid rectangle width")
	}
	if height <= 1 {
		return errors.New("invalid rectangle height")
	}

	if x+width > canvasWidth {
		return errors.New("width of the rectangle overlaps the canvas")
	}
	if y+height > canvasHeight {
		return errors.New("height of the rectangle overlaps the canvas")
	}

	return nil
}

func (c *Canvas) validateRectangleFilling(fill, outline string) error {
	if fill == "" && outline == "" {
		return errors.New("one of either `fill` or `outline` should not be empty")
	}

	return nil
}
