package sexpr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO change this to its own error type with a column field?
func errAt(col int, msg string) error {
	return errors.New(fmt.Sprint(msg, " at col ", col))
}

// readNumTo reads out a int from a string, ending the number at whatever is in
// delims. Also returns the width of the int that was read.
func readNumTo(s, delims string) (int, int, error) {
	// Slice string to the delim
	width := strings.IndexAny(s, delims)
	if width == -1 {
		width = len(s)
	}
	s = s[:width]

	// Parse as float
	res, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, 0, err
	}

	return int(res), width, nil
}

// Parse turns an s-expression string into a Fn that can be evaluated or
// otherwise manipulated. An error is returned if the expression is malformed
// for any reason.
func Parse(input string) (Fn, error) {
	stack := make([]Fn, 0)

	for i := 0; i < len(input); i++ {

		switch input[i] {
		case ' ':
		case '\n':
			// Spaces are a no-op.

		case '(':
			// Should be the start of an s-expression.
			i += 1
			if i >= len(input) {
				return nil, errAt(i, "Open paren at end of expression")
			}

			// Get the string for the operator. Operators are always followed by
			// a space.
			width := strings.IndexByte(input[i:], ' ')
			if width == -1 {
				return nil, errAt(i, "Operator not followed by space")
			}

			stack = append(stack, &node{op: input[i : i+width]})
			i += width - 1

		case ')':
			// s-expression closed. make sure we have enought to build it
			if len(stack) < 3 {
				return nil, errAt(i, "Expression closed prematurely")
			}

			// pop the last two elements:
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// And push them onto their parent.
			parent, ok := stack[len(stack)-1].(*node)
			if !ok {
				return nil, errAt(i, "Malformed expression missing parent")
			}
			parent.left = left
			parent.right = right

		case 't':
			// This is the variable of time.
			stack = append(stack, varAtom{})

		case '[':
			// Should be an array literal.
			// Consume the bracket
			i += 1

			// Parse each num
			arr := make(arrayAtom, 0)
			for {
				// Nums in array are delimited by spaces and the ending bracket.
				num, width, err := readNumTo(input[i:], " ]")
				if err != nil {
					return nil, errAt(i, "Failed to parse num in array atom")
				}

				// Add to array
				arr = append(arr, num)

				// Skip by the number
				i += width

				// Consume whitespace and garbage
				for input[i] == ' ' {
					i += 1
				}

				// Close if we hit the end
				if input[i] == ']' {
					break
				}
			}

			stack = append(stack, arr)

		default:
			// Everything else should be a number literal.
			num, width, err := readNumTo(input[i:], " )")
			if err != nil {
				return nil, errAt(i, "Failed to parse number atom")
			}
			stack = append(stack, numAtom(num))
			i += width - 1
		}

	}

	if len(stack) != 1 {
		return nil, errors.New("Multiple expressions at end of parse")
	}

	return stack[0], nil
}
