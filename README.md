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

import (
	"fmt"
	"log"

	"github.com/dmarichuk/tymbol"
)

func main() {
	// Get some data in column-oriented representation
	data := [][]interface{}{{1, 2, 3}, {"Bob", "Alice", "Francis"}, {10.0, 9.88, 5.002}}

	// Creates table
	table, err := tymbol.NewTable("Players' scores", []string{"id", "Name", "Status"}, data)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Draw a table
	result := table.Draw()
	fmt.Println(result)
	/*
				Players' scores
		#==============#==============#==============#
		#      id      #     Name     #    Status    #
		#==============#==============#==============#
		|      1       |     Bob      |      10      |
		+--------------+--------------+--------------+
		|      2       |    Alice     |     9.88     |
		+--------------+--------------+--------------+
		|      3       |   Francis    |   5.00002    |
		+--------------+--------------+--------------+
	*/

	// Resets table
	table.ResetCanvas()

	// tymbol have some reasonable defaults but you can change them

	// Change some symbols in table
	table.Options.SetHHeaderSym('o')     // Header horizontal symbol
	table.Options.SetCrossHeaderSym('@') // Header cross symbol
	table.Options.SetVHeaderSym('&')     // Header vertical symbol
	table.Options.SetVLineSym('!')       // Line vertical symbol

	// Set aligns
	err = table.Options.SetTitleAlign("left") // returns error if not valid
	table.Options.SetHeaderAlign("right")
	table.Options.SetCellAlign("right")

	table.Options.SetCellLength(10) // Increase fixed length
	table.Options.SetCellPadding(0) // ..and turn off cell padding

	result = table.Draw()
	fmt.Println(result)
	/*
		Players' scores
		@oooooooooo@oooooooooo@oooooooooo@
		&        id&      Name&    Status&
		@oooooooooo@oooooooooo@oooooooooo@
		!         1!       Bob!        10!
		+----------+----------+----------+
		!         2!     Alice!      9.88!
		+----------+----------+----------+
		!         3!   Francis!     5.002!
		+----------+----------+----------+
	*/

	// By default cell has fixed length 10 and padding 2
	// But if you want it compact it has fit content option
	// It will use the longest string in column as fixed length
	table, err = tymbol.NewTable("Players' scores", []string{"id", "Name", "Status"}, data)
	if err != nil {
		log.Fatal(err)
		return
	}
	table.Options.SetCellFitContent(true)

	result = table.Draw()
	fmt.Println(result)
	/*
				Players' scores
		#======#===========#==========#
		#  id  #   Name    #  Status  #
		#======#===========#==========#
		|  1   |    Bob    |    10    |
		+------+-----------+----------+
		|  2   |   Alice   |   9.88   |
		+------+-----------+----------+
		|  3   |  Francis  |  5.002   |
		+------+-----------+----------+
	*/

	// But if you have large blobs of text or long float numbers, probably you
	// would like to use fixed length. It can divide data on lines
	data = [][]interface{}{{1, 2, 3}, {"From fairest creatures we desire increase", "When forty winters shall besiege thy brow", "Look in thy glass, and tell the face thou viewest"}}
	table, err = tymbol.NewTable("Shakespear Sonnets ", []string{"id", "Sonnets"}, data)
	if err != nil {
		log.Fatal(err)
		return
	}
	table.Options.SetCellLength(14)
	table.Options.SetCellAlign("left")

	result = table.Draw()
	fmt.Println(result)

	/*
			   Shakespear Sonnets
		#==================#==================#
		#        id        #     Sonnets      #
		#==================#==================#
		|                  |  From fairest c  |
		|  1               |  reatures we de  |
		|                  |  sire increase   |
		+------------------+------------------+
		|                  |  When forty win  |
		|  2               |  ters shall bes  |
		|                  |  iege thy brow   |
		+------------------+------------------+
		|                  |  Look in thy gl  |
		|  3               |  ass, and tell   |
		|                  |  the face thou   |
		|                  |  viewest         |
		+------------------+------------------+
	*/
}
```
