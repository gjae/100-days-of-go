package main

import (
	"flag"
	"gjae/graph-mesh/algos"
	"gjae/graph-mesh/algos/shortways"
	"gjae/graph-mesh/graph"
	"log"
)

func main() {

	cols := flag.Int("col", 0, "Indicar la cantidad de columnas para la malla.")
	rows := flag.Int("row", 0, "Indica la cantidad de filas que tendra la malla")

	flag.Parse()

	if *cols == 0 || *rows == 0 {
		log.Fatal("Indique una cantidad de filas y columnas usando los parametros -row [filas] -col [columnas]")
	}
	g := graph.NewGraph(*cols, *rows, true)
	g.BuildMesh()

	wayFinder := shortways.NewShortWayFinder(algos.PrimMST(g))
	wayFinder.Print()
}
