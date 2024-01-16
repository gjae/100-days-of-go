package graph

import (
	"fmt"
	"math/rand"
	"time"
)

type Edge struct {
	v      int
	w      int
	weight float32
}

type GraphMesh struct {
	AdjList  map[int][]*Edge
	size     int
	Cols     int
	Rows     int
	Directed bool
}

/**
* Crea un objeto grafo, usando una implementación de lista
* de adyacencia usando un map con una clave entera y de valor una lista
* de aristas
 */
func NewGraph(cols, rows int, directed bool) *GraphMesh {

	vTotal := cols * rows
	size := vTotal
	list := make(map[int][]*Edge)
	for i := 0; i < vTotal; i++ {
		list[i] = []*Edge{}
	}

	return &GraphMesh{size: size, AdjList: list, Cols: cols, Rows: rows, Directed: directed}
}

/*
* Agrega una arista entre un origen y un destino
* Se presupone que es un grafo no dirigido por lo que
* se agrega la arista en ambos sentidos
* */
func (g *GraphMesh) AddEdge(origin, dest int, weight float32) {
	g.AdjList[origin] = append(g.AdjList[origin], &Edge{v: origin, w: dest, weight: weight})

	if g.Directed {
		g.AdjList[dest] = append(g.AdjList[dest], &Edge{v: dest, w: origin, weight: weight})
	}
}

func (g *GraphMesh) AdjacentsOf(vertex int) []*Edge {
	if _, ok := g.AdjList[vertex]; ok {
		return g.AdjList[vertex]
	}
	return []*Edge{}
}

/*
* Se crea la malla con las formulas:
* - la formula (i * g.Cols) + j da el vertice actual
* - la formula (i * g.Cols) + j + 1 da el vertice siguiente al vertice actual
* - la formula (i * g.Cols) + j + g.Cols aplica las aristas verticales de la malla
* Ejemplo con un grafo 3 * 3:
* 0 - 1 - 2
* |	  |   |
* 3 - 4 - 5
* |   |   |
* 6 - 7 - 8
* */
func (g *GraphMesh) BuildMesh() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			currentVertex := (i * g.Cols) + j
			if j < g.Cols-1 {
				g.AddEdge(currentVertex, currentVertex+1, rand.Float32())
			}
			if i < g.Rows-1 {
				g.AddEdge(currentVertex, currentVertex+g.Cols, rand.Float32())
			}
		}
	}
}

func (g *GraphMesh) Print() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			currentVertex := (i * g.Cols) + j
			adVertex := g.AdjList[currentVertex]
			numAdj := len(adVertex)
			if i == 0 {
				if numAdj == 2 {
					if j == 0 {
						fmt.Print("┌")
					} else if j == g.Cols-1 {
						fmt.Print("┐")
					} else {
						fmt.Print("─")
					}
				} else if numAdj == 3 {
					fmt.Print("┬")
				} else if numAdj == 1 {
					if j == 0 {
						fmt.Print("┌")
					} else if j < g.Cols-1 {
						fmt.Print("─")
					} else {
						fmt.Print("┐")
					}
				}
			} else if i < g.Rows-1 {
				if numAdj == 3 {
					if j == 0 {
						fmt.Print("├")
					} else if j == g.Cols-1 {
						fmt.Print("┐")
					} else {
						fmt.Print("┬")
					}
				} else if numAdj == 4 {
					fmt.Print("┼")
				} else if numAdj == 2 {
					if j == 0 {
						fmt.Print("├")
					} else if j == g.Cols-1 {
						fmt.Print("┤")
					} else {
						fmt.Print("─")
					}
				} else {
					if j == 0 {
						fmt.Print("├")
					} else if j < g.Cols-1 {
						fmt.Print("─")
					} else {
						fmt.Print("┤")

					}
				}
			} else if i == g.Rows-1 {
				if numAdj == 2 {
					if j == 0 {
						fmt.Print("└")
					} else if j == g.Cols-1 {
						fmt.Print("┘")
					} else {
						fmt.Print("┴")
					}
				} else if numAdj == 3 {
					fmt.Print("┴")
				} else if numAdj == 1 {
					if j == 0 || j == g.Cols-1 {
						fmt.Print("┴")
					} else if j < g.Cols-1 {
						fmt.Print("┴")
					}
				}
			}
		}
		fmt.Println()
	}
}

func (g *GraphMesh) AddNewEdge(e *Edge) {
	g.AddEdge(e.v, e.w, e.weight)
}

/**
* Metodos principales del tipo Edge (arista)
* */
func (e *Edge) PrintEdge() {
	fmt.Printf("Desde %d hasta %d. Peso %f\n", e.v, e.w, e.weight)
}

func (g *GraphMesh) Size() int {
	return g.size
}

func (e *Edge) Weight() float32 {
	return e.weight
}

func (e *Edge) Vertexes() (int, int) {
	return e.v, e.w
}

func (e *Edge) Oposed(vertex int) int {
	if vertex == e.v {
		return e.w
	}

	return e.v
}

func (e *Edge) Origin() int {
	return e.v
}
