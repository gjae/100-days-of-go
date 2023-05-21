/**
* Esta es una representación de un grafo dado V vertices ingresados por STDIN
* utilizando algoritmo de fuerza bruta
*
* Autor: Giovanny Avila
* Materia: Introducción a la Teoría de Grafos
 */

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Par struct {
	A int
	B int
}

type Grafos map[string]map[int][]int

func calcularPares(elementos []int) []Par {
	totalElementos := len(elementos)

	var p []Par

	for i := 0; i < totalElementos; i++ {
		for j := 0; j < totalElementos; j++ {
			p = append(p, Par{A: i, B: j})
		}
	}

	return p
}

func calcularElementos(totalVert int) []int {
	var elementos []int

	for i := 0; i < totalVert; i++ {
		elementos = append(elementos, i)
	}

	return elementos
}

func generarCombinaciones(pares []Par) [][]Par {
	var combinaciones [][]Par

	// Inicializar combinacion vacia
	combinacionActual := []Par{}

	generarCombinacionesRecursivas(pares, 0, combinacionActual, &combinaciones)

	return combinaciones
}

func generarCombinacionesRecursivas(pares []Par, index int, combinacionActual []Par, combinaciones *[][]Par) {

	// Se agrega la combinacion actual a las combinaciones
	*combinaciones = append(*combinaciones, combinacionActual)

	for i := index; i < len(pares); i++ {

		// Se crea nueva combinacion
		// Y se crea una copia de la combinacion actual
		nuevaCombinacion := make([]Par, len(combinacionActual))
		copy(nuevaCombinacion, combinacionActual)

		nuevaCombinacion = append(nuevaCombinacion, pares[i])

		generarCombinacionesRecursivas(pares, i+1, nuevaCombinacion, combinaciones)
	}

}

func generarGrafos(combinaciones [][]Par) Grafos {
	grafos := make(Grafos)

	for _, v := range combinaciones {
		temp := make(map[int][]int)
		var comb string
		for _, vers := range v {
			vertice1, vertice2 := vers.A, vers.B

			// la instruccion _, ok := temp[ALGO] comprueba que el mapa tenga la key "ALGO"
			// si el mapa no tiene dicha key entonces "ok" es false, lo que indica que vertice (REPRESENTADO POR vertice1 y vertice2)
			// por tanto se agrega un "slice" de enteros que representa cada vertice
			if _, ok := temp[vertice1]; !ok {
				var vAux []int
				temp[vertice1] = vAux
			}

			if _, ok := temp[vertice2]; !ok {
				var vAux []int
				temp[vertice2] = vAux
			}

			// Los vecinos se toman de manera invertida
			// el vecino del vertice1 es el vertice2 y viceversa
			temp[vertice1] = append(temp[vertice1], vertice2)
			temp[vertice2] = append(temp[vertice2], vertice1)
		}

		// Este paso es opcional, unicamente lo hago
		// con el motivo de, al imprimirlo , mostrar los pares
		// que representan cada grafo
		comb = fmt.Sprintf("%v", v)

		grafos[comb] = temp
	}

	return grafos
}

func main() {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("%v ", err)
	}

	/**
	* Primero se crea el conjunto de V elementos
	* donde V es el numero de vertices
	 */
	pares := calcularPares(calcularElementos(v))

	posiblesGrafos := int(math.Pow(2, float64(len(pares))))

	fmt.Printf("Total de grafos posibles: %d\n", posiblesGrafos)

	combinaciones := generarCombinaciones(pares)

	grafosGenerados := generarGrafos(combinaciones)

	for pares, grafos := range grafosGenerados {
		if rand.Intn(posiblesGrafos)%2 == 0 {
			fmt.Printf("Grafos generados para los pares %v\n", pares)
			for vertice, vecinos := range grafos {
				fmt.Printf("Vertice %d -> Vecinos: %v\n", vertice, vecinos)
			}
			break
		}
	}
}
