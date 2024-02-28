package quickunion

import "fmt"

/*
algorithm: quick union (lazy approach)
Initialize: N
Union: N
find: 1

defect: Union is to expensive O(n^2)
*/

type unionFind struct {
	id []int
}

func NewUnionFind(N int) *unionFind {
	u := &unionFind{id: make([]int, N)}
	for i := range u.id {
		u.id[i] = i
	}
	return u
}

func (u *unionFind) Union(p int, q int) {
	var pRoot = u.id[p]

	println(pRoot)
	println(u.id[pRoot])

	for pRoot != u.id[pRoot] {
		pRoot = u.id[pRoot]
	}

	var qRoot = u.id[q]

	for qRoot != u.id[qRoot] {
		qRoot = u.id[qRoot]
	}

	u.id[pRoot] = qRoot
}

func (u *unionFind) Connected(p int, q int) bool {
	var pRoot = u.id[p]

	for pRoot != u.id[pRoot] {
		pRoot = u.id[pRoot]
	}

	var qRoot = u.id[p]

	for qRoot != u.id[qRoot] {
		qRoot = u.id[qRoot]
	}

	return qRoot == pRoot
}

func (u *unionFind) Show() {
	fmt.Printf("union-find: %v\n", u.id)
}
