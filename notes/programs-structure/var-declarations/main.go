package main

import (
	"fmt"
)

// `var` declarations creates a variable with a certain type, assigns it with a
// custom name and its value.
//
// Each `var` declarations must follow the form:
//
//    var name type = expr
//
// The `type` OR `= expr` may be omitted, but NOT both.
var foo string = "Foo"
var bar string
var baz = "Baz"

// Note that as `= expr` is omitted, the variable will assume the default value
// for its type.
//
// - `0` to numbers
// - `false` to boleans
// - `""` to strings
// - `nil` to other types (such as slices, pointers, maps, channels and functions)
//
// This "zero value" mechanism will ensure that each variable will always have a
// value. So there are no "uninitialized variables" in Go.

func main() {
	fmt.Println("Foo", foo)
	fmt.Println("Foo", bar)
	fmt.Println("Foo", baz)
}
