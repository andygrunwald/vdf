package vdf_test

import (
	"strings"
	"testing"

	"github.com/andygrunwald/vdf"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok vdf.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: vdf.EOF},
		{s: `#`, tok: vdf.Illegal, lit: `#`},
		{s: ` `, tok: vdf.WS, lit: " "},
		{s: "\t", tok: vdf.WS, lit: "\t"},
		{s: "\n", tok: vdf.WS, lit: "\n"},
		{s: "\r", tok: vdf.WS, lit: "\r"},

		// Misc characters
		{s: `{`, tok: vdf.CurlyBraceOpen, lit: "{"},
		{s: `}`, tok: vdf.CurlyBraceClose, lit: "}"},
		{s: `"`, tok: vdf.QuotationMark, lit: "\""},
		{s: `\`, tok: vdf.EscapeSequence, lit: "\\"},

		// Identifiers
		{s: `foo`, tok: vdf.Ident, lit: `foo`},
		{s: `Zx12_3U_-`, tok: vdf.Ident, lit: `Zx12_3U_`},
	}

	for i, tt := range tests {
		s := vdf.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan(false)
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
