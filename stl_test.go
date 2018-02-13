package main

import (
	"strings"
	"fmt"
)

func ExampleStlReader() {
	var stl = `solid Object01
	  facet normal -1.583127e-002 1.253177e-001 9.919903e-001
		outer loop
		  vertex -2.131976e+001 -1.033176e+001 3.937008e+001
		  vertex -2.131976e+001 -5.408154e-001 3.813319e+001
		  vertex -2.375467e+001 -8.484154e-001 3.813319e+001
		endloop
	  endfacet
	  facet normal -4.649920e-002 1.174435e-001 9.919903e-001
		outer loop
		  vertex -2.131976e+001 -1.033176e+001 3.937008e+001
		  vertex -2.375467e+001 -8.484154e-001 3.813319e+001
		  vertex -2.603659e+001 -1.751890e+000 3.813319e+001
		endloop
	  endfacet
	endsolid Object01
	`

	reader := NewSTLReader(strings.NewReader(stl))
	fmt.Println(reader.ReadTriangle())
	fmt.Println(reader.ReadTriangle())
	fmt.Println(reader.ReadTriangle())

	// Output:
	// &{{-21.31976 -10.33176 39.37008 1} {-21.31976 -0.5408154 38.13319 1} {-23.75467 -0.8484154 38.13319 1}}
	// &{{-21.31976 -10.33176 39.37008 1} {-23.75467 -0.8484154 38.13319 1} {-26.03659 -1.75189 38.13319 1}}
	// <nil>
}

func ExampleScalingStlReader() {
	var stl = `solid Object01
	  facet normal -1.583127e-002 1.253177e-001 9.919903e-001
		outer loop
		  vertex -2e+000 -2e+000 -2e+000
		  vertex 0e+000 0e-000 0e+000
		  vertex 0e+000 1e-000 0e+000
		endloop
	  endfacet
	endsolid Object01
	`

	reader := NewSTLReader(strings.NewReader(stl))
	fmt.Println(reader.ReadModel(true))

	// Output:
	// &{[{{-0.3333333333333333 -0.5 -0.3333333333333333 1} {0.3333333333333333 0.16666666666666666 0.3333333333333333 1} {0.3333333333333333 0.5 0.3333333333333333 1}}]}

}
