package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Cube struct {
	x float64
}

type Sphere struct {
	r float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * c.r * c.r * c.r
}


type shapes []Shape3D

func (a shapes) Len() int {
	return len(a)
}

func (a shapes) Less(i int, j int) bool {
	return a[i].Vol() < a[j].Vol()
}

func (a shapes) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func PrintShapes(a shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube volumen: %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Printf("Unknow data type!")
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}

		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}

	PrintShapes(data)
	// Sorting
	sort.Sort(shapes(data))
	//  Reverse sorting

	sort.Sort(sort.Reverse(shapes(data)))
	PrintShapes(data)
}