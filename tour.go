/**
 * Packages
 *
 * Every Go program is made up of packages.
 * Programs start running in package 'main'.
 * By convention, the package name is the same as the last element of the import path.
 * - For instance, the "math/rand" package comprises files that begin with the statement package rand.
 */
package main

import (
  "fmt"
  "math"
  "math/rand"
  "math/cmplx"
)

/**
 * Functions
 *
 * A function can take zero or more arguments.
 * In this example, 'add' takes two parameters of type 'int'.
 * Notice that the type comes after the variable name.
 */
func add(x int, y int) int {
  return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func addShort(x, y int) int {
  return x + y
}

/**
 * Functions - Multiple Results
 *
 * A function can return any number of results.
 */
func swap(x, y string) (string, string) {
  return y, x
}

/**
 * Functions - Named return values
 *
 * Go's return values may be named. If so, they are treated as variables defined at the top of the function.
 * These names should be used to document the meaning of the return values.
 * A return statement without arguments returns the named return values. This is known as a "naked" return.
 * Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
 */
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}


/**
 * Main
 */
func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Printf("Another number %d over here\n", rand.Intn(10))

  /**
   * Exported Names
   *
   * In Go, a name is exported if it begins with a capital letter.
   * - For example, 'Pizza' is an exported name, as is 'Pi', which is exported from the 'math' package.
   * - 'pizza' and 'pi' do not start with a capital letter, so they are not exported.
   */
  fmt.Printf("PI - %v\n", math.Pi)

  /**
   * Functions (examples)
   */
  fmt.Printf("add(42, 18) = %v\n", add(42, 18))
  hello, world := swap("hello", "world")
  fmt.Printf("swap('hello', 'world') = %v %v\n", hello, world)


  /**
   * Variables
   *
   * The var statement declares a list of variables; as in function argument lists, the type is last.
   * A var statement can be at package or function level.
   */
  var python, java bool
  var a, b int = 1, 2 // initialize multiple variables
  var c, d = 3, 4 // type can be omitted - variable will take the type of initializer
  e, f := "e", false // short-hand variable declaration (CAN ONLY BE USED INSIDE FUNCTIONS; IMPLICT TYPE (no type))
  fmt.Printf("a=%v, b=%v, c=%v, d=%v, e=%v, f=%v\n", a, b, c, d, e, f)
  fmt.Printf("python=%v, java=%v\n", python, java)

  /**
   * Basic types
   *
   * bool
   * string
   *
   * int  int8  int16  int32  int64
   * uint uint8 uint16 uint32 uint64 uintptr
   *
   * byte // alias for uint8
   * rune // alias for int32, * represents a Unicode code point
   *
   * float32 float64
   * complex64 complex128
   *
   */
  var (
    ToBe    bool        = false
    MaxInt  uint64      = 1<<64 - 1
    z       complex128  = cmplx.Sqrt(-5 + 12i)
  )

  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)

  /**
   * Variables - Zero values
   *
   * Variables declared without an explicit initial value are given their zero value.
   * The zero value is:
   * - 0 for numeric types,
   * - false for the boolean type, and
   * - "" (the empty string) for strings.
   */
  var ii int
  var ff float64
  var bb bool
  var ss string
  fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)

  /**
   * Variables - Type conversions (casting)
   *
   * The expression T(v) converts the value v to the type T.
   * - i := 42
   * - f := float64(i)
   * - u := uint(f)
   *
   * Unlike in C, in Go assignment between items of different type requires an explicit conversion.
   */
  var aaa, bbb int = 3, 4
  var fff float64 = float64(aaa / bbb)
  fmt.Println(aaa, bbb, fff)
}
