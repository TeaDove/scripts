package cgodev

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random() + 2)
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
