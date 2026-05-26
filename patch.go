package gomonkey

import (
	"reflect"
	"unsafe"
)

type Patches struct {
	originals    map[uintptr][]byte
	targets      map[uintptr]uintptr
	values       map[reflect.Value]reflect.Value
	valueHolders map[reflect.Value]reflect.Value
}

type Params []interface{}
type OutputCell struct {
	Values Params
	Times  int
}

func ApplyFunc(target, double interface{}) *Patches { _ = "STUB: not implemented"; return nil }

func ApplyMethod(target interface{}, methodName string, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyMethodFunc(target interface{}, methodName string, doubleFunc interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyPrivateMethod(target interface{}, methodName string, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyGlobalVar(target, double interface{}) *Patches { _ = "STUB: not implemented"; return nil }

func ApplyFuncVar(target, double interface{}) *Patches { _ = "STUB: not implemented"; return nil }

func ApplyFuncSeq(target interface{}, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyMethodSeq(target interface{}, methodName string, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyFuncVarSeq(target interface{}, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyFuncReturn(target interface{}, output ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyMethodReturn(target interface{}, methodName string, output ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func ApplyFuncVarReturn(target interface{}, output ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func create() *Patches { _ = "STUB: not implemented"; return nil }

func NewPatches() *Patches { _ = "STUB: not implemented"; return nil }

func (this *Patches) Origin(fn func()) { _ = "STUB: not implemented"; return }

func (this *Patches) ApplyFunc(target, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyMethod(target interface{}, methodName string, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyMethodFunc(target interface{}, methodName string, doubleFunc interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyPrivateMethod(target interface{}, methodName string, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyGlobalVar(target, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyFuncVar(target, double interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyFuncSeq(target interface{}, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyMethodSeq(target interface{}, methodName string, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyFuncVarSeq(target interface{}, outputs []OutputCell) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyFuncReturn(target interface{}, returns ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyMethodReturn(target interface{}, methodName string, returns ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyFuncVarReturn(target interface{}, returns ...interface{}) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) Reset() { _ = "STUB: not implemented"; return }

func (this *Patches) ApplyCore(target, double reflect.Value) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) ApplyCoreOnlyForPrivateMethod(target unsafe.Pointer, double reflect.Value) *Patches {
	_ = "STUB: not implemented"
	return nil
}

func (this *Patches) check(target, double reflect.Value) { _ = "STUB: not implemented"; return }

func replace(target, double uintptr) []byte { _ = "STUB: not implemented"; return nil }

func getDoubleFunc(funcType reflect.Type, outputs []OutputCell) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func GetResultValues(funcType reflect.Type, results ...interface{}) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

type funcValue struct {
	_ uintptr
	p unsafe.Pointer
}

func getPointer(v reflect.Value) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func entryAddress(p uintptr, l int) []byte { _ = "STUB: not implemented"; return nil }

func pageStart(ptr uintptr) uintptr { _ = "STUB: not implemented"; return 0 }

func funcToMethod(funcType reflect.Type, doubleFunc interface{}) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func castRType(val interface{}) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
