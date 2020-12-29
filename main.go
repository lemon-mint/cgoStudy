package main

// #include <stdio.h>
// #include <stdlib.h>
// void cprint(char* s) {
//   printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cPrint("Hello, World!")
}

func cPrint(data string) {
	buf := C.CString(data)
	C.cprint(buf)
	C.free(unsafe.Pointer(buf))
}
