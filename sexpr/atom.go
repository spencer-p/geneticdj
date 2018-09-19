package sexpr

import (
	"fmt"
	"io"
)

// numAtom is a single number. In the expression (+ 3 4), "3" and "4" are both
// numAtom.
type numAtom int

func (atom numAtom) Eval(t int) int {
	// The time variable never changes a constant, but we need it to satisfy the
	// interface.
	return int(atom)
}

func (atom numAtom) String() string {
	return fmt.Sprint(int(atom))
}

func (atom numAtom) WriteBytes(buf io.Writer) {
	fmt.Fprint(buf, int(atom))
}

// arrayAtom is an array of numbers, such as [1 2 3]. arrayAtom should always
// be owned by an index operator. Note that no commas are used in arrayAtoms.
type arrayAtom []int

func (atom arrayAtom) Eval(t int) int {
	panic("sexpr: attempt to evaluate an arrayAtom")
}

func (atom arrayAtom) String() string {
	return fmt.Sprint([]int(atom))
}

func (atom arrayAtom) WriteBytes(buf io.Writer) {
	fmt.Fprint(buf, atom)
}

func (atom arrayAtom) Index(i int) int {
	if i < 0 || i >= len(atom) {
		panic(fmt.Sprintf("atom has length %d but got index %d", len(atom), i))
	}
	return atom[i]
}

// varAtom is the variable t itself. It always evaluates to itself.
type varAtom struct{}

func (atom varAtom) Eval(t int) int {
	return t
}

func (atom varAtom) String() string {
	return "t"
}

func (atom varAtom) WriteBytes(buf io.Writer) {
	fmt.Fprint(buf, "t")
}
