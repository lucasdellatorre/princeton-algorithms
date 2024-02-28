package quickfind

import "testing"

func TestUnionEmptyList(t *testing.T) {
	want := []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}
	u := NewUnionFind(10)
	u.Union(4, 3)
	got := u.ToSlice()

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Union(%d, %d) == %v, want %v", 4, 3, got, want)
		}
	}

	t.Log("Passed: ", u.ToSlice())
}

func TestUnion(t *testing.T) {
	want := []int{0, 1, 1, 8, 8, 5, 5, 7, 8, 8}
	u := NewUnionFind(10)
	u.Union(4, 3)
	u.Union(3, 8)
	u.Union(6, 5)
	u.Union(9, 4)
	u.Union(2, 1)

	got := u.ToSlice()

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Union(%d, %d) == %v, want %v", 4, 3, got, want)
		}
	}

	t.Log("\nPassed: ", u.ToSlice())
}
