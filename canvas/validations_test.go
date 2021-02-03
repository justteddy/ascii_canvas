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
		x, y     int
		expected error
	}{
		{
			x:        -1,
			y:        1,
			expected: errors.New("invalid x coordinate"),
		},
		{
			x:        1,
			y:        -1,
			expected: errors.New("invalid y coordinate"),
		},
		{
			x:        12,
			y:        2,
			expected: errors.New("invalid x coordinate"),
		},
		{
			x:        2,
			y:        12,
			expected: errors.New("invalid y coordinate"),
		},
		{
			x:        13,
			y:        2,
			expected: errors.New("invalid x coordinate"),
		},
		{
			x:        2,
			y:        13,
			expected: errors.New("invalid y coordinate"),
		},
		{
			x:        0,
			y:        0,
			expected: nil,
		},
		{
			x:        5,
			y:        8,
			expected: nil,
		},
	}

	canvas := New()
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := canvas.validatePointPosition(c.x, c.y)
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
			width:    12,
			height:   5,
			expected: errors.New("width of the rectangle overlaps the canvas"),
		},
		{
			x:        1,
			y:        1,
			width:    5,
			height:   12,
			expected: errors.New("height of the rectangle overlaps the canvas"),
		},
		{
			x:        5,
			y:        5,
			width:    8,
			height:   3,
			expected: errors.New("width of the rectangle overlaps the canvas"),
		},
		{
			x:        5,
			y:        5,
			width:    3,
			height:   8,
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
			width:    11,
			height:   11,
			expected: nil,
		},
	}

	canvas := New()
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := canvas.validateRectanglePosition(c.x, c.y, c.width, c.height)
			if err != nil {
				require.Error(t, c.expected)
				assert.Equal(t, c.expected.Error(), err.Error())
			} else {
				require.Nil(t, c.expected)
			}
		})
	}
}
