# vdf

[![GoDoc](https://godoc.org/github.com/andygrunwald/vdf?status.svg)](https://godoc.org/github.com/andygrunwald/vdf)
[![Build Status](https://travis-ci.org/andygrunwald/vdf.svg?branch=master)](https://travis-ci.org/andygrunwald/vdf)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygrunwald/vdf)](https://goreportcard.com/report/github.com/andygrunwald/vdf)
[![Coverage Status](https://coveralls.io/repos/github/andygrunwald/vdf/badge.svg?branch=master)](https://coveralls.io/github/andygrunwald/vdf?branch=master)

A Lexer and Parser for [Valves Data Format (known as vdf)](https://developer.valvesoftware.com/wiki/KeyValues) written in Go.

## Installation

It is go gettable

```
$ go get github.com/andygrunwald/vdf
```
    
(optional) to run unit / example tests:

```
$ cd $GOPATH/src/github.com/andygrunwald/vdf
$ go test -v ./...
```

## Usage

Given a file named `example.vdf` with content:

```
"Example"
{
	"TimeNextStatsReport"      "1234567890"
	"ContentStatsID"           "-7123456789012345678"
}
```

Can be parsed with this Go code:

```go
package main

import (
	"os"
	"log"
	"github.com/andygrunwald/vdf"
)

func main() {
	f, err := os.Open("./example.vdf")
	if err != nil {
		log.Fatal(err)
	}

	p := vdf.NewParser(f)
	m, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
```

And it will output:

```
map[
	Example:map[
		TimeNextStatsReport:1234567890
		ContentStatsID:-7123456789012345678
	]
]
```
## Inspiration

The code and project idea is heavily inspired and driven by [@benbjohnson](https://github.com/benbjohnson) article [Handwritten Parsers & Lexers in Go](https://blog.gopheracademy.com/advent-2014/parsers-lexers/) and his example [sql-parser](https://github.com/benbjohnson/sql-parser). Thank you Ben!

## Parser in other languages

* PHP and JavaScript: [rossengeorgiev/vdf-parser](https://github.com/rossengeorgiev/vdf-parser)
* PHP: [https://github.com/devinwl/keyvalues-php](devinwl/keyvalues-php)
* PHP: [lukezbihlyj/vdf-parser](https://github.com/lukezbihlyj/vdf-parser)
* C#: [sanmadjack/VDF](https://github.com/sanmadjack/VDF)
* Java: [DHager/hl2parse](https://github.com/DHager/hl2parse)
* And many more: [Github search for vdf valve](https://github.com/search?p=1&q=vdf+valve&ref=searchresults&type=Repositories&utf8=%E2%9C%93)
		
## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).