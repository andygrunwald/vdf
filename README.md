# vdf

[![GoDoc](https://godoc.org/github.com/andygrunwald/vdf?status.svg)](https://pkg.go.dev/github.com/andygrunwald/vdf)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygrunwald/vdf)](https://goreportcard.com/report/github.com/andygrunwald/vdf)

A Lexer and Parser for [Valves Data Format (known as vdf)](https://developer.valvesoftware.com/wiki/KeyValues) written in Go.

## Installation

It is go gettable

```
$ go get github.com/andygrunwald/vdf@master
```
    
(optional) to run unit / example tests:

```
$ cd $GOPATH/src/github.com/andygrunwald/vdf
$ go test -v ./...
```

## Usage

Given a file named [`gamestate_integration_consolesample.cfg`](testdata/gamestate_integration_consolesample.cfg) with content:

```
"Console Sample v.1"
{
	"uri" 		"http://127.0.0.1:3000"
	"timeout" 	"5.0"
	"buffer"  	"0.1"
	"throttle" 	"0.5"
	"heartbeat"	"60.0"
	[...]
}
```

Can be parsed with this Go code:

```go
package main

import (
	"fmt"
	"os"

	"github.com/andygrunwald/vdf"
)

func main() {
	f, err := os.Open("gamestate_integration_consolesample.cfg")
	if err != nil {
		panic(err)
	}

	p := vdf.NewParser(f)
	m, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}
```

And it will output:

```
map[
	Console Sample v.1:map[
		uri:http://127.0.0.1:3000
		timeout:5.0
		buffer:0.1
		throttle:0.5
		heartbeat:60.0
		[...]
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
