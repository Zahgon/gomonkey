// Customized reflect package for gomonkey，copy most code from go/src/reflect/type.go

package creflect

import (
	"reflect"
	"unsafe"
)

// rtype is the common implementation of most values.
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, may be zero
}

func Create(t reflect.Type) *rtype { _ = "STUB: not implemented"; return nil }

type funcValue struct {
	_ uintptr
	p unsafe.Pointer
}

func funcPointer(v reflect.Method, ok bool) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

func MethodByName(r reflect.Type, name string) (fn unsafe.Pointer, ok bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

func (t *rtype) Method(p method) (fn unsafe.Pointer) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

type tflag uint8
type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type textOff int32 // offset from top of text section

//go:linkname resolveTextOff reflect.resolveTextOff
func resolveTextOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

func (t *rtype) textOff(off textOff) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

//go:linkname resolveNameOff reflect.resolveNameOff
func resolveNameOff(ptrInModule unsafe.Pointer, off int32) unsafe.Pointer

func (t *rtype) nameOff(off nameOff) name { _ = "STUB: not implemented"; return *new(name) }

const (
	tflagUncommon tflag = 1 << 0
)

// uncommonType is present only for defined types or types with methods
type uncommonType struct {
	pkgPath nameOff // import path; empty for built-in types like int, string
	mcount  uint16  // number of methods
	xcount  uint16  // number of exported methods
	moff    uint32  // offset from this uncommontype to [mcount]method
	_       uint32  // unused
}

// ptrType represents a pointer type.
type ptrType struct {
	rtype
	elem *rtype // pointer element (pointed at) type
}

// funcType represents a function type.
type funcType struct {
	rtype
	inCount  uint16
	outCount uint16 // top bit is set if last input parameter is ...
}

func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// interfaceType represents an interface type.
type interfaceType struct {
	rtype
	pkgPath name      // import path
	methods []imethod // sorted by hash
}

type imethod struct {
	name nameOff // name of method
	typ  typeOff // .(*FuncType) underneath
}

type String struct {
	Data unsafe.Pointer
	Len  int
}

func (t *rtype) uncommon(r reflect.Type) *uncommonType { _ = "STUB: not implemented"; return nil }

// Method on non-interface type
type method struct {
	name nameOff // name of method
	mtyp typeOff // method type (without receiver)
	ifn  textOff // fn used in interface call (one-word receiver)
	tfn  textOff // fn used for normal method call
}

func (t *uncommonType) methods() []method { _ = "STUB: not implemented"; return nil }
