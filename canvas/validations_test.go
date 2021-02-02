package canvas

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatePointPosition(t *testing.T) {
	cases := []struct {
		width, height int
		x, y          int
		expected      error
	}{
		{
			width:    1,
			height:   1,
			x:        0,
			y:        1,
			expected: errors.New("invalid x coordinate"),
		},
		{
			width:    1,
			height:   1,
			x:        1,
			y:        0,
			expected: errors.New("invalid y coordinate"),
		},
		{
			width:    3,
			height:   3,
			x:        4,
			y:        2,
			expected: errors.New("invalid x coordinate"),
		},
		{
			width:    3,
			height:   3,
			x:        2,
			y:        4,
			expected: errors.New("invalid y coordinate"),
		},
		{
			width:    3,
			height:   3,
			x:        3,
			y:        3,
			expected: nil,
		},
		{
			width:    3,
			height:   3,
			x:        1,
			y:        1,
			expected: nil,
		},
		{
			width:    10,
			height:   5,
			x:        10,
			y:        5,
			expected: nil,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			canv := New(c.width, c.height)
			err := canv.validatePointPosition(c.x, c.y)
			if err != nil {
				require.Error(t, c.expected)
				assert.Equal(t, c.expected.Error(), err.Error())
			} else {
				require.Nil(t, c.expected)
			}
		})
	}
}

func TestValidateRectanglePosition(t *testing.T) {
	canv := New(10, 10)
	cases := []struct {
		x, y          int
		width, height int
		expected      error
	}{
		{
			x:        1,
			y:        1,
			width:    0,
			height:   2,
			expected: errors.New("invalid rectangle width"),
		},
		{
			x:        1,
			y:        1,
			width:    2,
			height:   0,
			expected: errors.New("invalid rectangle height"),
		},
		{
			x:        1,
			y:        1,
			width:    1,
			height:   2,
			expected: errors.New("invalid rectangle width"),
		},
		{
			x:        1,
			y:        1,
			width:    2,
			height:   1,
			expected: errors.New("invalid rectangle height"),
		},
		{
			x:        1,
			y:        1,
			width:    11,
			height:   5,
			expected: errors.New("width of the rectangle overlaps the canvas"),
		},
		{
			x:        1,
			y:        1,
			width:    5,
			height:   11,
			expected: errors.New("height of the rectangle overlaps the canvas"),
		},
		{
			x:        5,
			y:        5,
			width:    7,
			height:   3,
			expected: errors.New("width of the rectangle overlaps the canvas"),
		},
		{
			x:        5,
			y:        5,
			width:    3,
			height:   7,
			expected: errors.New("height of the rectangle overlaps the canvas"),
		},
		{
			x:        5,
			y:        5,
			width:    6,
			height:   2,
			expected: nil,
		},
		{
			x:        5,
			y:        5,
			width:    2,
			height:   6,
			expected: nil,
		},
		{
			x:        1,
			y:        1,
			width:    10,
			height:   10,
			expected: nil,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := canv.validateRectanglePosition(c.x, c.y, c.width, c.height)
			if err != nil {
				require.Error(t, c.expected)
				assert.Equal(t, c.expected.Error(), err.Error())
			} else {
				require.Nil(t, c.expected)
			}
		})
	}
}
