package main

import "testing"

type testcase struct {
	input    string
	expected int
}

type captcha func(string) int

func runTest(t *testing.T, cases []testcase, f captcha, fname string) {
	for _, c := range cases {
		got := f(c.input)
		if got != c.expected {
			t.Errorf("%s(%s) == %d, expected %d", fname, c.input, got, c.expected)
		}
	}

}

func TestPart1(t *testing.T) {
	cases := []testcase{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	runTest(t, cases, captchaP1, "captchaP1")
}

func TestPart2(t *testing.T) {
	cases := []testcase{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	runTest(t, cases, captchaP2, "captchaP2")
}
