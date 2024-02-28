package quickunion

/*
algorithm: quick union (lazy approach)
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
	var pRoot = u.id[p]

	for pRoot != u.id[pRoot] {
		pRoot = u.id[pRoot]
	}

	var qRoot = u.id[q]

	for qRoot != u.id[qRoot] {
		qRoot = u.id[qRoot]
	}

	u.id[pRoot] = qRoot
}

func (u *UnionFind) Connected(p int, q int) bool {
	var pRoot = u.id[p]

	for pRoot != u.id[pRoot] {
		pRoot = u.id[pRoot]
	}

	var qRoot = u.id[q]

	for qRoot != u.id[qRoot] {
		qRoot = u.id[qRoot]
	}

	return qRoot == pRoot
}

func (u *UnionFind) ToSlice() []int {
	return u.id
}
