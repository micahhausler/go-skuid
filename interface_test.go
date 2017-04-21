package geo_test

import "math"
import "testing"

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

// START OMIT
func TestArea(t *testing.T) {
	geos := []struct {
		input geometry
		want  float64
	}{
		{circle{1}, math.Pi},
		{rect{2, 2}, 4},
	}
	for _, geo := range geos {
		if got := geo.input.area(); got != geo.want {
			t.Errorf("Wanted %b, Got %b", geo.want, got)
		}
	}
}

// END OMIT
