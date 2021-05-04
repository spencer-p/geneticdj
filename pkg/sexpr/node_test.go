package sexpr

import (
	"testing"
)

func TestString(t *testing.T) {
	input := "(* (/ 3 4) (i [3 4] (% (& t 4) 2)))"
	fn, err := Parse(input)
	if err != nil {
		t.Fatal("Failed to parse input expression")
	}

	output := fn.String()
	if input != output {
		t.Error("Input does not match output. Expected\n", input, "\nbut got\n", output)
	}
}
