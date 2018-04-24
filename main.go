package main

import (
	"github.com/mesb/mchess/board"
)

func main() {
	b := board.InitBoard()
	board.Printer(*b)
}
