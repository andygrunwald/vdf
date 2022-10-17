package vdf

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: EOF},
		{s: `#`, tok: Illegal, lit: `#`},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},
		{s: "\r", tok: WS, lit: "\r"},

		// Misc characters
		{s: `{`, tok: CurlyBraceOpen, lit: "{"},
		{s: `}`, tok: CurlyBraceClose, lit: "}"},
		{s: `"`, tok: QuotationMark, lit: "\""},
		{s: `\`, tok: EscapeSequence, lit: "\\"},
		{s: "//", tok: CommentDoubleSlash, lit: "//"},

		// Identifiers
		{s: `foo`, tok: Ident, lit: `foo`},
		{s: `Zx12_3U_-`, tok: Ident, lit: `Zx12_3U_`},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan(false)
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

func FuzzScanner_ScanWithoutWhitespace(f *testing.F) {
	testcases := []string{
		// Special tokens (EOF, ILLEGAL, WS)
		``, `#`, ` `, "\t", "\n", "\r",

		// Misc characters
		`{`, `}`, `"`, `\`, "//",

		// Identifiers
		`foo`, `Zx12_3U_-`,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		s := NewScanner(strings.NewReader(input))
		s.Scan(false)
	})
}

func FuzzScanner_ScanWithWhitespace(f *testing.F) {
	testcases := []string{
		// Special tokens (EOF, ILLEGAL, WS)
		``, `#`, ` `, "\t", "\n", "\r",

		// Misc characters
		`{`, `}`, `"`, `\`, "//",

		// Identifiers
		`foo`, `Zx12_3U_-`,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		s := NewScanner(strings.NewReader(input))
		s.Scan(true)
	})
}
