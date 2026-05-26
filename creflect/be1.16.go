//go:build !go1.17
// +build !go1.17

package creflect

// name is an encoded type name with optional extra data.
type name struct {
	bytes *byte
}

func (n name) name() (s string) { _ = "STUB: not implemented"; return "" }
