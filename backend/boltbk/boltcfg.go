package boltbk

import (
	"encoding/json"
	"fmt"

	"github.com/gravitational/teleport/backend"
)

// cfg represents JSON config for bolt backlend
type cfg struct {
	Path string `json:"path"`
}

// FromString initialized the backend from backend-specific string
func FromString(v string) (backend.Backend, error) {
	if len(v) == 0 {
		return nil, fmt.Errorf(
			`please supply a valid dictionary, e.g. {"path": "/opt/bolt.db"}`)
	}
	var c *cfg
	if err := json.Unmarshal([]byte(v), &c); err != nil {
		return nil, fmt.Errorf("invalid backend configuration format, err: %v", err)
	}
	return New(c.Path)
}
