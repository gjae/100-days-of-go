package busquedas

import (
	"gjae/matady/grafo"
)

type Busqueda struct {
	Conteo   int
	marcados []bool
	graf     *grafo.Grafo
	AristaA  []int
}

func New(g *grafo.Grafo) *Busqueda {
	marcados := make([]bool, g.GetTam())
	aristasA := make([]int, g.GetTam())

	return &Busqueda{
		marcados: marcados,
		graf:     g,
		AristaA:  aristasA,
	}
}

func (b *Busqueda) Contar() {
	b.Conteo++
}

func (b *Busqueda) Buscar(origen int) {
	b.marcados[origen] = true
	b.Contar()

	for _, ady := range b.graf.AdyacentesDe(origen) {
		if !b.marcados[ady] {
			b.Buscar(ady)
			b.AristaA[ady] = ady
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
