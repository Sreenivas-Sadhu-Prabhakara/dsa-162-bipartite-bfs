# Bipartite Check (BFS)

> Striver SDE Sheet 162 · Day 23 Graph — BFS 2-coloring: levels alternate colors, and any edge inside a level betrays an odd cycle. Go, stdlib only, tested.

## Problem

Can the graph's vertices be split into two sets with every edge crossing between them? Equivalently (König): is the graph free of odd-length cycles?

## Approach

Try to build the 2-coloring directly: color a start vertex, BFS outward, give every newly met neighbor the opposite color (`3 − color`, a branch-free flip between 1 and 2), and fail the moment an edge connects two same-colored vertices. BFS makes the equivalence visible — its levels alternate colors, so a same-color edge is precisely an edge within a level, which closes an odd cycle. Component loop as always: one all-even component says nothing about the next (test pinned). Bipartiteness is what makes matching algorithms and "two-team" partition problems tractable.

**Complexity:** O(V+E) time, O(V) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
