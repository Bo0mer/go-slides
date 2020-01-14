package main

import (
	"fmt"
	"unsafe"
)

type iface struct {
	itab *struct {
		inter unsafe.Pointer
		_type unsafe.Pointer
		hash  uint32
		_     [4]byte
		fun   [1]uintptr
	}
	dat unsafe.Pointer
}

// START1 OMIT
func print(x interface{}) {
	v := *(*iface)(unsafe.Pointer(&x))
	fmt.Println(v.itab, v.dat)
}

func main() {
	var x interface{}
	print(x)
	var bytes []byte // nil
	x = bytes
	print(x)
	x = nil
	print(x)
}

// END1 OMIT
