package example

import (
	"bytes"
	"fmt"
)

func Pointers() {
	var x int = 5
	var p *int

	p = &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p:", p)
	fmt.Println("Value pointed by p:", *p)

	*p = 10
	fmt.Println("Modified value of x:", x)

	modifyValue(p)
	fmt.Println("Value after modifyValue:", x)
}

func modifyValue(p *int) {
	*p = 15
}

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

// func ByteSlice() {
// 	var buffer [256]byte
// 	slice := sliceHeader{
// 		Length:        50,
// 		ZerothElement: &buffer[100],
// 	}
// 	print(slice)
// }

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func AddArray() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func DelOne() {
	var buffer [256]byte
	slice := buffer[10:20]
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func Truncate() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)
}
