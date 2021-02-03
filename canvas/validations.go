package canvas

import "github.com/pkg/errors"

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
	if err := c.validatePointPosition(x, y); err != nil {
		return err
	}

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
