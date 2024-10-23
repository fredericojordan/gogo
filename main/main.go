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
	b := board.Board{9, board.Bitboard{211661359156224}, board.Bitboard{252483329525809152}}
	b.Print()
	b = board.Board{9, board.Bitboard{941625864}, board.Bitboard{1031866943494}}
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
