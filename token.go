package vdf

// Token is our own type that represents a single token
// to work with during parsing
type Token int

const (
	_ Token = iota

	// ILLEGAL represents a token that we
	// don`t know in contect of the VDF format
	ILLEGAL

	// EOF represents the End of File token.
	// This is used if the file is end
	EOF

	// WS represents a whitespace.
	// This can be a space or a tab.
	WS

	// IDENT represents a key or a value.
	// Typically this is a simple string
	IDENT // Keys or values

	// CURLY_BRACE_OPEN represents a open curly brace "{"
	CURLY_BRACE_OPEN

	// CURLY_BRACE_CLOSE represents a close curly brace "}"
	CURLY_BRACE_CLOSE

	// QUOTATION_MARK represents a quote mark '"'
	QUOTATION_MARK

	// ESCAPE_SEQUENCE represents an escape character "\"
	ESCAPE_SEQUENCE
)

var eof = rune(0)
