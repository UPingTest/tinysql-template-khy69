package main

import (
	"fmt"
)

// func:function or function pointer
func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// for{

	// }
	// while(true)
	// Go only runs the selected case, not all the cases that follow
	// Go's switch cases need not be constants, and the values involved need not be integers
	// Switch without a condition is the same as switch true.
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)
	// }
	// calculate arguments now,call it when the surrounding function returns
	// a series of "defer statement",call them in the reverse order
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("world")
	defer fmt.Println("sfr")

	fmt.Println("hello")
}
