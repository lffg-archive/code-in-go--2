package main

import (
	"fmt"
)

// There are four kinds of declarations in Go: `var`, `const`, `type` and
// `func`.

// There are PACKAGE declarations, such as this `main` function.
// The names in the package level are visible throughout the whole package. If
// the name is capitalized, it may be read by package consumers.

func main() {
	// ... And LOCAL declarations, such as this variable `name`.
	// Local names are only visible to its scopes.
	name := "Luiz"
	fmt.Println(name)
}
