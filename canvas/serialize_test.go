package canvas

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerialize(t *testing.T) {
	t.Run("marshal empty canvas", func(t *testing.T) {
		c := New()
		actual, err := c.Marshal()
		require.NoError(t, err)

		expected, err := ioutil.ReadFile("testdata/empty_canvas.json")
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("marshal canvas with rectangle and fill", func(t *testing.T) {
		c := New()
		err := c.DrawRectangle(4, 4, 4, 5, "", "*")
		require.NoError(t, err)

		err = c.FloodFill(0, 0, "-")
		require.NoError(t, err)

		actual, err := c.Marshal()
		require.NoError(t, err)

		expected, err := ioutil.ReadFile("testdata/rectangle_and_fill.json")
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("unmarshal canvas with rectangle and fill", func(t *testing.T) {
		data, err := ioutil.ReadFile("testdata/rectangle_and_fill.json")
		require.NoError(t, err)

		c, err := Unmarshal(data)
		require.NoError(t, err)

		upperLeftCorner := c.pickPointSymbol(4, 4)
		upperRightCorner := c.pickPointSymbol(7, 4)
		lowerLeftCorner := c.pickPointSymbol(4, 8)
		lowerRightCorner := c.pickPointSymbol(7, 8)

		// check the corners of the rectangle
		assert.Equal(t, "*", upperLeftCorner)
		assert.Equal(t, "*", upperRightCorner)
		assert.Equal(t, "*", lowerLeftCorner)
		assert.Equal(t, "*", lowerRightCorner)

		// check that the inside of the rectangle is empty
		rectangleInside := c.pickPointSymbol(5, 5)
		assert.Equal(t, "", rectangleInside)

		// check filled area
		filled := c.pickPointSymbol(1, 1)
		assert.Equal(t, "-", filled)

	})
}
