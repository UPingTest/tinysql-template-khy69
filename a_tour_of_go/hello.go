package main

import (
	"math/cmplx"
)

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last
// func add(x, y int) int {
// 	return x + y
// }
// func swap(x, y string) (string, string) {
// 	return y, x
// }
// A return statement without arguments returns the named return values.
// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }
var (
	maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// overflow
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// In Go, a name is exported if it begins with a capital letter
func main() {
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// var i = 1
	// i := 1
	// If an initializer is present, the type can be omitted
	// in Go assignment between items of different type requires an explicit conversion.
	// i := 6.32
	// j := int(i)
	// fmt.Println(Big)
	// s := "hello"
	// fmt.Printf("%q", s)

	// Constants cannot be declared using the := syntax.
}
