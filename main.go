package main

// #include "mylib.h"
import "C"

func main() {
	name := C.CString("User")
	C.greet(name)
}
