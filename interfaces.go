package main

import "fmt"
import "math"

// START SHOW OMIT
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64  { return r.width * r.height }
func (r rect) perim() float64 { return 2*r.width + 2*r.height }

func (c circle) area() float64  { return math.Pi * c.radius * c.radius }
func (c circle) perim() float64 { return 2 * math.Pi * c.radius }

// END SHOW OMIT
// START RUN OMIT

func measure(g geometry) {
	fmt.Printf("g = %#v\n", g)
	fmt.Printf("g.area() = %f\n", g.area())
	fmt.Printf("g.perim() = %f\n", g.perim())
	fmt.Println()
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

// END RUN OMIT
