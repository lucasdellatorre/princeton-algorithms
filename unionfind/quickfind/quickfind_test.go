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
			break
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
			break
		}
	}

	t.Log("\nPassed: ", u.ToSlice())
}

func TestConnectionSuccess(t *testing.T) {
	want := true
	u := NewUnionFind(10)
	u.Union(4, 3)
	u.Union(3, 8)
	u.Union(6, 5)
	u.Union(9, 4)
	u.Union(2, 1)

	got := u.Connected(8, 9)

	if got != want {
		t.Errorf("Connected(%d, %d) == %v, want %v", 8, 9, got, want)
	}
}

func TestConnectionFailure(t *testing.T) {
	want := false
	u := NewUnionFind(10)
	u.Union(4, 3)
	u.Union(3, 8)
	u.Union(6, 5)
	u.Union(9, 4)
	u.Union(2, 1)

	got := u.Connected(5, 4)

	if got != want {
		t.Errorf("Connected(%d, %d) == %v, want %v", 8, 9, got, want)
	}
}
