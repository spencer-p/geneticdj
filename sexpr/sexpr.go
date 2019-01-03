/*
Package sexpr provides a library to parse, create, evaluate, and manipulate
functions of a single variable represented as s-expressions.
*/
package sexpr

import (
	"io"
)

// Fn is any s-expression function that can be evaluated. Atoms and other
// operators implement this interface.
type Fn interface {
	// Eval evaluates the function with the given argument.
	Eval(int) int

	// String returns an s-expression representation of the function.
	String() string

	// WriteBytes writes the ASCII bytes that make up an s-expression
	// representation of the function.
	WriteBytes(io.Writer)
}
