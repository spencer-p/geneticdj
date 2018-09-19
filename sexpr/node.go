package sexpr

import (
	"fmt"
	"io"
	"math"
	"strings"
)

var (
	// Most of these are standard operations. "i" is the index operator for
	// arrays. This slice is to be used as reference and not be mutated. Go does
	// not support slice constants.
	OPS = []string{"^", "|", "&", "%", "/", "*", "+", "-", "<<", ">>", "i"}
)

// node is any place in a function that branches. It has has an operator, which
// it performs on its left and right operands.
type node struct {
	op          string
	left, right Fn
}

func (n *node) String() string {
	var buf strings.Builder
	n.WriteBytes(&buf)
	return buf.String()
}

func (n *node) WriteBytes(buf io.Writer) {
	// Writes "(op left right)"
	fmt.Fprintf(buf, "(%s ", n.op)
	n.left.WriteBytes(buf)
	fmt.Fprint(buf, " ")
	n.right.WriteBytes(buf)
	fmt.Fprint(buf, ")")
}

func (n *node) Eval(t int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// The index operation is a special case, so we'll handle it first
	if n.op == "i" {
		arr, ok := n.left.(arrayAtom)
		if !ok {
			panic("Index operation does not have an array as left parameter")
		}

		// Normalize the index
		index := n.right.Eval(t)
		if index < 0 {
			if index == math.MinInt64 {
				// Unbelievably, multiplying the MinInt64 with -1 is a noop. So
				// we manually change this specific value. Note that this is
				// because 2's complement integers actually cannot represent
				// the additive complement of the min.
				index = math.MaxInt64
			} else {
				index *= -1
			}
		}
		index = index % len(arr)

		// Return the result
		return arr.Index(index)
	}

	// Other cases
	left, right := n.left.Eval(t), n.right.Eval(t)
	switch n.op {
	case "^":
		return left ^ right
	case "|":
		return left | right
	case "&":
		return left & right

	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right

	// For bitshifts, we need a natural number to shift by. If the right is not
	// a natural number, the shift becomes a no-op.
	case "<<":
		if right < 0 {
			return left
		}
		return left << uint(right)
	case ">>":
		if right < 0 {
			return left
		}
		return left >> uint(right)

	// For the division based ops, we check for div by zero errors. Div by zero
	// essentially becomes a no-op.
	case "/":
		if right == 0 {
			return left
		}
		return left / right
	case "%":
		if right == 0 {
			return left
		}
		return left % right
	default:
		panic(fmt.Sprint("Unrecognized operator: ", n.op))
	}
}
