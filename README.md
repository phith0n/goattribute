# goattribute

`goattribute` is a lightweight Go library that allows you to set (and get) attributes of a struct dynamically, using dot notation (e.g., `a.b.c`). This can be particularly useful when dealing with dynamic or runtime configuration settings, JSON manipulation, or other similar use cases.

Although most of the code for this project is generated by [GPT4](https://chat.openai.com/), it has full unit tests and can be used safely.

## Features

- Set attributes of a struct using dot notation
- Supports nested structs and pointer types
- Easy to use with a clean and simple API

## Installation

To install `goattribute`, use go get:

```shell
go get github.com/phith0n/goattribute
```

## Usage

Here is a simple example of how to use `goattribute`:

```go
package main

import "github.com/phith0n/goattribute"

type inputTestStruct struct {
	Std  int8   `json:"std"`
	Name string `json:"name"`
}

type outputTestStruct struct {
	Filename string `json:"filename"`
}

type configTestStruct struct {
	Name   string             `json:"name"`
	Input  *inputTestStruct   `json:"input"`
	Output []outputTestStruct `json:"output"`
}

func main() {
	var config configTestStruct
	var attr = goattribute.New(&config)
	attr.SetAttr("Name", "hello")
	attr.SetAttr("Input.Name", "world")
	attr.SetAttr("Input.Std", 2)
	attr.SetAttr("Output[0].Filename", "test.txt")
}
```

## Caveats

The library uses Go's reflection (reflect package) to manipulate the struct fields. Therefore, the performance might not be as good as direct field access. Please be cautious when using this library in performance-critical applications.

Currently, the library does not support slice or map types for attributes. Support for these types might be added in future releases.

This project will convert int-based variable automatically, take care of the accuracy and range by yourself.

## License

`goattribute` is released under the MIT License. See the [LICENSE](LICENSE) file for more information.