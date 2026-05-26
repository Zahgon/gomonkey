package dsl

import . "github.com/agiledragon/gomonkey/v2"

type Behavior interface {
	Apply() []Params
}

type ReturnBehavior struct {
	rets   []Params
	params Params
}

func (this *ReturnBehavior) Apply() []Params { _ = "STUB: not implemented"; return nil }

type RepeatBehavior struct {
	rets     []Params
	behavior Behavior
	times    int
}

func (this *RepeatBehavior) Apply() []Params { _ = "STUB: not implemented"; return nil }
