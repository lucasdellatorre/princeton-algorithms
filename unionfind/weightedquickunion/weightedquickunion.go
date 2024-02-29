package weightedquickunion

import "fmt"

/*
algorithm: weight quick union (lazy approach)
Initialize: N
Union: log N(cost of finding roots)
connected: log N
*/

type UnionFind struct {
	id []int
	sz []int
}

func NewUnionFind(N int) *UnionFind {
	id := make([]int, N)
	sz := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
		sz[i] = 1
	}
	return &UnionFind{id: id, sz: sz}
}

func (u *UnionFind) root(i int) int {
	for i != u.id[i] {
		fmt.Println(u.id)
		u.id[i] = u.id[u.id[i]] // path compression
		i = u.id[i]
	}
	return i
}

func (u *UnionFind) Union(p int, q int) {
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

func (u *UnionFind) Connected(p int, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UnionFind) ToSlice() []int {
	return u.id
}
