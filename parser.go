package vdf

import (
	"bytes"
	"fmt"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan(respectWhitespace bool) (Token, string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit := p.s.Scan(respectWhitespace)

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (Token, string) {
	tok, lit := p.scan(false)
	if tok == WS {
		tok, lit = p.scan(false)
	}
	return tok, lit
}

// Parse is the main entry point of the vdf parser.
// If parsed the complete VDF content and returns
// a map as a key / value pair.
// The value is a string (normal value) or a map[string]interface{}
// again if there is a nested structure.
func (p *Parser) Parse() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	key := ""

	// The first part is a simple ident (as a main map key)
	tok, lit := p.scanMapKey()
	if tok != Ident {
		return nil, fmt.Errorf("Found %q, expected an ident as a first part", lit)
	}
	key = lit

	tok, lit = p.scanIgnoreWhitespace()
	if tok != CurlyBraceOpen {
		return nil, fmt.Errorf("Found %q, expected an ident as a first part", lit)
	}

	p.unscan()
	m[key] = p.parseMap()

	return m, nil
}

func (p *Parser) scanMapKey() (Token, string) {
	tok, lit := p.scanIgnoreWhitespace()

	// Get the key
	if tok == QuotationMark {
		return p.scanIdentSurroundedQuotationMark()
	} else if tok == Ident {
		return tok, lit
	}

	return Illegal, lit
}

func (p *Parser) parseMap() map[string]interface{} {
	m := make(map[string]interface{})
	key := ""

	// The first part should be a open curly brace
	tok, _ := p.scanIgnoreWhitespace()
	if tok != CurlyBraceOpen {
		return m
	}

	for {
		// At first: A key
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case QuotationMark:
			_, key = p.scanIdentSurroundedQuotationMark()
		case Ident:
			key = lit
		case CurlyBraceClose:
			return m
		default:
			return m
		}

		// After this: A value or a map again
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case QuotationMark:
			_, m[key] = p.scanIdentSurroundedQuotationMark()
		case Ident:
			m[key] = lit
		case CurlyBraceOpen:
			p.unscan()
			m[key] = p.parseMap()
		case CurlyBraceClose:
			return m
		default:
			return m
		}
	}
}

func (p *Parser) scanIdentSurroundedQuotationMark() (Token, string) {
	var buf bytes.Buffer
	escaped := false

	for {
		tok, lit := p.scan(true)

		if tok == QuotationMark && escaped == false {
			// We don`t unscan here, because
			// we don`t need this quotation mark anymore.
			break
		}

		// If the current character is escaped and it is NOT an escape sequence
		// set character handling back to normal.
		if escaped == true && tok != EscapeSequence {
			escaped = false
		}

		// If we have an escape sequence and the current state is not escaped
		// mark the next character as escaped.
		if tok == EscapeSequence && escaped == false {
			escaped = true
			continue
		}

		buf.WriteString(lit)

		// If the current character is escaped and it is a backslash
		// reset the character handing.
		// This is only triggered if you want to add a \ into a key or a value.
		// Then you have to add "\\".
		if escaped == true && tok == EscapeSequence {
			escaped = false
		}
	}

	return Ident, buf.String()
}
