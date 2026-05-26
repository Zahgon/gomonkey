package dsl

type Constraint interface {
	Eval(x interface{}) bool
}

type AnyConstraint struct {
}

func (this *AnyConstraint) Eval(x interface{}) bool { _ = "STUB: not implemented"; return false }

type EqConstraint struct {
	x interface{}
}

func (this *EqConstraint) Eval(x interface{}) bool { _ = "STUB: not implemented"; return false }
