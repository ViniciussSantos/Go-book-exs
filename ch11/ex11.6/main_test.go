package main

import "testing"

// pc[i] is the population count of i.
var pc [256]byte

func initTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int {
	var count int
	for i := uint(0); i < 8; i++ {
		count += int(x >> i & 1)
	}
	return count
}

func PopCountClear(x uint64) int {

	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func BenchmarkPopcountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(uint64(i))
	}
}

func BenchmarkPopcountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(uint64(i))
	}
}

func BenchmarkPopcountTable(b *testing.B) {
	initTable()
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(i))
	}
}
