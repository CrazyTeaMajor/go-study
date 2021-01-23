// otherTest
package main

import "unsafe"

const (
	a = "Myname"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func test() {
	println(a, b, c)
}
