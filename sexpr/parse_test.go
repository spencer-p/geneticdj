package sexpr

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestReadNumTo(t *testing.T) {
	num, width, _ := readNumTo("-1235$#%789", "$#")
	expect(t, num, -1235, "read incorrect number")
	expect(t, width, 5, "read incorrect number width")
}

func TestNumAtomParse(t *testing.T) {
	fn, err := Parse("3")
	if err != nil {
		t.Error("Failed to parse:", err)
	}
	expect(t, fn.Eval(0), 3, "atom does not parse and eval")
}

func TestArrAtomParse(t *testing.T) {
	fn, err := Parse("[1 3 5]")
	if err != nil {
		t.Error("Failed to parse:", err)
	}

	arr, ok := fn.(arrayAtom)
	if !ok {
		t.Fatal("Parse did not return an arrayAtom for the array input")
	}

	expect(t, arr.Index(0), 1, "arrAtom wrong at index 0")
	expect(t, arr.Index(2), 5, "arrAtom wrong at index 2")
}

func TestAddParse(t *testing.T) {
	fn, err := Parse("(+ 1 2)")
	if err != nil {
		t.Error("Failed to parse:", err)
	}
	expect(t, fn.Eval(0), 3, "could not add 1+2")
}

func TestEvals(t *testing.T) {
	t.Helper()

	tests := []struct {
		input    string
		expected int
	}{
		// Basic math
		{"(+ 1 2)", 3},
		{"(* (+ 1 2) (- 7 4))", 9},
		{"(/ (+ 1 2) (- 7 4))", 1},
		{"(- (* 7 6) (+ 5 4))", 33},
		{"(+ (* 0 2) (* -1 100))", -100},
		{"(<< 1 2)", 4},

		// Array index test
		{"(i [65 2 3] t)", 65},

		// Crash resistance checks
		{"(/ 5 0)", 5},
		{"(% 5 0)", 5},
		{"(<< 1 -2)", 1},
	}

	for _, test := range tests {
		test := test
		t.Run("Parse and Eval: "+test.input, func(t *testing.T) {
			fn, err := Parse(test.input)
			if err != nil {
				t.Error("Could not parse", test.input, "because", err)
			}
			expect(t, fn.Eval(0), test.expected, "Eval failed for "+test.input)
		})
	}
}

func TestDoubler(t *testing.T) {
	// This function should always double t.
	fn, err := Parse("(* t 2)")
	if err != nil {
		t.Error("Failed to parse:", err)
	}
	for i := 0; i < 10; i++ {
		expect(t, fn.Eval(i), int(2*i), "Doubler failed")
	}
}

func TestGenerateReflexivity(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Helper()

	const NUM_TESTS = 10

	for i := 0; i < NUM_TESTS; i++ {
		t.Run(fmt.Sprint("Iteration #", i), func(t *testing.T) {
			randFn := Preset{
				MaxDepth:      5,
				NumRange:      [2]int{0, 10},
				ArrayLenRange: [2]int{2, 10},
				OddsBranch:    0.5,
				OddsVar:       0.25,
			}.Generate()

			asString := randFn.String()
			parsedAgain, err := Parse(asString)
			if err != nil {
				t.Fatal("Could not parse string of fn:", err)
			}

			finalString := parsedAgain.String()

			if asString != finalString {
				t.Error("Parse not reflexive. Got\n", finalString, "\nbut expected\n", asString)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	// Expression to parse
	input := "(>> (% (| (^ (i [-21 -26 45 63 -68 -71 46 -36 26 19 8 -60 -36] t) (| (^ (* (>> (| 81 53) (<< 59 t)) (* (% 100 -19) (| -80 -87))) (& (>> (| 5 -83) (+ -56 t)) (>> (>> -12 15) 56))) (% (>> (- (>> 83 20) (& t 87)) (<< -43 (i [-82 75 -41 21 87 62 -69 73 31 -22 -14 -66 -27 75 -6 93 19 -90 2 -45 96] 35))) (| (<< 52 (+ 12 -45)) (i [31 74 64 50 43 96 16 -46 -19 59 45 83 68 -21 78 -93 17 -76 -3 40 -17 2 -34 54] (^ -3 7)))))) (% (| (>> (<< (<< (& t t) (i [58 96 98 -37 -45 91 -79 38 3 69 -32 -33 7 11 -51 23 -34 61 -67 87 -35 8 -78 91 -3 -79] -41)) (^ (i [-81 -56 92 18 -94 20 62 -33 -29] 70) (* 64 -65))) (i [56 -50 96 47 -94 -47 -77 71 7 27 14 68 -33 91 -53 -46 18 -49 -29] -8)) -31) (/ (i [61 67 -66 -93 -68 76 61 66 -28 78 -100 20 76 -12 0 -80 -31 97 35 36 -19 0 -14] t) (^ 44 t)))) (% (- (* (* (| 33 (i [10 36 43 -73 -51 77 32 25 -96 29 -37 -10 -33 -50 -98 54 -94 -89 -51 6 -79] (/ 62 84))) (& (i [-50 -23 64 -92 76 -81 8 -9 81 74 26 62 88 -76 94 -84 46 32 46] (% 51 t)) 97)) (+ (>> (>> (% t -28) (>> -15 65)) (* t (^ 35 t))) -75)) 15) (<< (<< (/ -14 (* (<< 25 (^ 100 -57)) (>> (i [-38 37 90 -93 -96 9 13 -77 -97 64 -11 48 -69 -26 -48 -7 20 99] 21) -87))) (^ -45 56)) (i [-51 68 72 -55] (& (| 53 (- (| -92 14) (* 85 42))) (| 55 67)))))) (<< 84 (| (& (>> (/ (>> (& -3 (^ -44 t)) 2) (- (| (i [-98 -65 -73 -16 63 -61 2 37 8 -64 -43 24] -6) t) -9)) -38) (* (- (<< (>> (/ 50 89) -63) (/ t -42)) (+ (>> (- -21 79) (+ t t)) (- (& t -39) (/ -69 51)))) -51)) (/ (i [71 44 1 -100 -94 -72 56] t) t))))"

	b.Log(b.N)
	for i := 0; i < b.N; i++ {
		Parse(input)
	}
}
