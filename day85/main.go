package main

import (
	"bufio"
	"fmt"
	Bfs "gjae/matady/busqueda/bfs"
	Dfs "gjae/matady/busqueda/dfs"
	"gjae/matady/grafo"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func LeerArchivo(ruta *string) *bufio.Scanner {
	var archivo *os.File
	var err error

	if ruta == nil {
		archivo, err = os.Open("grafo.in")
	} else {
		archivo, err = os.Open(*ruta)
	}

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(archivo)
	return scanner
}

func existe(elementos []int, v int) bool {
	for _, vi := range elementos {
		if v == vi {
			return true
		}
	}

	return false
}

func BuscarCaminoAleatorio(vertices []int, grafo *grafo.Grafo) {
	rand.Seed(time.Now().UnixNano())
	var caminoConVertices []int

	for len(caminoConVertices) <= len(vertices) {
		v := vertices[rand.Intn(len(vertices))]
		caminoConVertices = append(caminoConVertices, v)
	}
}

func main() {
	var ruta *string

	if len(os.Args) > 1 {
		ruta = &os.Args[1]
	}

	scanner := LeerArchivo(ruta)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	var vertices []int

	numVertices, _ := strconv.Atoi(scanner.Text())
	dirigido, _ := strconv.Atoi(scanner.Text())
	mat := grafo.New(numVertices, dirigido == 0)

	for scanner.Scan() {
		auxVert := strings.Split(scanner.Text(), " ")
		if len(auxVert) == 2 {
			v1, _ := strconv.Atoi(auxVert[0])
			v2, _ := strconv.Atoi(auxVert[1])
			mat.AgregarArista(v1, v2)
			if !existe(vertices, v1) {
				vertices = append(vertices, v1)
			}
			if !existe(vertices, v2) {
				vertices = append(vertices, v2)
			}
		}
	}

	rand.Seed(time.Now().Unix())

	origen := 0
	destino := 0
	for origen >= destino {
		origen = rand.Intn(numVertices)
		destino = rand.Intn(numVertices)

		// Esta validacion adicional para el caso de los grafos dirigidos
		// si no se hace y se toma un origen que no tiene vertices adyacentes
		// al intentar imprimir el camino se quedará colgado el programa
		if len(mat.AdyacentesDe(origen)) == 0 {
			origen = 0
			destino = 0
		}

	}

	mat.PrintGrafo()
	dfs := Dfs.New(mat, origen)

	// Buscar el camino mas largo desde un origen
	// retorna dicho origen y el destino al camino mas largo encontrado
	_, origen, destino = dfs.GetCaminoLargo()

	// busca el camino mas corto entre un origen y un destino
	// dado por la busqueda del camino mas largo encontrado

	fmt.Println()
	bfs := Bfs.New(mat, origen, destino)

	_ = bfs.GetCaminoCorto()
}
