package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

type Options []string

func getRandomOption() string {
	options := Options{"Uno", "Dos", "tres"}

	return options[rand.Intn(len(options))]
}

func main() {
	var userOption string
	var nextPlay int = 1
	for nextPlay == 1 {
		fmt.Println("Ingrese una opción: ")
		fmt.Scanf("%s", &userOption)
		correctOption := getRandomOption()
		if userOption == correctOption {
			fmt.Println("Respuesta correcta ...")
		} else {
			fmt.Printf("Respuesta incorrecta ... %v \n", correctOption)
		}
		fmt.Println("Desea volver a jugar? [1: Si, 0: No]: ")
		fmt.Scanf("%d", &nextPlay)
	}
}	