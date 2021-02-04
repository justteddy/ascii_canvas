package canvas

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Marshal serializes field to JSON
func (c *Canvas) Marshal() ([]byte, error) {
	data, err := json.Marshal(&c.field)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal canvas data")
	}

	return data, nil
}

// Unmarshal deserializes field from JSON and returns new Canvas
func Unmarshal(data []byte) (*Canvas, error) {
	c := New()
	if err := json.Unmarshal(data, &c.field); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal data to canvas")
	}

	return c, nil
}
