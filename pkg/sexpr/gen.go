package sexpr

import (
	"math/rand"
)

// Returns a number from [i, j]
func between(i, j int) int {
	return rand.Intn(j-i+1) + i
}

// Preset is a struct enumerating the options required to generate a new Fn. A
// Preset must be used to create new Fns.
type Preset struct {
	// MaxDepth defines the maximum depth of the parse tree. Avoid large values
	// for MaxDepth, as the size of the tree will be O(2^MaxDepth).
	MaxDepth int
	// NumRange is an inclusive range of random numbers that can be generated.
	NumRange [2]int
	// ArrayLenRange is an inclusive range of possible lengths for array atom
	// lengths.
	ArrayLenRange [2]int
	// OddsBranch is the probability that any node is an operator instead of an
	// atom (operators are "branches" as they have children, the operands.)
	OddsBranch float64
	// OddsVar is the probability that any given atom is the time variable.
	OddsVar float64
}

// Generate generates a new Fn based on the provided options.
func (o Preset) Generate() Fn {
	// Simple probability to check if we are going to branch
	branch := rand.Float64() < o.OddsBranch

	// We are inserting a var object if we are either not branching because of
	// probability or from maxdepth, and of course if the odds are in our favor.
	isVar := (!branch || o.MaxDepth == 1) && rand.Float64() < o.OddsVar

	if o.MaxDepth == 1 || !branch {
		// Either base case (max depth) or we just want an atom.
		if isVar {
			return varAtom{}
		} else {
			return numAtom(between(o.NumRange[0], o.NumRange[1]))
		}
	} else {
		// Recursive case, generate a node and its children
		n := &node{op: OPS[between(0, len(OPS)-1)]}

		// Crank down the recursion level before recursing
		o.MaxDepth -= 1

		// Generate left
		if n.op == "i" {
			// The index operator requires an array to index.
			atom := make(arrayAtom, between(o.ArrayLenRange[0], o.ArrayLenRange[1]))
			for i := 0; i < len(atom); i++ {
				atom[i] = between(o.NumRange[0], o.NumRange[1])
			}
			n.left = atom
		} else {
			n.left = o.Generate()
		}

		// Generate right
		n.right = o.Generate()

		return n
	}
}
