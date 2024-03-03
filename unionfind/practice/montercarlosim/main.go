package main

import (
	"fmt"
	"math/rand"
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
	n    int
}

func main() {
	// New(rand.NewSource(time.Now().UnixNano()))
	n := 4
	u := newUnionFind(n)

	for i := 0; i < n; i++ {
		u.union(i, n*n)
	}

	for i := n * (n - 1); i < n*n; i++ {
		u.union(i, n*n+1)
	}

	randomNumbers := generateUniqueRandomNumbers(n * n)

	fmt.Println(randomNumbers)

	for i := 0; i < n*n; i++ {
		u.Open(randomNumbers[i])
		u.printUf()
	}
}

func generateUniqueRandomNumbers(n int) []int {
	randomNumbers := make([]int, n)
	usedNumbers := make(map[int]bool)

	for i := 0; i < n; i++ {
		var num int
		for {
			num = rand.Intn(n)
			if !usedNumbers[num] {
				break
			}
		}
		randomNumbers[i] = num
		usedNumbers[num] = true
	}

	return randomNumbers
}

func (u *UnionFind) isOpen(i int) bool {
	return !u.grid[i]
}

func (u *UnionFind) Open(i int) {
	u.grid[i] = false
	if i-1 > 0 && u.isOpen(i-1) { // left
		u.union(i, i-1)
	}
	if i+1 < len(u.grid) && u.isOpen(i+1) { // right
		u.union(i, i+1)
	}
	if i+u.n < len(u.grid) && u.isOpen(i+u.n) { // top
		u.union(i, i+1)
	}
	if i-u.n > 0 && u.isOpen(i-u.n) { // left
		u.union(i, i-1)
	}
	// u.union(i, i-u.n) // top
	// u.union(i, i+u.n) // bottom
}

func (u *UnionFind) union(p int, q int) {
	i := u.root(p)
	j := u.root(q)

	if i == j {
		return
	}

	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]

	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]

	}
}

func (u *UnionFind) root(i int) int {
	for i != u.id[i] {
		u.id[i] = u.id[u.id[i]] // path compression
		i = u.id[i]
	}
	return i
}

func (u *UnionFind) connection(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UnionFind) percolates() bool {
	return u.connection(4*4, 4*4+1)
}

func newUnionFind(n int) *UnionFind {
	id := make([]int, n*n+2)
	sz := make([]int, n*n+2)
	grid := make([]bool, n*n)

	initializeIdMatrix(id)
	initializeSizeMatrix(sz)
	initializeGridMatrix(grid)

	uf := &UnionFind{
		id:   id,
		sz:   sz,
		grid: grid,
		n:    n,
	}
	return uf
}

func initializeIdMatrix(matrix []int) {
	for i := range matrix {
		matrix[i] = i
	}
}

func initializeSizeMatrix(matrix []int) {
	for i := range matrix {
		matrix[i] = 1
	}
}

func initializeGridMatrix(matrix []bool) {
	for i := range matrix {
		matrix[i] = true
	}
}

func (u *UnionFind) printIdMatrix() {
	for i := 0; i < u.n*u.n; i++ {
		if i > 0 && i%u.n == 0 {
			fmt.Println()
		}
		fmt.Printf(" %d ", u.id[i])
	}
}

func (u *UnionFind) printSizeMatrix() {
	for i := 0; i < u.n*u.n; i++ {
		if i > 0 && i%u.n == 0 {
			fmt.Println()
		}
		fmt.Printf(" %d ", u.sz[i])
	}
}

func (u *UnionFind) printGridMatrix() {
	for i := 0; i < u.n*u.n; i++ {
		if i > 0 && i%u.n == 0 {
			fmt.Println()
		}
		fmt.Printf(" %t ", u.grid[i])
	}
}

func (u *UnionFind) printUf() {
	u.printIdMatrix()
	fmt.Println()
	u.printSizeMatrix()
	fmt.Println()
	u.printGridMatrix()
	fmt.Println()
	fmt.Println("Percolates?", u.percolates())
}
