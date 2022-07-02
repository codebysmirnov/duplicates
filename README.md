# Duplicates

Simple function set for work with duplicates.

## Install

```bash
go install github.com/codebysmirnov/duplicates@latest
```

## Requirements

1. Go1.18 (because functions use generics)
2. Use comparable types

## Usage

```go
package main

import (
	"fmt"

	"duplicates/duplicates"
)

func example() {
	type example struct {
		Name string
		Age  int
	}

	e1 := example{
		Name: "Josh",
		Age:  18,
	}
	e2 := example{
		Name: "John",
		Age:  18,
	}
	e3 := example{
		Name: "John",
		Age:  18,
	}
	e4 := example{
		Name: "Mark",
		Age:  28,
	}
	e5 := example{
		Name: "Mark",
		Age:  28,
	}

	examples := []example{e1, e2, e3, e4, e5}

	if duplicates.Detect(examples) {
		dups := duplicates.GetAll(examples)
		fmt.Println("Duplicates:", dups)
		count := duplicates.Count(examples)
		fmt.Println("Count:", count)
		cleaned := duplicates.Remove(examples)
		fmt.Println("Cleaned slice from duplicates:", cleaned)
	}
}

func main() {
	example()
}

```