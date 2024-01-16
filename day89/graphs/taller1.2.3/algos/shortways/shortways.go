package shortways

import (
	"fmt"
	"gjae/graph-mesh/graph"
)

type Graph struct {
	Mesh     *graph.GraphMesh
	dijkstra int
	bfs      int
	aStar    int
}

const (
	DIJKSTRA = iota
	BFS
	ASTAR
)

func NewShortWayFinder(old *graph.GraphMesh) *Graph {
	return &Graph{
		Mesh: old,
	}
}

func (g *Graph) RelaxCounter(relaxSource int) {
	if relaxSource == DIJKSTRA {
		g.dijkstra++
	} else if relaxSource == BFS {
		g.bfs++
	} else {
		g.aStar++
	}
}

func (g *Graph) Print() {
	entry := -1
	exit := -1

	for k, _ := range g.Mesh.AdjList {
		if len(g.Mesh.AdjacentsOf(k)) > 1 {
			if entry == -1 {
				entry = k
			} else if exit == -1 && entry != k {
				exit = k
			}
		}

		if entry != -1 && exit != -1 {
			break
		}
	}

	if entry > exit {
		entry, exit = exit, entry
	}

	RunDijkstra(g, entry, exit)
	RunBfs(g, entry, exit)
	RunAStar(g, entry, exit)

	fmt.Printf("\nDijkstra: %d\n", g.dijkstra)
	fmt.Printf("Bfs: %d\n", g.bfs)
	fmt.Printf("A*: %d\n", g.aStar)
}
