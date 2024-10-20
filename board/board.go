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

func (b *Board) DownShift(shift int) { b.LeftShift(shift * b.Size) }

func (b *Board) LeftShift(shift int) {
	if shift >= 64 {
		b.leftShiftArray(shift / 64)
	}
	b.leftShiftMod(shift % 64)
}

func (b *Board) leftShiftArray(shift int) {
	for shift > 0 {
		b.Black[0], b.Black[1], b.Black[2], b.Black[3], b.Black[4], b.Black[5] = 0, b.Black[0], b.Black[1], b.Black[2], b.Black[3], b.Black[4]
		b.White[0], b.White[1], b.White[2], b.White[3], b.White[4], b.White[5] = 0, b.White[0], b.White[1], b.White[2], b.White[3], b.White[4]
		shift -= 1
	}
}

func (b *Board) leftShiftMod(shift int) {
	var leftover uint64 = 0
	for index, element := range b.Black {
		b.Black[index] = (element << shift) + leftover
		leftover = element >> (64 - shift)
	}
	leftover = 0
	for index, element := range b.White {
		b.White[index] = (element << shift) + leftover
		leftover = element >> (64 - shift)
	}
}

func (b *Board) UpShift(shift int) { b.RightShift(shift * b.Size) }

func (b *Board) RightShift(shift int) {
	if shift >= 64 {
		b.rightShiftArray(shift / 64)
	}
	b.rightShiftMod(shift % 64)
}

func (b *Board) rightShiftArray(shift int) {
	for shift > 0 {
		b.Black[0], b.Black[1], b.Black[2], b.Black[3], b.Black[4], b.Black[5] = b.Black[1], b.Black[2], b.Black[3], b.Black[4], b.White[5], 0
		b.White[0], b.White[1], b.White[2], b.White[3], b.White[4], b.White[5] = b.White[1], b.White[2], b.White[3], b.White[4], b.White[5], 0
		shift -= 1
	}
}

func (b *Board) rightShiftMod(shift int) {
	var leftover uint64 = 0
	for index, element := range slices.Backward(b.Black[:]) {
		b.Black[index] = (element >> shift) + leftover
		leftover = element << (64 - shift)
	}
	leftover = 0
	for index, element := range slices.Backward(b.White[:]) {
		b.White[index] = (element >> shift) + leftover
		leftover = element << (64 - shift)
