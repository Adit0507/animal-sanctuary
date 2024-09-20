package main

import (
	"fmt"
	"math/rand"
	"time"
)

type elephant struct {
	name string
}

func (el *elephant) String() string {
	return el.name
}

func (el *elephant) move() string {
	switch rand.Intn(2) {
	case 0:
		return "rolls around"
	default:
		return "tramples the ground"
	}
}

func (el *elephant) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "bamboo"
	default:
		return "bananas"
	}
}

type gopher struct {
	name string
}

func (g *gopher) String() string {
	return g.name
}

func (g *gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "Scurries along the ground"
	default:
		return "burrows in the sand"
	}
}

func (g *gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "carrot"
	case 1:
		return "lettuce"
	case 2:
		return "radish"
	case 3:
		return "corn"
	default:
		return "root"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())
	animals := []animal{
		&elephant{name: " Dumbo"},
		&gopher{name: " Go gopher"},
	}

	var sol, hour int
	for {
		fmt.Printf("%2d:00", hour)
		if hour < sunrise || hour > sunset {
			fmt.Println(" Animals are sleeping")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)
		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
