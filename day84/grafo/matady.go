package grafo

import "fmt"

type Grafo struct {
	mat [][]int
	tam int
}

func New(tam int) *Grafo {
	mat := make([][]int, tam)

	for i := 0; i < tam; i++ {
		mat[i] = make([]int, tam)
	}

	return &Grafo{mat: mat, tam: tam}
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

func (g *Grafo) AgregarArista(f, c int) {
	g.mat[f][c] = 1
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
	for i := 0; i <= g.tam; i++ {
		for j := 0; j <= g.tam; j++ {
			// Si se encuentra discrepancia  entonces es dirigido
			if g.mat[i][j] != g.mat[j][i] {
				return true
			}
		}
	}

	return false
}
