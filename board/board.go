package board

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Size  int
	Black [6]uint64
	White [6]uint64
}

func (b Board) hasBlack(i, j int) bool {
	point := b.Size*i + j
	index := point / 64
	shift := uint(point % 64)
	return b.Black[index]&(1<<shift) != 0
}

func (b Board) hasWhite(i, j int) bool {
	point := b.Size*i + j
	index := point / 64
	shift := uint(point % 64)
	return b.White[index]&(1<<shift) != 0
}

func (b Board) isEmpty(i, j int) bool {
	return !b.hasBlack(i, j) && !b.hasWhite(i, j)
}

func (b Board) pointIcon(i, j int) string {
	switch {
	case b.hasWhite(i, j):
		return "●"
	case b.hasBlack(i, j):
		return "○"
	default:
		return "·"
	}
}

func (b Board) Print() {
	fmt.Println(b)

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Print(" " + b.pointIcon(i, j) + " ")
		}
		fmt.Println()
	}
}

func (b Board) RandomEmpty() (int, int) {
	i, j := rand.Intn(b.Size), rand.Intn(b.Size)
	for !b.isEmpty(i, j) {
		i, j = rand.Intn(b.Size), rand.Intn(b.Size)
	}
	return i, j
}

func (b *Board) PlaceWhite(i, j int) {
	point := b.Size*i + j
	index := point / 64
	shift := uint(point % 64)
	b.White[index] = b.White[index] | (1 << shift)
}

func (b *Board) PlaceBlack(i, j int) {
	point := b.Size*i + j
	index := point / 64
	shift := uint(point % 64)
	b.Black[index] = b.Black[index] | (1 << shift)
}
