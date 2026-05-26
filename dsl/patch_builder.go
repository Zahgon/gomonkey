package dsl

import (
	. "github.com/agiledragon/gomonkey/v2"
)

type FuncPara struct {
	target      interface{}
	constraints []Constraint
	behaviors   []Behavior
}

type PatchBuilder struct {
	patches  *Patches
	funcPara FuncPara
}

func NewPatchBuilder(patches *Patches) *PatchBuilder { _ = "STUB: not implemented"; return nil }

func (this *PatchBuilder) Func(target interface{}) *PatchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (this *PatchBuilder) Stubs() *PatchBuilder { _ = "STUB: not implemented"; return nil }

func (this *PatchBuilder) With(matcher ...Constraint) *PatchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (this *PatchBuilder) Will(behavior Behavior) *PatchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (this *PatchBuilder) Then(behavior Behavior) *PatchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (this *PatchBuilder) End() { _ = "STUB: not implemented"; return }
