package main

import (
	"testing"
	"reflect"
	"math"
	"fmt"
)

func TestCloneTriangle(t *testing.T) {
	t1 := NewTriangle(.5, .5, 0,
		-.5, -.5, 0,
		-.5, .5, 0)
	t2 := t1.Clone()
	if !reflect.DeepEqual(t1, t2) {
		t.FailNow()
	}

	t2.Apply(TransM(NewV4(1, 0, 0)))
	if reflect.DeepEqual(t1, t2) {
		t.FailNow()
	}
}

func TestModelClone(t *testing.T) {
	m1 := &Model{triangles: []Triangle{*NewTriangle(
		.5, .5, 0,
		-.5, -.5, 0,
		-.5, .5, 0)}}
	m2 := m1.Clone()
	if !reflect.DeepEqual(m1, m2) {
		t.FailNow()
	}

	m2.Apply(TransM(NewV4(1, 0, 0)))
	if reflect.DeepEqual(m1, m2) {
		t.FailNow()
	}
}

func TestModelMerge(t *testing.T) {
	m1 := Model{triangles: []Triangle{*NewTriangle(
		.5, .5, 0,
		-.5, -.5, 0,
		-.5, .5, 0)}}
	m2 := *m1.Clone().Apply(TransM(NewV4(0, 0, 1)))

	m3 := *m1.Merge(m2)

	if !reflect.DeepEqual(m3.triangles, append(m1.triangles, *NewTriangle(
		.5, .5, 1,
		-.5, -.5, 1,
		-.5, .5, 1))) {
		t.Fail()
	}
}

func assertAlmostEqual(t *testing.T, v1 float64, v2 float64) {
	if math.Abs(v1 - v2) > .0000001 {
		t.FailNow()
	}
}

func assertAlmostEqualV4(t *testing.T, v1 V4, v2 V4) {
	assertAlmostEqual(t, v2.x, v1.x)
	assertAlmostEqual(t, v2.y, v1.y)
	assertAlmostEqual(t, v2.z, v1.z)
	assertAlmostEqual(t, v2.w, v1.w)
}

func TestModelRotate(t *testing.T) {
	m := Model{triangles: []Triangle{*NewTriangle(
		0, 0, 0,
		1, 1, 0,
		1, 0, 0)}}
	m.Rot(rad(90), rad(90), rad(90))
	v1 := m.triangles[0].v1
	assertAlmostEqual(t,0, v1.x)
	assertAlmostEqual(t,0, v1.y)
	assertAlmostEqual(t,0, v1.z)

	v2 := m.triangles[0].v2
	assertAlmostEqual(t,0, v2.x)
	assertAlmostEqual(t,-1, v2.y)
	assertAlmostEqual(t,1, v2.z)

	v3 := m.triangles[0].v3
	assertAlmostEqual(t,0, v3.x)
	assertAlmostEqual(t,0, v3.y)
	assertAlmostEqual(t,1, v3.z)
}

func TestVectorCrossProduct(t *testing.T) {
	cross := Cross(*NewV4(1, 2, 0), *NewV4(4, 5, 6))
	assertAlmostEqual(t, 12, cross.x)
	assertAlmostEqual(t, -6, cross.y)
	assertAlmostEqual(t, -3, cross.z)
	assertAlmostEqual(t, 1, cross.w)
}

func TestVectorDotProduct(t *testing.T) {
	assertAlmostEqual(t, 6, Dot(*NewV4(1, 2, 0), *NewV4(-4, 5, 6)))
	fmt.Println(NewV4(-40, 5, 6).Length())
}

func TestVectorLength(t *testing.T) {
	assertAlmostEqual(t, 1.7320508075688772, NewV4(1, 1, 1).Length())
	assertAlmostEqual(t, 1.4142135623730951, NewV4(1, 1, 0).Length())
	assertAlmostEqual(t, 5, NewV4(3, 4, 0).Length())
}

func TestVectorAngle(t *testing.T) {
	assertAlmostEqual(t, 0, Angle(*NewV4(1, 1, 1), *NewV4(1, 1, 1)))
	assertAlmostEqual(t, math.Pi, Angle(*NewV4(1, 1, 1), *NewV4(-1, -1, -1)))
	assertAlmostEqual(t, math.Pi / 2, Angle(*NewV4(1, 1, 0), *NewV4(-1, 1, 0)))
	assertAlmostEqual(t, math.Pi / 2, Angle(*NewV4(2, 2, 0), *NewV4(-1, 1, 0)))
}

func TestNormalComputation(t *testing.T) {
	// counter-clockwise vertex winding:
	triangle := NewTriangle(.5, .5, 0,  -.5, .5, 0,  -.5, -.5, 0)
	assertAlmostEqualV4(t, *NewV4(0, 0, 1), triangle.Normal())

	triangle = NewTriangle(-.5, .5, 0,  -.5, -.5, 0, .5, .5, 0)
	assertAlmostEqualV4(t, *NewV4(0, 0, 1), triangle.Normal())

	// clockwise vertex winding:
	triangle = NewTriangle(-.5, .5, 0, .5, .5, 0,  -.5, -.5, 0).Apply(RotX(rad(45)))
	assertAlmostEqualV4(t, *NewV4(0, 0.7071067811865476, -0.7071067811865476), triangle.Normal())

	triangle = NewTriangle(.5, .5, 0,  -.5, -.5, 0,  -.5, .5, 0)
	assertAlmostEqualV4(t, *NewV4(0, 0, -1), triangle.Normal())
}
