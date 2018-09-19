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
	Eval(int) int
	String() string
	WriteBytes(io.Writer)
}
