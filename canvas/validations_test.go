package canvas

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsSingleASCII(t *testing.T) {
	cases := []struct {
		s       string
		isASCII bool
	}{
		{
			s:       "g",
			isASCII: true,
		},
		{
			s:       "#",
			isASCII: true,
		},
		{
			s:       "abc",
			isASCII: false,
		},
		{
			s:       "Ж",
			isASCII: false,
		},
	}

	canv := New()
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := canv.validateIsSingleASCII(c.s)
			require.Equal(t, c.isASCII, err == nil)
		})
	}
}

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
			x:        canvasWidth,
			y:        2,
			expected: errors.New("invalid x coordinate"),
		},
		{
			x:        2,
			y:        canvasHeight,
			expected: errors.New("invalid y coordinate"),
		},
		{
			x:        canvasWidth + 1,
			y:        2,
			expected: errors.New("invalid x coordinate"),
		},
		{
			x:        2,
			y:        canvasHeight + 1,
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
			width:    canvasWidth,
			height:   5,
			expected: errors.New("width of the rectangle overlaps the canvas"),
		},
		{
			x:        1,
			y:        1,
			width:    5,
			height:   canvasHeight,
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
			width:    canvasWidth - 1,
			height:   canvasHeight - 1,
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
