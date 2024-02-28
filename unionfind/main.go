package main

import (
	"github.com/lucasdellatorre/princeton-algorithms/quickfind"
	"github.com/lucasdellatorre/princeton-algorithms/quickunion"
)

func main() {
	quickFindTest()
	quickUnionTest()
}

func quickFindTest() {
	u := quickfind.NewUnionFind(10)
	u.Union(4, 3)
	u.Union(3, 8)
	u.Union(6, 5)
	u.Union(9, 4)
	u.Union(2, 1)

	u.Show()
}

func quickUnionTest() {
	u := quickunion.NewUnionFind(10)
	u.Union(4, 3)
	u.Union(3, 8)
	u.Union(6, 5)
	u.Union(9, 4)
	u.Union(2, 1)

	u.Show()

}
