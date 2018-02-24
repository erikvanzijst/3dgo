// Copyright 2018 Erik van Zijst -- erik.van.zijst@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"testing"
	"reflect"
	"math"
	"github.com/stretchr/testify/assert"
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
    assert.InDelta(t, v1.x, v2.x, 1e-6, "V4.x")
    assert.InDelta(t, v1.y, v2.y, 1e-6, "V4.y")
    assert.InDelta(t, v1.z, v2.z, 1e-6, "V4.z")
    assert.InDelta(t, v1.w, v2.w, 1e-6, "V4.w")
}

func assertAlmostEqualM4(t *testing.T, m1 *M4, m2 *M4, e float64) {
	assert.InDelta(t, m1.a0, m2.a0, e, "a0")
	assert.InDelta(t, m1.a1, m2.a1, e, "a1")
	assert.InDelta(t, m1.a2, m2.a2, e, "a2")
	assert.InDelta(t, m1.a3, m2.a3, e, "a3")

	assert.InDelta(t, m1.b0, m2.b0, e, "b0")
	assert.InDelta(t, m1.b1, m2.b1, e, "b1")
	assert.InDelta(t, m1.b2, m2.b2, e, "b2")
	assert.InDelta(t, m1.b3, m2.b3, e, "b3")

	assert.InDelta(t, m1.c0, m2.c0, e, "c0")
	assert.InDelta(t, m1.c1, m2.c1, e, "c1")
	assert.InDelta(t, m1.c2, m2.c2, e, "c2")
	assert.InDelta(t, m1.c3, m2.c3, e, "c3")

	assert.InDelta(t, m1.d0, m2.d0, e, "d0")
	assert.InDelta(t, m1.d1, m2.d1, e, "d1")
	assert.InDelta(t, m1.d2, m2.d2, e, "d2")
	assert.InDelta(t, m1.d3, m2.d3, e, "d3")
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
}

func TestVectorLength(t *testing.T) {
    assertAlmostEqual(t, 1.7320508075688772, NewV4(1, 1, 1).Length())
    assertAlmostEqual(t, 1.4142135623730951, NewV4(1, 1, 0).Length())
    assertAlmostEqual(t, 5, NewV4(3, 4, 0).Length())
}

func TestV4_Normalize(t *testing.T) {
    assert.InDelta(t, 1., NewV4(1, 1, 1).Normalize().Length(), 1e-6, "Normalized length")
    assert.InDelta(t, 0., NewV4(0, 0, 0).Normalize().Length(), 1e-6, "Normalized length")
    assert.InDelta(t, NewV4(1, 1, 1).Normalize().Length(), NewV4(3, 3, 3).Normalize().Length(), 1e-6, "Normalized length")
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

func TestM4_Determinant(t *testing.T) {
	assert.Equal(t, 1., new(M4).SetIdentity().Determinant())
	assert.InEpsilonf(t, -6.199999999999999, (&M4{
		4, 0, 1, 2,
		-1, 1, 0, 3,
		-1, 0, 2, 0,
		3.3, 0, 1, 1,
	}).Determinant(), 1e-6, "Determinant failed")

	m := RotX(rad(-23)).Mul(RotY(rad(2))).Mul(TransM(NewV4(1, 2, 3))).Mul(ScaleM(0.3, 1.9, .9))
	assert.InDelta(t, 0.5129999999999999, m.Determinant(), 1e-6, "Determinant failed")
}

func TestM4_Inverse(t *testing.T) {
	m := new(M4).SetIdentity()
	assertAlmostEqualM4(t, m, m.Inverse(), 1e-6)

	m = RotX(rad(-23)).Mul(RotY(rad(2))).Mul(TransM(NewV4(1, 2, 3))).Mul(ScaleM(0.3, 1.9, .9))
	assertAlmostEqualM4(t, &M4{
		a0: 3.3313027567303193, a1:-0.04545439910091964, a2:-0.1070838536589986, a3:-3.3333333333333335,
		b0:8.453818109048769e-19, b1:0.4844762386591792, b2:-0.2056479623627757, b3:-1.0526315789473686,
		c0:0.0387772185583344, c1:0.4338812284922221, c2:1.0221601186299176, c3:-3.333333333333333,
		d0:0, d1:0, d2:0, d3:1,},
		m.Inverse(), 1e-6)
	assertAlmostEqualM4(t, new(M4).SetIdentity(), m.Inverse().Mul(m), 1e-6)
}

func TestRot(t *testing.T) {
    xaxis := NewV4(1, 0, 0)
    yaxis := NewV4(0, 1, 0)
    zaxis := NewV4(0, 0, 1)
    origin := NewV4(0, 0, 0)

    assertAlmostEqualV4(t, *origin, *origin.MultiplyM(Rot(origin, xaxis, rad(45))))
    assertAlmostEqualV4(t, *origin, *origin.MultiplyM(Rot(origin, yaxis, rad(45))))
    assertAlmostEqualV4(t, *origin, *origin.MultiplyM(Rot(origin, zaxis, rad(45))))

    p := NewV4(1, 1, 1)
    assertAlmostEqualV4(t, *NewV4(1, -1, 1), *p.MultiplyM(Rot(origin, xaxis, rad(90))))
    p = NewV4(1, 1, 1)
    assertAlmostEqualV4(t, *NewV4(1, 1, -1), *p.MultiplyM(Rot(origin, yaxis, rad(90))))
    p = NewV4(1, 1, 1)
    assertAlmostEqualV4(t, *NewV4(-1, 1, 1), *p.MultiplyM(Rot(origin, zaxis, rad(90))))

    // assert that a negative angle is the same as flipping the direction vector:
    p = NewV4(1, 1, 1)
    assertAlmostEqualV4(t, *NewV4(1, 1, -1), *p.MultiplyM(Rot(origin, xaxis, rad(-90))))
    p = NewV4(1, 1, 1)
    assertAlmostEqualV4(t, *NewV4(1, 1, -1), *p.MultiplyM(Rot(origin, NewV4(-1, 0, 0), rad(90))))

    // articulation > 1
    assertAlmostEqualV4(t,
        *NewV4(-1, -2, 3),
        *NewV4(-1, 3, 2).MultiplyM(Rot(origin, xaxis, rad(90))))

    // rotation around oneself
    assertAlmostEqualV4(t,
        *NewV4(-1, 3, 2),
        *NewV4(-1, 3, 2).MultiplyM(
            Rot(NewV4(-1, 3, 2), NewV4(-1, -1, 1), rad(90))))

	// rotation around custom vector from the origin
    assertAlmostEqualV4(t,
        *NewV4(0.2928932188134524, 1.7071067811865475, 0),
        *NewV4(1, 1, 1).MultiplyM(
            Rot(origin, &V4{-1, -1, 0, 1}, rad(90))))

    // rotation around custom vector from custom point
    assertAlmostEqualV4(t,
        *NewV4(0.4961430174654501, 2.1688893982361113, -3.0473954553288047),
        *NewV4(2, 1, -3).MultiplyM(
            Rot(NewV4(.5, .7, -1.1), &V4{-1, -.8, 12, 1}, rad(90))))

}
