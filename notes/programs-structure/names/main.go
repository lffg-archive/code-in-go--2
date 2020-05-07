package main

import (
	"fmt"
)

// Name is an example of type, whichare a kind of "name" in Go.
type Name = string

// A variable name is also a kind of "name" in Go.
var name Name = "Luiz"

// Constants are also a kind of name. Here we used `iota` to provide incremented
// values.
const (
	FOO = iota
	BAR
	BAZ
)

// Names can contain any unicode letter:
var áéíóú bool = true

// They can also start with underscores:
var _test string = "This is a test"

// Funcions are also a kind of name.
func main() {
	fmt.Println(name, FOO, BAR, BAZ, áéíóú, _test)
}
