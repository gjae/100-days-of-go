package shortways

import (
	"container/heap"
	"gjae/graph-mesh/graph"
	"math"
)

func Coordinates(mesh *graph.GraphMesh, srcVertex int) (int, int) {
	for y := 0; y < mesh.Rows; y++ {
		for x := 0; x < mesh.Cols; x++ {
			currentVertex := (y * mesh.Cols) + x
			if currentVertex == srcVertex {
				return x, y
			}
		}
	}

	return -1, -1
}

func ManhattanDistance(mesh *graph.GraphMesh, edge *graph.Edge) float32 {
	v, w := edge.Origin(), edge.Oposed(edge.Origin())
	srcx, srcy := Coordinates(mesh, v)
	dstx, dsty := Coordinates(mesh, w)

	dx := math.Abs(float64(srcx - dstx))
	dy := math.Abs(float64(srcy - dsty))

	return float32(dx + dy)
}

func (d *DijkstraAlg) RelaxBfs(q Queue, target int) {
	v := q.Source
	for _, a := range d.newGraph.Mesh.AdjacentsOf(v) {
		w := a.Oposed(v)

		if d.EdgeA[w] == nil {
			manhattan := ManhattanDistance(d.newGraph.Mesh, a)
			d.EdgeA[w] = a
			heap.Push(&d.Queue, Queue{
				Source: w,
				Dist:   manhattan,
			})
			d.newGraph.RelaxCounter(BFS)
		}

		if w == target {
			d.TargetFound = true
			break
		}

	}
}

func RunBfs(graph *Graph, origin int, target int) {
	d := NewDijkstra(graph.Mesh, graph)

	heap.Init(&d.Queue)
	d.EdgeA[origin] = nil
	d.DistA[origin] = 0
	heap.Push(&d.Queue, Queue{Source: origin, Dist: 0})

	for d.Queue.Len() > 0 && !d.TargetFound {
		v := heap.Pop(&d.Queue)
		d.RelaxBfs(v.(Queue), target)
	}
}
