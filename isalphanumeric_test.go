package isalphanumeric

import "testing"

var implementations = []struct {
	Name string
	Fn   func(string) bool
}{
	{Name: "regex", Fn: IsAlphaNumericRegex},
	{Name: "loop", Fn: IsAlphaNumericLoop},
	{Name: "simd", Fn: IsAlphaNumericSIMD},
}

type testCase struct {
	Name  string
	Input string
	Want  bool
	Bench bool
}

func (tc testCase) String() string {
	if tc.Name != "" {
		return tc.Name
	}
	return tc.Input
}

var testCases = []testCase{
	{Name: "16b-valid", Input: "0123456789abcDEF", Want: true, Bench: true},
	{Input: "0123456789abcDEF0123456789abcDEF", Want: true},
	{Input: "0123456789abcDE-", Want: false},
	{Input: "0123456789abcDEö", Want: false},
	{Input: "0123456789abcDE.", Want: false},

	{
		Name:  "1024b-valid",
		Input: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345670123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345670123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345670123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567",
		Want:  true,
		Bench: true,
	},
}

func TestIsAlphaNumeric(t *testing.T) {
	for _, impl := range implementations {
		t.Run(impl.Name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.String(), func(t *testing.T) {
					if got := impl.Fn(tc.Input); got != tc.Want {
						t.Fatalf("case=%q got=%t want=%t", tc.Input, got, tc.Want)
					}
				})
			}
		})
	}
}

func BenchmarkIsAlphaNumeric(b *testing.B) {
	for _, impl := range implementations {
		b.Run(impl.Name, func(b *testing.B) {
			for _, tc := range testCases {
				if !tc.Bench {
					continue
				}

				b.Run(tc.String(), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.Fn(tc.Input)
					}
				})
			}
		})
	}
}
