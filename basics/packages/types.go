/**
 * Go Types
 */

// bool
//
// string
//
// int   int8    int16   int32   int64
// uint  uint8   uint16  uint32  uint64  uintptr
//
// byte  // alias for uint8
//
// rune  // alias for int32
//       // represents a Unicode code point
//
// float32 float64
//
// complex64 complex128

/**
 * Go Zero Values
 */

// 0     for numeric
// false for boolean
// ""    for strings

package main

import (
	"fmt"
	"math/cmplx"
)

// block variable declaration
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const format = "%T(%v)\n"

	var i int = 42
	var f float64
	var b bool
	var s string

	v := 42

	fmt.Printf(format, ToBe, ToBe)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

	fmt.Printf("\n%v, %v, %v, %q\n", i, f, b, s)

	// type conversions
	flt := float64(i)

	fmt.Println(flt)

	fmt.Printf("v is of type %T\n", v)
}
