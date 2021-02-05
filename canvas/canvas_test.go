package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDrawRectangles(t *testing.T) {
	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|*|*|*| | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
	*/
	t.Run("draw 1 rectangle without outline", func(t *testing.T) {
		c := New()
		err := c.DrawRectangle(4, 4, 5, 5, "*", "")
		require.NoError(t, err)

		// check corners
		assert.Equal(t, "*", c.pickPointSymbol(4, 4))
		assert.Equal(t, "*", c.pickPointSymbol(4, 8))
		assert.Equal(t, "*", c.pickPointSymbol(8, 4))
		assert.Equal(t, "*", c.pickPointSymbol(8, 8))

		// check inside
		assert.Equal(t, "*", c.pickPointSymbol(5, 6))
		assert.Equal(t, "*", c.pickPointSymbol(6, 6))
		assert.Equal(t, "*", c.pickPointSymbol(7, 6))

		// check outside
		assert.Equal(t, "", c.pickPointSymbol(0, 0))
		assert.Equal(t, "", c.pickPointSymbol(11, 11))
	})

	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | |.|.|.|.|.| | | |
		| | | | |.| | | |.| | | |
		| | | | |.| | | |.| | | |
		| | | | |.| | | |.| | | |
		| | | | |.|.|.|.|.| | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
	*/
	t.Run("draw 1 rectangle without fill", func(t *testing.T) {
		c := New()
		err := c.DrawRectangle(4, 4, 5, 5, "", ".")
		require.NoError(t, err)

		// check corners
		assert.Equal(t, ".", c.pickPointSymbol(4, 4))
		assert.Equal(t, ".", c.pickPointSymbol(4, 8))
		assert.Equal(t, ".", c.pickPointSymbol(8, 4))
		assert.Equal(t, ".", c.pickPointSymbol(8, 8))

		// check inside
		assert.Equal(t, "", c.pickPointSymbol(5, 6))
		assert.Equal(t, "", c.pickPointSymbol(6, 6))
		assert.Equal(t, "", c.pickPointSymbol(7, 6))
	})

	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | |@|@|@| | | | | | | |
		| | |@|X|@| | | | | | | |
		| | |@|@|@| | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | |X|X|X|X|X| |
		| | | | | | |X|0|0|0|X| |
		| | | | | | |X|0|0|0|X| |
		| | | | | | |X|0|0|0|X| |
		| | | | | | |X|X|X|X|X| |
		| | | | | | | | | | | | |
	*/
	t.Run("draw 2 not overlapped rectangles", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(2, 2, 3, 3, "X", "@")
		require.NoError(t, err)

		err = c.DrawRectangle(6, 6, 5, 5, "0", "X")
		require.NoError(t, err)

		// first - check corners
		assert.Equal(t, "@", c.pickPointSymbol(2, 2))
		assert.Equal(t, "@", c.pickPointSymbol(4, 2))
		assert.Equal(t, "@", c.pickPointSymbol(2, 4))
		assert.Equal(t, "@", c.pickPointSymbol(4, 4))

		// first - check inside
		assert.Equal(t, "X", c.pickPointSymbol(3, 3))

		// second - check corners
		assert.Equal(t, "@", c.pickPointSymbol(2, 2))
		assert.Equal(t, "@", c.pickPointSymbol(4, 2))
		assert.Equal(t, "@", c.pickPointSymbol(2, 4))
		assert.Equal(t, "@", c.pickPointSymbol(4, 4))

		// second - check inside
		assert.Equal(t, "X", c.pickPointSymbol(3, 3))
	})

	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|*|*|*| | | |
		| | | | |*|*|.|.|.|.|.| |
		| | | | |*|*|.|*|*| |.| |
		| | | | |*|*|.|*|*| |.| |
		| | | | | | |.| | | |.| |
		| | | | | | |.|.|.|.|.| |
		| | | | | | | | | | | | |
	*/
	t.Run("draw 2 overlapped rectangles", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(4, 4, 5, 5, "*", "")
		require.NoError(t, err)

		err = c.DrawRectangle(6, 6, 5, 5, "", ".")
		require.NoError(t, err)

		// check intersection
		assert.Equal(t, ".", c.pickPointSymbol(6, 6))
		assert.Equal(t, ".", c.pickPointSymbol(7, 6))
		assert.Equal(t, ".", c.pickPointSymbol(8, 6))

		assert.Equal(t, ".", c.pickPointSymbol(6, 7))
		assert.Equal(t, "*", c.pickPointSymbol(7, 7))
		assert.Equal(t, "*", c.pickPointSymbol(8, 7))

		assert.Equal(t, ".", c.pickPointSymbol(6, 8))
		assert.Equal(t, "*", c.pickPointSymbol(7, 8))
		assert.Equal(t, "*", c.pickPointSymbol(8, 8))

		// check inside of both rectangles
		assert.Equal(t, "*", c.pickPointSymbol(5, 5))
		assert.Equal(t, "", c.pickPointSymbol(9, 9))
	})
}

func TestFloodFill(t *testing.T) {
	/*
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
	*/
	t.Run("fill whole canvas", func(t *testing.T) {
		c := New()

		err := c.FloodFill(0, 0, "x")
		require.NoError(t, err)

		assert.Equal(t, "x", c.pickPointSymbol(0, 0))
		assert.Equal(t, "x", c.pickPointSymbol(11, 11))
		assert.Equal(t, "x", c.pickPointSymbol(6, 6))
	})

	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		|.|.|.|.|.|.|.|.|.|.|.|.|
		|.|x|x|x|x|x|x|x|x|x|x|.|
		|.|.|.|.|.|.|.|.|.|.|.|.|
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
	*/
	t.Run("fill area inside of rectangle", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(0, 5, 12, 3, "", ".")
		require.NoError(t, err)

		err = c.FloodFill(1, 6, "x")
		require.NoError(t, err)

		assert.Equal(t, ".", c.pickPointSymbol(0, 6))
		assert.Equal(t, "x", c.pickPointSymbol(1, 6))
		assert.Equal(t, ".", c.pickPointSymbol(11, 6))
	})

	/*
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|x|x|x|x|x|x|x|x|x|x|x|x|
		|.|.|.|.|.|.|.|.|.|.|.|.|
		|.| | | | | | | | | | |.|
		|.|.|.|.|.|.|.|.|.|.|.|.|
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
	*/
	t.Run("fill area outside of rectangle", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(0, 5, 12, 3, "", ".")
		require.NoError(t, err)

		err = c.FloodFill(0, 0, "x")
		require.NoError(t, err)

		// check rectangle
		assert.Equal(t, ".", c.pickPointSymbol(0, 6))
		assert.Equal(t, "", c.pickPointSymbol(1, 6))
		assert.Equal(t, ".", c.pickPointSymbol(11, 6))

		// check fill area
		assert.Equal(t, "x", c.pickPointSymbol(0, 0))
		assert.Equal(t, "x", c.pickPointSymbol(11, 4))

		// check empty area
		assert.Equal(t, "", c.pickPointSymbol(0, 8))
		assert.Equal(t, "", c.pickPointSymbol(11, 11))
	})

	/*
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		|@|@|@|@|@|@|@|@|@|@|@|@|
		|@|x|x|x|x|x|x|x|x|x|x|@|
		|@|@|@|@|@|@|@|@|@|@|@|@|
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
		| | | | | | | | | | | | |
	*/
	t.Run("fill border of rectangle", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(0, 5, 12, 3, "x", ".")
		require.NoError(t, err)

		err = c.FloodFill(0, 5, "@")
		require.NoError(t, err)

		assert.Equal(t, "@", c.pickPointSymbol(0, 6))
		assert.Equal(t, "x", c.pickPointSymbol(1, 6))
		assert.Equal(t, "@", c.pickPointSymbol(11, 6))
		assert.Equal(t, "@", c.pickPointSymbol(11, 5))
		assert.Equal(t, "@", c.pickPointSymbol(11, 7))
	})

	/*
		|-|-|-|-|-|-|-|-|-|-|-|-|
		|-|-|-|-|-|-|-|-|@|@|@|@|
		|-|-|-|-|-|-|-|-|@|@|@|@|
		|-|-|-|-|-|-|-|-|@|@|@|@|
		|-|-|-|-|-|-|-|-|@|@|@|@|
		|.|.|.|.|.|-|-|-|-|-|-|-|
		|.| | | |.|-|-|-|-|-|-|-|
		|.| | | |.|-|-|-|-|-|-|-|
		|.|.|.|.|.|-|-|-|-|-|-|-|
		|-|-|-|-|-|-|-|-|-|-|-|-|
		|-|-|-|-|-|-|-|-|-|-|-|-|
		|-|-|-|-|-|-|-|-|-|-|-|-|
	*/
	t.Run("complex filling area", func(t *testing.T) {
		c := New()

		err := c.DrawRectangle(0, 5, 5, 4, "", ".")
		require.NoError(t, err)

		err = c.DrawRectangle(8, 1, 4, 4, "@", "")
		require.NoError(t, err)

		err = c.FloodFill(0, 0, "-")
		require.NoError(t, err)

		// check fill area
		assert.Equal(t, "-", c.pickPointSymbol(0, 0))
		assert.Equal(t, "-", c.pickPointSymbol(7, 4))
		assert.Equal(t, "-", c.pickPointSymbol(7, 7))
		assert.Equal(t, "-", c.pickPointSymbol(0, 11))
		assert.Equal(t, "-", c.pickPointSymbol(11, 11))

		// check first rectangle
		assert.Equal(t, "", c.pickPointSymbol(1, 6))
		assert.Equal(t, ".", c.pickPointSymbol(0, 5))
		assert.Equal(t, ".", c.pickPointSymbol(0, 8))
		assert.Equal(t, ".", c.pickPointSymbol(4, 5))
		assert.Equal(t, ".", c.pickPointSymbol(4, 8))

		// check second rectangle
		assert.Equal(t, "@", c.pickPointSymbol(9, 2))
		assert.Equal(t, "@", c.pickPointSymbol(8, 1))
		assert.Equal(t, "@", c.pickPointSymbol(8, 4))
		assert.Equal(t, "@", c.pickPointSymbol(11, 1))
		assert.Equal(t, "@", c.pickPointSymbol(11, 4))
	})
}
