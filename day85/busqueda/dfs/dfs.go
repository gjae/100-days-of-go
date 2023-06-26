package busqueda

import (
	"container/list"
	"fmt"
	"gjae/matady/grafo"
)

type Dfs struct {
	Conteo   int
	marcados []bool
	AristaA  []int
	distA    []int
	graf     *grafo.Grafo
	origen   int
	destino  int
	colores  []int
	pila     *list.List
}

func New(g *grafo.Grafo, origen int) *Dfs {
	marcados := make([]bool, g.GetTam())
	aristasA := make([]int, g.GetTam())

	colores := make([]int, g.GetTam())

	dfs := &Dfs{
		marcados: marcados,
		graf:     g,
		AristaA:  aristasA,
		origen:   origen,
		colores:  colores,
		distA:    make([]int, g.GetTam()),
		pila:     list.New().Init(),
	}

	dfs.Buscar(origen)
	return dfs
}

func (b *Dfs) Contar() {
	b.Conteo++
}

func (b *Dfs) Buscar(origen int) {
	b.marcados[origen] = true
	b.colores[origen] = 1
	b.Contar()

	for _, ady := range b.graf.AdyacentesDe(origen) {
		if !b.marcados[ady] {
			b.AristaA[ady] = origen
			b.marcados[ady] = true
			b.distA[ady] = b.distA[origen] + 1
			b.pila.PushFront(origen)
			b.Buscar(ady)
		} else {
			if b.colores[ady] == 1 {
				fmt.Println("Hay ciclos en el vertice ", ady, " Desde ", origen)
			}
		}
	}

	b.colores[origen] = 2
}

func (b *Dfs) EsConexo() bool {
	for _, v := range b.marcados {
		if !v {
			return false
		}
	}

	return true
}

func (b *Dfs) GetConteo() int {
	return b.Conteo
}

func (d *Dfs) ImprimirCaminino(v int) {

	pila := list.New().Init()
	for i := v; i != d.origen; i = d.AristaA[i] {
		pila.PushFront(i)
	}

	pila.PushFront(d.origen)
	for el := pila.Front(); el != nil; el = el.Next() {
		val := el.Value.(int)
		fmt.Printf("%d ", val)
	}

}

func (b *Dfs) GetCaminoLargo() (int, int, int) {
	maxV := 0
	maxK := 0
	for k, v := range b.distA {
		if v > maxV && len(b.graf.AdyacentesDe(maxK)) > 0 {
			maxK = k
			maxV = v
			b.destino = k
		}
	}

	fmt.Println(b.origen, maxK)
	fmt.Println(maxV)
	b.ImprimirCaminino(b.destino)
	return b.distA[b.destino], b.origen, b.destino
}
