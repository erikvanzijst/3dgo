// A very crude ASCII STL parser.
// https://en.wikipedia.org/wiki/STL_(file_format)
//
// Limitations: can only parse triangle facets (vertex triplets).
package main

import (
	"io"
	"bufio"
	"fmt"
	"regexp"
	"container/list"
	"math"
)

var pat = regexp.MustCompile("\\s*vertex\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s*")

type STLReader struct {
	scanner *bufio.Scanner
}

func NewSTLReader(reader io.Reader) *STLReader {
	return &STLReader{
		scanner: bufio.NewScanner(reader),
	}
}

func (r *STLReader) readVertex() *V4 {
	for {
		if !r.scanner.Scan() {
			return nil
		}
		line := r.scanner.Text()
		vertices := pat.FindStringSubmatch(line)
		if len(vertices) != 4 {
			continue
		}
		var x, y, z float64
		fmt.Sscanf(vertices[1], "%f", &x)
		fmt.Sscanf(vertices[2], "%f", &y)
		fmt.Sscanf(vertices[3], "%f", &z)
		return NewV4(x, y, z)
	}
}

// ReadTriangle returns the next triangle (facet) from the stream.
// When the end of the file is reached, nil is returned.
func (r *STLReader) ReadTriangle() *Triangle {
	v1 := r.readVertex()
	v2 := r.readVertex()
	v3 := r.readVertex()
	if v1 ==  nil {
		return nil
	} else if v2 == nil || v3 == nil {
		panic("Invalid number of vertices in STL file.")
	}
	return &Triangle{v1: *v1, v2: *v2, v3: *v3}
}

// ReadModel returns the model as defined in the loaded STL file.
// Models can be scaled to fit in the ((-1, -1, -1), ..., (1, 1, 1))
// bounding box using the `scale` parameter. Use `scale=false` to keep
// the STL file's original vertex values.
func (r *STLReader) ReadModel(scale bool) *Model {
	var minV = func(v1 *V4, v2 *V4) *V4 {
		return NewV4(math.Min(v1.x, v2.x), math.Min(v1.y, v2.y), math.Min(v1.z, v2.z))
	}
	var maxV = func(v1 *V4, v2 *V4) *V4 {
		return NewV4(math.Max(v1.x, v2.x), math.Max(v1.y, v2.y), math.Max(v1.z, v2.z))
	}

	elements := list.New()
	for t := r.ReadTriangle(); t != nil; t = r.ReadTriangle() {
		elements.PushBack(t)
	}
	var min, max *V4
	model := Model{make([]Triangle, elements.Len())}
	for i, e := 0, elements.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Triangle)
		if i == 0 {
			min, max = &t.v1, &t.v1
		}
		model.triangles[i] = *t

		for _, v := range []V4{t.v1, t.v2, t.v3} {
			min = minV(min, &v)
			max = maxV(max, &v)
		}
		i++
	}

	if scale && elements.Len() > 0 {
		// translate model to the origin:
		model.Move(
			((max.x - min.x) / 2.) - max.x,
			((max.y - min.y) / 2.) - max.y,
			((max.z - min.z) / 2.) - max.z)

		// and scale it to fill the ((-.5, -.5, -.5), (.5, .5, .5)) bounding box:
		factor := 1. / math.Max(max.x - min.x, math.Max(max.y - min.y, max.z - min.z))
		model.Apply(ScaleM(factor, factor, factor))
	}
	return &model
}
