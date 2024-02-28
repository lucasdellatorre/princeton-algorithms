package quickfind

import "fmt"

/*
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
	if u.Connected(p, q) {
		return
	}

	var masterId = u.id[q]
	var wrongId = u.id[p]

	for i := range u.id {
		if u.id[i] == wrongId {
			u.id[i] = masterId
		}
	}
}

func (u *unionFind) Connected(p int, q int) bool {
	return u.id[p] == u.id[q]
}

func (u *unionFind) Show() {
	fmt.Printf("union-find: %v\n", u.id)
}
