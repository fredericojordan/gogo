package main

import (
	"gogo/board"
	"time"
)

func main() {
	puzzles()
	// randomFill()
}

func puzzles() {
	var b board.Board
	b = board.Board{9, [6]uint64{211661359156224}, [6]uint64{252483329525809152}}
	b.Print()
	b = board.Board{9, [6]uint64{941625864}, [6]uint64{1031866943494}}
	b.Print()
}

func randomFill() {
	b := board.Board{Size: 9}
	for i := 0; i < 40; i++ {
		b.PlaceBlack(b.RandomEmpty())
		b.Print()
		time.Sleep(200 * time.Millisecond)
		b.PlaceWhite(b.RandomEmpty())
		b.Print()
		time.Sleep(200 * time.Millisecond)
	}
}
