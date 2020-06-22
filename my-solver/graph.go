package main

import (
	"sort"
)

type Edge struct {
	in int
	out int
	weight float64
}

type Edges []Edge

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].weight < edges[j].weight
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func makeEdges(dist map[int]map[int]float64) Edges {
	var edges Edges
	for i := 0; i < len(dist); i++ {
		for j := i + 1; j < len(dist); j++ {
			edges = append(edges, Edge{
				in: i,
				out: j,
				weight: dist[i][j],
			})
			edges = append(edges, Edge{
				in: j,
				out: i,
				weight: dist[i][j],
			})
		}
	}
	return edges
}

func sortEdges(edges Edges) Edges {
	sort.Sort(edges)
	return edges
}

type Graph struct {
	edges Edges
	nodes []int
}