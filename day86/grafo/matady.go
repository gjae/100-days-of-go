package grafo

import "fmt"

type Grafo struct {
	mat      [][]int
	tam      int
	dirigido bool
}

func New(tam int, dirigido bool) *Grafo {

	mat := make([][]int, tam)

	for i := 0; i < tam; i++ {
		mat[i] = make([]int, tam)
	}

	return &Grafo{mat: mat, tam: tam, dirigido: dirigido}
}

func (g *Grafo) AgregarArista(f, c int) {
	g.mat[f][c] = 1

	if !g.dirigido {
		g.mat[c][f] = 1
	}

}

func (g *Grafo) ExisteArco(f, c int) bool {
	return g.mat[f][c] == 1
}

func (g *Grafo) VectorExgrado() {
	for i, r := range g.mat {
		var grado int
		for j, _ := range r {
			if g.mat[i][j] == 1 {
				grado++
			}
		}

		fmt.Printf("Vertice %d, Exgrados: %d\n", i+1, grado)
	}
}

func (g *Grafo) VectorInGrado() {
	for i, r := range g.mat {
		var grado int
		for j, _ := range r {
			if g.mat[j][i] == 1 {
				grado++
			}
		}
		fmt.Printf("Vertice: %d, Ingrado: %d\n", i+1, grado)
	}
}

func (g *Grafo) ChequearCamino(camino []int) bool {

	for i := 0; i < len(camino)-1; i++ {
		v1 := camino[i]
		v2 := camino[i+1]
		if !g.ExisteArco(v1, v2) {
			return false
		}
	}

	return true
}

func (g *Grafo) EsDirigido() bool {
	return g.dirigido
}

func (g *Grafo) AdyacentesDe(v int) []int {
	var adyacentes []int
	vertice := g.mat[v]

	for key, v := range vertice {
		if v == 1 {
			adyacentes = append(adyacentes, key)
		}
	}

	return adyacentes
}

func (g *Grafo) GetTam() int {
	return g.tam
}

func (g *Grafo) PrintGrafo() {
	for i, _ := range g.mat {
		if i == 0 {
			fmt.Print("  | ")
		}
		if i < g.tam {
			fmt.Printf(" %d ", i)
		}
	}

	fmt.Println()
	for i, r := range g.mat {
		fmt.Printf("%d | ", i)
		for _, v := range r {
			fmt.Printf(" %d ", v)
		}
		fmt.Println()
	}
}

func (g *Grafo) VerticeConMayorGrado() int {
	var vertMayGrado int
	var numAdy int
	for v, _ := range g.mat {
		adys := len(g.AdyacentesDe(v))
		if len(g.AdyacentesDe(v)) > numAdy {
			vertMayGrado = v
			numAdy = adys
		}
	}

	return vertMayGrado
}
