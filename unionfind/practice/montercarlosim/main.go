package main

import (
	"fmt"
	"math"
)

// Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty. A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites. If there is a full site in the bottom row, then we say that the system percolates.
// https://introcs.cs.princeton.edu/java/24percolation/
// https://coursera.cs.princeton.edu/algs4/assignments/percolation/specification.php
// Ideia por trás
// Uma matriz toda bloqueada, escolher um posicao aleatoriamente em toda a matriz. tentar abrir ela
// checar se a matriz foi percolada, ou seja, a parte superior esta conectada com a parte inferior
//

type UnionFind struct {
	id   []int
	sz   []int
	grid []bool
}

func main() {
	n := 4
	u := newUnionFind(n*n + 2) // + 2 top and bottom virtual

	u.printUf()
	u.percolates()
}

func Union() {

}

func (u *UnionFind) root(i int) int {
	for i != u.id[i] {
		fmt.Println(u.id)
		u.id[i] = u.id[u.id[i]] // path compression
		i = u.id[i]
	}
	return i
}

func (u *UnionFind) connection(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UnionFind) percolates() bool {
	virtualTop := u.id[len(u.id)-1]
	virtualBottom := u.id[len(u.id)-2]
	return u.connection(virtualBottom, virtualTop)
}

func newUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		id:   make([]int, n),
		sz:   make([]int, n),
		grid: make([]bool, n),
	}

	for i := range uf.id {
		uf.id[i] = i
		uf.sz[i] = 1
		uf.grid[i] = true
	}

	return uf
}

func printIntegerMatrix(matrix []int) {
	n := int(math.Sqrt(float64(len(matrix)))) // remove this for god sake
	for i := range matrix {
		if i > 0 && i%n == 0 {
			fmt.Println()
		}
		fmt.Printf(" %d ", matrix[i])
	}
}

func printBooleanMatrix(matrix []bool) {
	n := int(math.Sqrt(float64(len(matrix)))) // remove this for god sake
	for i := range matrix {
		if i > 0 && i%n == 0 {
			fmt.Println()
		}
		fmt.Printf(" %t ", matrix[i])
	}
}

func (u *UnionFind) printUf() {
	printIntegerMatrix(u.id)
	fmt.Println()
	printIntegerMatrix(u.sz)
	fmt.Println()
	printBooleanMatrix(u.grid)
	fmt.Println()
	u.percolates()
}
