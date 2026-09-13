package main

import "testing"

func TestStringCompress(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "single char", in: "a", want: "a"},
		{name: "all singles", in: "abc", want: "abc"},
		{name: "single pair", in: "aa", want: "2a"},
		{name: "pairs", in: "aabbcc", want: "2a2b2c"},
		{name: "mixed runs", in: "aaaabbcc", want: "4a2b2c"},
		{name: "multi digit run", in: "aaaaaaaaaa", want: "10a"},
		{name: "single then multi digit run", in: "abbbbbbbbbbbb", want: "a12b"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			in := []byte(tc.in)
			got := string(stringCompress(in))
			if got != tc.want {
				t.Fatalf("stringCompress(%q) = %q, want %q", tc.in, got, tc.want)
			}

			if len(got) > len(tc.in) {
				t.Fatalf("stringCompress(%q) grew output: got len %d > input len %d", tc.in, len(got), len(tc.in))
			}
		})
	}
}
