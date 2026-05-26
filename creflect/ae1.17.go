//go:build go1.17
// +build go1.17

package creflect

// name is an encoded type name with optional extra data.
type name struct {
	bytes *byte
}

func (n name) data(off int, whySafe string) *byte { _ = "STUB: not implemented"; return nil }

func (n name) readVarint(off int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (n name) name() (s string) { _ = "STUB: not implemented"; return "" }
