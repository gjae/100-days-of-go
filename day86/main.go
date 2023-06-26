package main

import (
	"bufio"
	"fmt"
	busqueda "gjae/grafos/busqueda/bfs"
	"gjae/grafos/grafo"
	"log"
	"os"
	"strconv"
	"strings"
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
	mat := grafo.New(numVertices, false)

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

	origen := mat.VerticeConMayorGrado()
	mat.PrintGrafo()
	fmt.Print("\nNúmero cromatico: ")
	_ = busqueda.New(mat, origen)
}
