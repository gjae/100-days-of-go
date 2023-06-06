package main

import (
	"bufio"
	"fmt"
	Anchura "gjae/matady/busquedas/anchura"
	Profundidad "gjae/matady/busquedas/profundidad"
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

	fmt.Println("Buscando camino: ", caminoConVertices)
	fmt.Println("Existe el camino: ", grafo.ChequearCamino(caminoConVertices))
	fmt.Println("Es dirigido: ", grafo.EsDirigido())
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
	mat := grafo.New(numVertices)

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

	mat.PrintGrafo()
	fmt.Println("---------")
	BuscarCaminoAleatorio(vertices, mat)
	bp := Profundidad.New(mat)
	bp.Buscar(1)

	fmt.Println("(Busqueda en profundidad) : Conteo para 1: ", bp.GetConteo())
	fmt.Println("(Busqueda en profundidad) : Es conexo: ", bp.EsConexo())
	fmt.Print("------------------\n")
	ba := Anchura.New(mat)
	ba.Buscar(1)
	fmt.Println("(Busqueda en Anchura) : Es conexo: ", ba.EsConexo())

}
