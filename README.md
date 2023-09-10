# Tymbol

## Short description

Tymbol is a go module designed for creating tables composed of customizable symbols using provided data.

## Installation

```sh
go get -u github.com/dmarichuk/tymbol
```

## Basic usage

```go
package main

import(
    "fmt"
    
    "github.com/dmarichuk/tymbol"
)

func main() {
    data := [][]interface{}{{1, 2, 3}, {"Bob", "Alice", "Francis"}, {✅, ✅, ❌}}

    table, err := tymbol.Table("Players' status", []string{"id", "Name", "Status"})
    if err != nil {
        fmt.Fatalln(err)
    }

    table.Draw()
}
```
