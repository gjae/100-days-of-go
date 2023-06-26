package busqueda

import (
	"container/list"
	"fmt"
	"gjae/matady/grafo"
)

type Bfs struct {
	Conteo   int
	marcados []bool
	AristaA  []int
	distA    []int
	graf     *grafo.Grafo
	origen   int
	destino  int
	colores  []int
	Queue    *list.List
	pila     *list.List
}

func New(g *grafo.Grafo, origen int, destino int) *Bfs {
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
		destino:  destino,
		Queue:    list.New().Init(),
		pila:     list.New().Init(),
	}

	bfs.Buscar(origen)

	return bfs
}

func (b *Bfs) Contar() {
	b.Conteo++
}

func (b *Bfs) Buscar(origen int) {
	b.marcados[origen] = true
	b.Contar()
	b.Queue.PushBack(origen)
	b.colores[origen] = 1

	for el := b.Queue.Front(); el != nil; el = el.Next() {
		v := el.Value.(int)
		for _, w := range b.graf.AdyacentesDe(v) {
			if !b.marcados[w] {
				b.marcados[w] = true
				b.AristaA[w] = v
				b.distA[w] = b.distA[v] + 1
				b.Queue.PushBack(w)
			}
		}
		b.colores[origen] = 2
		b.pila.PushFront(v)
	}

}

func (b *Bfs) EsConexo() bool {
	for _, v := range b.marcados {
		if !v {
			return false
		}
	}

	return true
}

func (b *Bfs) GetConteo() int {
	return b.Conteo
}

func (b *Bfs) ImprimirCaminino(v int) {
	pila := list.New().Init()
	for i := v; i != b.origen; i = b.AristaA[i] {
		pila.PushFront(i)
	}

	pila.PushFront(b.origen)
	for el := pila.Front(); el != nil; el = el.Next() {
		val := el.Value.(int)
		fmt.Printf("%d ", val)
	}
}

func (b *Bfs) GetCaminoCorto() int {

	fmt.Println(b.distA[b.destino])
	b.ImprimirCaminino(b.destino)
	return b.distA[b.destino]
}
