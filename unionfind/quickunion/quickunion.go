package quickunion

/*
algorithm: quick union (lazy approach)
Initialize: N
Union: N (cost of finding roots)
find: N (worst-case)

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

func (u *UnionFind) root(i int) int {
	for i != u.id[i] {
		i = u.id[i]
	}
	return i
}

func (u *UnionFind) Union(p int, q int) {
	i := u.root(p)
	j := u.root(q)
	u.id[i] = j
}

func (u *UnionFind) Connected(p int, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UnionFind) ToSlice() []int {
	return u.id
}
