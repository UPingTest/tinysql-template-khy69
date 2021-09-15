package main

import "fmt"

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
// 	res := make([][]uint8, dy)
// 	for i := range res {
// 		res[i] = make([]uint8, dx)
// 		for j := range res[i] {
// 			res[i][j] = uint8(i * j)
// 		}
// 	}
// 	return res
// }
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func main() {
	// pic.Show(Pic)
	// delete(m, "sfr")
	// fmt.Println(m)
	// // 	If key is in m, ok is true. If not, ok is false.

	// // If key is not in the map, then elem is the zero value for the map's element type.
	// elem, present := m["sfr"]
	// fmt.Println(elem, present)
	v := Vertex{1, 2}
	v.add(Vertex{2, 4})
	fmt.Println(v)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
// built-in type need to using another name
type Myfloat float64
type Vertex struct {
	X int
	Y int
}

// modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Vertex) add(w Vertex) {
	v.X += w.X
	v.Y += w.Y
}

var m = map[string]Vertex{
	"khy": {10, 5},
	"sfr": {6, 11},
}

// var(
// 	v1=Vertex{1,2}
// 	v2=Vertex{X:1}
// 	v3=&v1

// )
// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func main() {
// v := Vertex{1, 2}
// p := &v
// p.X = 3
// fmt.Println(v.X)
// array:[n]T,index:0 to size-1
// a := [2]string{"hello", "world"}
// fmt.Println(a)
// primes := [6]int{2, 3, 5, 7, 11, 13}
// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array
// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// var s []int = primes[1:4]
// fmt.Println(s)
// s[0] = 1
// 	When slicing, you may omit the high or low bounds to use their defaults instead.

// The default is zero for the low bound and the length of the slice(array) for the high bound.
// s := []struct {
// 	i int
// 	b bool
// }{
// 	{1, true},
// 	{2, false},
// }
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// The zero value of a slice is nil
// has no underlying array.
// a := make([]int, 5)
// a = append(a, 6)
// fmt.Println(a)
// two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index
// 	for i, j := range pow {
// 		fmt.Println(i, j)
// 	}
// }
