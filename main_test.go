package main

import "testing"

func TestIsBipartite(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true}, // even cycle
		{3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, false},        // odd cycle
		{2, [][2]int{{0, 1}}, true},
		{5, [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, false}, // odd cycle in 2nd component
		{3, nil, true}, // no edges: trivially bipartite
		{0, nil, true},
	}
	for _, c := range cases {
		if got := IsBipartite(c.n, c.edges); got != c.want {
			t.Errorf("IsBipartite(n=%d, %v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
