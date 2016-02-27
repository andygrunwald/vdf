package vdf

type Token int

const (
	_ Token = iota

	// Special tokens
	ILLEGAL
	EOF
	WS

	// Literals
	IDENT // Keys or values

	// Misc characters
	CURLY_BRACE_OPEN  // {
	CURLY_BRACE_CLOSE // }
	QUOTATION_MARK    // "
)

var eof = rune(0)
