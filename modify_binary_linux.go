package gomonkey

func modifyBinary(target uintptr, bytes []byte) { _ = "STUB: not implemented"; return }

func mprotectCrossPage(addr uintptr, length int, prot int) error {
	_ = "STUB: not implemented"
	return nil
}
