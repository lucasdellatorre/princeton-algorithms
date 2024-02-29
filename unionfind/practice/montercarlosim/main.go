package main

import (
	"fmt"
)

// Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty. A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites. If there is a full site in the bottom row, then we say that the system percolates.
// https://introcs.cs.princeton.edu/java/24percolation/

type UnionFind struct {
	id   [][]int
	sz   [][]int
	grid [][]int
}

func main() {
	u := newUnionFind(10)
	printMatrix(u.id)
	fmt.Println()
	printMatrix(u.sz)
	fmt.Println()
	printMatrix(u.grid)

}

func newUnionFind(N int) *UnionFind {
	id := make([][]int, N)
	sz := make([][]int, N)
	grid := make([][]int, N)
	for i := range grid {
		grid[i] = make([]int, N)
		sz[i] = make([]int, N)
		id[i] = make([]int, N)
		for j := range sz[i] {
			sz[i][j] = 1
			id[i][j] = i*10 + j
		}
	}
	return &UnionFind{id: id, sz: sz, grid: grid}
}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Print("[")
		for j := range matrix[i] {
			fmt.Printf(" %d ", matrix[i][j])

		}
		fmt.Print("]\n")
	}

}
