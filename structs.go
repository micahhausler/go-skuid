package main

import "fmt"

// START OMIT
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{Name: "Alice", Age: 30})
	fmt.Println(Person{Name: "Fred"})

	fmt.Println(&Person{Name: "Ann", Age: 40}) // An `&` prefix yields a pointer to the struct.
	s := Person{Name: "Sean", Age: 50}         // Access struct fields with a dot.
	fmt.Println(s.Name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.Age)

	sp.Age = 51 // Structs are mutable.
	fmt.Println(sp.Age)
}

// END OMIT
