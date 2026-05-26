package dsl

import . "github.com/agiledragon/gomonkey/v2"

func Any() Constraint { _ = "STUB: not implemented"; return *new(Constraint) }

func Eq(x interface{}) Constraint { _ = "STUB: not implemented"; return *new(Constraint) }

func Return(x ...interface{}) Behavior { _ = "STUB: not implemented"; return *new(Behavior) }

func Repeat(behavior Behavior, times int) Behavior {
	_ = "STUB: not implemented"
	return *new(Behavior)
}
