package quickfind

/*
Initialize: N
Union: N
find: 1

defect: Union is to expensive O(n^2)
*/

type UnionFind struct {
	id []int
}

func NewUnionFind(N int) *UnionFind {
	id := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
	}
	return &UnionFind{id: id}
}

func (u *UnionFind) Union(p int, q int) {
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

func (u *UnionFind) Connected(p int, q int) bool {
	return u.id[p] == u.id[q]
}

func (u *UnionFind) ToSlice() []int {
	return u.id
}
