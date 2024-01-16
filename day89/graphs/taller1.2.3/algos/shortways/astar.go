package shortways

import (
	"container/heap"
)

func (d *DijkstraAlg) RelaxAStar(q Queue, target int) {
	v := q.Source

	for _, a := range d.newGraph.Mesh.AdjacentsOf(v) {
		w := a.Oposed(v)

		if d.DistA[w] > d.DistA[v]+a.Weight() {
			manhattan := ManhattanDistance(d.newGraph.Mesh, a)
			d.DistA[w] = d.DistA[v] + a.Weight()
			d.EdgeA[w] = a
			d.newGraph.RelaxCounter(ASTAR)
			if ok, indx := d.Queue.Has(w); ok {
				d.Queue[indx].Dist = d.DistA[w] + manhattan
				heap.Fix(&d.Queue, indx)
			} else {
				heap.Push(&d.Queue, Queue{Source: w, Dist: d.DistA[w] + manhattan})
			}
		}

		if w == target {
			d.TargetFound = true
			break
		}

	}
}

func RunAStar(graph *Graph, origin int, target int) {
	d := NewDijkstra(graph.Mesh, graph)

	heap.Init(&d.Queue)
	d.EdgeA[origin] = nil
	d.DistA[origin] = 0
	heap.Push(&d.Queue, Queue{Source: origin, Dist: 0})
	for d.Queue.Len() > 0 && !d.TargetFound {
		v := heap.Pop(&d.Queue)
		d.RelaxAStar(v.(Queue), target)
	}
}
