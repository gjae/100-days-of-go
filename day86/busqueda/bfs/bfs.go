package busqueda

import (
	"container/list"
	"fmt"
	"gjae/grafos/grafo"
)

type Bfs struct {
	Conteo   int
	marcados []bool
	AristaA  []int
	distA    []int
	graf     *grafo.Grafo
	origen   int
	colores  []int
	Queue    *list.List
}

func New(g *grafo.Grafo, origen int) *Bfs {
	marcados := make([]bool, g.GetTam())
	aristasA := make([]int, g.GetTam())

	colores := make([]int, g.GetTam())

	bfs := &Bfs{
		marcados: marcados,
		graf:     g,
		AristaA:  aristasA,
		origen:   origen,
		colores:  colores,
		distA:    make([]int, g.GetTam()),
		Queue:    list.New().Init(),
	}

	bfs.Buscar(origen, 1)
	for v, c := range bfs.colores {
		fmt.Printf("%d: C%d\n", v, c)
	}

	return bfs
}

func (b *Bfs) Contar() {
	b.Conteo++
}

func (b *Bfs) Buscar(origen, color int) {
	b.marcados[origen] = true
	b.Contar()
	b.Queue.PushBack(origen)
	auxColor := color
	dictColores := make(map[int]bool)

	b.colores[origen] = color

	for el := b.Queue.Front(); el != nil; el = el.Next() {
		v := el.Value.(int)

		// Se le asigna a lavariable "color" el "color" original
		color = auxColor
		for _, w := range b.graf.AdyacentesDe(v) {
			// Si se encuentra un vertice con el color "color" entonces
			// se ubica el siguiente color mas bajo
			if b.colores[w] == color {
				color++
			}

			if !b.marcados[w] {
				b.marcados[w] = true
				b.AristaA[w] = v
				b.distA[w] = b.distA[v] + 1
				b.Queue.PushBack(w)
			}
		}
		b.colores[v] = color
		dictColores[color] = true
	}

	// Muestra en pantalla la cantidad de colores usados
	fmt.Println(len(dictColores))
}

func (b *Bfs) EsConexo() bool {
	for _, v := range b.marcados {
		if !v {
			return false
		}
	}

	return true
}
