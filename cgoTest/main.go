package main

/*
#cgo CXXFLAGS: -O3 -DNDEBUG
#cgo LDFLAGS: -L${SRCDIR}/../build/lib/
#cgo LDFLAGS:  -lcppLib -lkernel32 -luser32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -lstdc++
#include "../lib/cppLib.h"
*/
import "C"

func main() {
	C.test()
}
