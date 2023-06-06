package busquedas

import (
	"container/list"
	"gjae/matady/grafo"
)

type Busqueda struct {
	Conteo   int
	marcados []bool
	graf     *grafo.Grafo
	AristaA  []int
	Queue    *list.List
}

func New(g *grafo.Grafo) *Busqueda {
	marcados := make([]bool, g.GetTam())
	aristasA := make([]int, g.GetTam())

	return &Busqueda{
		marcados: marcados,
		graf:     g,
		AristaA:  aristasA,
		Queue:    list.New().Init(),
	}
}

func (b *Busqueda) Contar() {
	b.Conteo++
}

func (b *Busqueda) Buscar(origen int) {
	b.marcados[origen] = true
	b.Contar()
	b.Queue.PushBack(origen)

	for el := b.Queue.Front(); el != nil; el = el.Next() {
		v := el.Value.(int)
		for _, w := range b.graf.AdyacentesDe(v) {
			if !b.marcados[w] {
				b.marcados[w] = true
				b.AristaA[w] = v
				b.Queue.PushBack(w)
			}
		}
	}
}

func (b *Busqueda) EsConexo() bool {
	for _, v := range b.marcados {
		if !v {
			return false
		}
	}

	return true
}

func (b *Busqueda) GetConteo() int {
	return b.Conteo
}
