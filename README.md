# vdf

A Parser for [Valves Data Format (known as vdf)](https://developer.valvesoftware.com/wiki/KeyValues) written Go. 
Comments are not preserved.

## Installation

It is go gettable

```
$ go get github.com/perplex/vdf
```
   

## Usage

```go
package main

import (
	"fmt"
	"github.com/perplex/vdf"
)

func main() {

	obj, err := vdf.ParseFile("path/to/example.vdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}

```

## Inspiration

This is based on the original fork of this repo from [andygrunwald](https://github.com/andygrunwald/vdf) and 
[simple-vdf](https://github.com/rossengeorgiev/vdf-parser) which could handle duplicate keys and various other nuances 
in vdf files.
