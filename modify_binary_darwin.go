package gomonkey

func PtrOf(val []byte) uintptr { _ = "STUB: not implemented"; return 0 }

func modifyBinary(target uintptr, bytes []byte) { _ = "STUB: not implemented"; return }

//go:cgo_import_dynamic mach_task_self mach_task_self "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic mach_vm_protect mach_vm_protect "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic sys_icache_invalidate sys_icache_invalidate "/usr/lib/libSystem.B.dylib"
func write(target, data uintptr, len int, page uintptr, pageSize, oriProt int) int
