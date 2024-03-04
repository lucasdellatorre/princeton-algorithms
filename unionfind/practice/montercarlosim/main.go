package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty. A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites. If there is a full site in the bottom row, then we say that the system percolates.
// https://introcs.cs.princeton.edu/java/24percolation/
// https://coursera.cs.princeton.edu/algs4/assignments/percolation/specification.php
// Todo: animacao dos nodos se conectando

/*
Recently, i was doing a course of algorithms 1 of the university of princeton and there is a exercise about union find that i would like to share with you guys.

But instead of java i implemented in golang, a language that i'm learning.

Union find is a data structure that we use to connect sets.

Union find can be use to resolve various problems, like Koshen-Kopeman algorithm in physics, Kruskal's minimin spanning three, Equivalence of finite state automata, but the exercise that the course introduces is Percolation.

We will use Monte Carlo simulation to study a natural model known as percolation.

Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty. A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites. If there is a full site in the bottom row, then we say that the system percolates.

I used a weighted union find with path compression and i create a animation do demonstrate this simulation and i would like to share with you guys: .

*/

type UnionFind struct {
	id   []int
	sz   []int
	grid []bool
	n    int
}

type Cell struct {
	Rect  pixel.Rect
	Color color.Color
}

func main() {
	pixelgl.Run(run)
}

func run() {
	var width float64 = 800
	var height float64 = 800
	cfg := pixelgl.WindowConfig{
		Title:  "Percolation",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	n := 5
	u := newUnionFind(n)

	randomNumbers := generateUniqueRandomNumbers(n * n)

	animationCells := make([]Cell, n*n)

	var PADDING_X float64 = 100
	var PADDING_Y float64 = height - 100
	var CELL_SIZE float64 = 50

	for !win.Closed() {
		for i := 0; !u.percolates(); i++ {
			win.Clear(colornames.Lightblue)
			time.Sleep(time.Microsecond * 200)
			u.Open(randomNumbers[i])

			var minX, maxX = PADDING_X - CELL_SIZE, PADDING_X
			var minY, maxY = PADDING_Y + CELL_SIZE, PADDING_Y + CELL_SIZE + CELL_SIZE
			for i := range u.grid {
				if u.grid[i] {
					animationCells[i].Color = colornames.Black
				} else {
					animationCells[i].Color = colornames.White
				}

				if i > 0 && i%n == 0 {
					minX = PADDING_X
					maxX = PADDING_X + CELL_SIZE
					minY = minY - CELL_SIZE
					maxY = maxY - CELL_SIZE
				} else {
					minX = minX + CELL_SIZE
					maxX = maxX + CELL_SIZE
				}
				animationCells[i].Rect = pixel.R(minX, minY, maxX, maxY)
			}

			for i := range animationCells {
				imd := imdraw.New(nil)
				imd.Color = animationCells[i].Color
				imd.Push(animationCells[i].Rect.Min, animationCells[i].Rect.Max)
				imd.Rectangle(0)
				imd.Draw(win)
			}
			fmt.Println("======")
			u.printGridMatrix()
			fmt.Println(animationCells)
			win.Update()
		}
		u.printUf()
		time.Sleep(time.Second * 10)
		win.SetClosed(true)
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

	// Check and join with adjacent open sites
	if i-1 >= 0 && i%u.n != 0 && !u.isOpen(i-1) { // left
		u.union(i, i-1)
	}
	if i+1 < len(u.grid) && (i+1)%u.n != 0 && !u.isOpen(i+1) { // right
		u.union(i, i+1)
	}
	if i-u.n >= 0 && u.isOpen(i-u.n) { // top
		u.union(i, i-u.n)
	}
	if i+u.n < len(u.grid) && u.isOpen(i+u.n) { // bottom
		u.union(i, i+u.n)
	}
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
	return u.connection(u.n*u.n, u.n*u.n+1)
}

func newUnionFind(n int) *UnionFind {
	id := make([]int, n*n+2)
	sz := make([]int, n*n+2)
	grid := make([]bool, n*n)

	initializeIdMatrix(id)
	initializeSizeMatrix(sz)
	initializeGridMatrix(grid)

	u := &UnionFind{
		id:   id,
		sz:   sz,
		grid: grid,
		n:    n,
	}

	for i := 0; i < n; i++ {
		u.union(i, n*n)
	}

	for i := n * (n - 1); i < n*n; i++ {
		u.union(i, n*n+1)
	}

	return u
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
