package sexpr

import "testing"

func expect(t *testing.T, got, expected int, message string) {
	if got != expected {
		t.Error(message+":", "got", got, "but expected", expected)
	}
}

func TestNumAtom(t *testing.T) {
	var atom numAtom = 10
	expect(t, atom.Eval(0), 10, "numAtom eval failed")
}

func TestArrayAtom(t *testing.T) {
	var atom arrayAtom = []int{1, 3, 6}
	expect(t, atom.Index(1), 3, "arrayAtom index fail at i=1")
}

func TestVarAtom(t *testing.T) {
	atom := varAtom{}
	for i := 0; i < 10; i++ {
		expect(t, atom.Eval(i), int(i), "varAtom did not evaluate to t")
	}
}
