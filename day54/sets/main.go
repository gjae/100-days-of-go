package main

import "fmt"

type Set struct {
	IntegerMap map[int]bool
}

func (set *Set) New() {
	set.IntegerMap = make(map[int]bool)
}

// adds the element to the set

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.IntegerMap[element] = true
	}
}

// Check if element is contained in set
func (set *Set) ContainsElement(element int) bool {
	_, ok := set.IntegerMap[element]

	return ok
}

// Delete an element from set
func (set *Set) DeleteElement(element int) {
	delete(set.IntegerMap, element)
}

// Return a new set with intersect elements betwen set and anotherset
func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()
	var value int

	for value, _ = range set.IntegerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}

	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()

	for value, _ := range set.IntegerMap {
		unionSet.AddElement(value)
	}

	for value, _ := range anotherSet.IntegerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

func main() {
	var set *Set

	set = &Set{}
	anotherSet := &Set{}
	anotherSet.New()
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	anotherSet.AddElement(3)
	anotherSet.AddElement(4)
	anotherSet.AddElement(1)

	fmt.Println(set)
	fmt.Print(set.ContainsElement(1))

	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))
	fmt.Println(anotherSet.Union(set))
}
