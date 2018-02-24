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
	"math"
)

// V4 represents a 4-dimensional (or homogeneous 3-dimensional) vector.
type V4 struct {
	x, y, z, w float64
}

func NewV4(x float64, y float64, z float64) *V4 {
	return &V4{x: x, y: y, z: z, w: 1}
}

// Length returns the length or magnitude of this vector.
func (v *V4) Length() float64 {
	return math.Sqrt(Dot(*v, *v))
}

// Normalize normalizes this vector and returns itself.
func (v *V4) Normalize() *V4 {
    l := v.Length()
    if l > 0 {
        v.x /= l
        v.y /= l
        v.z /= l
    }
    return v
}

// Add adds the specified vector to this vector and returns a new vector.
func (v *V4) Add(v2 *V4) V4 {
	return V4{
		x: v.x + v2.x,
		y: v.y + v2.y,
		z: v.z + v2.z,
		w: v.w,
	}
}

// Subtract subtracts the specified vector from this vector (v - v2) and returns a new vector.
func (v *V4) Subtract(v2 V4) V4 {
	return V4{
		x: v.x - v2.x,
		y: v.y - v2.y,
		z: v.z - v2.z,
		w: v.w,
	}
}

// MultiplyV multiplies this vector by the specified vector and returns itself.
func (v *V4) MultiplyV(v2 *V4) *V4 {
	v.x, v.y, v.z = v.x * v2.x, v.y * v2.y, v.z * v2.z
	return v
}

func (v *V4) MultiplyM(m *M4) *V4 {
	// Multiplies this vector with the specified matrix and returns itself.
	v.x, v.y, v.z, v.w =
		v.x * m.a0 + v.y * m.a1 + v.z * m.a2 + v.w * m.a3,
		v.x * m.b0 + v.y * m.b1 + v.z * m.b2 + v.w * m.b3,
		v.x * m.c0 + v.y * m.c1 + v.z * m.c2 + v.w * m.c3,
		v.x * m.d0 + v.y * m.d1 + v.z * m.d2 + v.w * m.d3
	return v
}

func Cross(v1 V4, v2 V4) V4 {
	// Computes the cross product of the specified matrices.
	return V4{
		x: v1.y * v2.z - v1.z * v2.y,
		y: v1.z * v2.x - v1.x * v2.z,
		z: v1.x * v2.y - v1.y * v2.x,
		w: 1,
	}
}

func Dot(v1 V4, v2 V4) float64 {
	return v1.x * v2.x + v1.y * v2.y + v1.z * v2.z
}

// Angle returns the angle in radians between the given vectors.
func Angle(v1 V4, v2 V4) float64 {
	v := Dot(v1, v2) / (v1.Length() * v2.Length())

	// Squash floating point rounding errors on parallel vectors
	if v > 1. {
		v = 1.
	} else if v < -1. {
		v = -1.
	}
	return math.Acos(v)
}

// V2 represents a 2-dimensional vector.
type V2 struct {
	x, y float64
}

// M4 represents a 4-dimensional matrix.
type M4 struct {
	a0, a1, a2, a3 float64
	b0, b1, b2, b3 float64
	c0, c1, c2, c3 float64
	d0, d1, d2, d3 float64
}

// SetIdentity turns this matrix into the identity matrix.
func (m *M4) SetIdentity() *M4 {
	m.a0, m.a1, m.a2, m.a3 = 1, 0, 0, 0
	m.b0, m.b1, m.b2, m.b3 = 0, 1, 0, 0
	m.c0, m.c1, m.c2, m.c3 = 0, 0, 1, 0
	m.d0, m.d1, m.d2, m.d3 = 0, 0, 0, 1
	return m
}

func (m *M4) Mul(m2 *M4) *M4 {
	// Multiplies this matrix with the specified matrix.
	a0 := m.a0 * m2.a0 + m.a1 * m2.b0 + m.a2 * m2.c0 + m.a3 * m2.d0
	a1 := m.a0 * m2.a1 + m.a1 * m2.b1 + m.a2 * m2.c1 + m.a3 * m2.d1
	a2 := m.a0 * m2.a2 + m.a1 * m2.b2 + m.a2 * m2.c2 + m.a3 * m2.d2
	a3 := m.a0 * m2.a3 + m.a1 * m2.b3 + m.a2 * m2.c3 + m.a3 * m2.d3

	b0 := m.b0 * m2.a0 + m.b1 * m2.b0 + m.b2 * m2.c0 + m.b3 * m2.d0
	b1 := m.b0 * m2.a1 + m.b1 * m2.b1 + m.b2 * m2.c1 + m.b3 * m2.d1
	b2 := m.b0 * m2.a2 + m.b1 * m2.b2 + m.b2 * m2.c2 + m.b3 * m2.d2
	b3 := m.b0 * m2.a3 + m.b1 * m2.b3 + m.b2 * m2.c3 + m.b3 * m2.d3

	c0 := m.c0 * m2.a0 + m.c1 * m2.b0 + m.c2 * m2.c0 + m.c3 * m2.d0
	c1 := m.c0 * m2.a1 + m.c1 * m2.b1 + m.c2 * m2.c1 + m.c3 * m2.d1
	c2 := m.c0 * m2.a2 + m.c1 * m2.b2 + m.c2 * m2.c2 + m.c3 * m2.d2
	c3 := m.c0 * m2.a3 + m.c1 * m2.b3 + m.c2 * m2.c3 + m.c3 * m2.d3

	d0 := m.d0 * m2.a0 + m.d1 * m2.b0 + m.d2 * m2.c0 + m.d3 * m2.d0
	d1 := m.d0 * m2.a1 + m.d1 * m2.b1 + m.d2 * m2.c1 + m.d3 * m2.d1
	d2 := m.d0 * m2.a2 + m.d1 * m2.b2 + m.d2 * m2.c2 + m.d3 * m2.d2
	d3 := m.d0 * m2.a3 + m.d1 * m2.b3 + m.d2 * m2.c3 + m.d3 * m2.d3

	m.a0, m.a1, m.a2, m.a3 = a0, a1, a2, a3
	m.b0, m.b1, m.b2, m.b3 = b0, b1, b2, b3
	m.c0, m.c1, m.c2, m.c3 = c0, c1, c2, c3
	m.d0, m.d1, m.d2, m.d3 = d0, d1, d2, d3

	return m
}

// Transpose returns a new matrix containing the transposition of this matrix.
func (m *M4) Transpose() *M4 {
	return &M4{
		m.a0, m.b0, m.c0, m.d0,
		m.a1, m.b1, m.c1, m.d1,
		m.a2, m.b2, m.c2, m.d2,
		m.a3, m.b3, m.c3, m.d3,
	}
}

func (m *M4) Determinant() float64 {
	// http://cg.info.hiroshima-cu.ac.jp/~miyazaki/knowledge/teche23.html
    return m.a0*m.b1*m.c2*m.d3 + m.a0*m.b2*m.c3*m.d1 + m.a0*m.b3*m.c1*m.d2 +
           m.a1*m.b0*m.c3*m.d2 + m.a1*m.b2*m.c0*m.d3 + m.a1*m.b3*m.c2*m.d0 +
           m.a2*m.b0*m.c1*m.d3 + m.a2*m.b1*m.c3*m.d0 + m.a2*m.b3*m.c0*m.d1 +
           m.a3*m.b0*m.c2*m.d1 + m.a3*m.b1*m.c0*m.d2 + m.a3*m.b2*m.c1*m.d0 -
           m.a0*m.b1*m.c3*m.d2 - m.a0*m.b2*m.c1*m.d3 - m.a0*m.b3*m.c2*m.d1 -
           m.a1*m.b0*m.c2*m.d3 - m.a1*m.b2*m.c3*m.d0 - m.a1*m.b3*m.c0*m.d2 -
           m.a2*m.b0*m.c3*m.d1 - m.a2*m.b1*m.c0*m.d3 - m.a2*m.b3*m.c1*m.d0 -
           m.a3*m.b0*m.c1*m.d2 - m.a3*m.b1*m.c2*m.d0 - m.a3*m.b2*m.c0*m.d1
}

// Inverse returns a new matrix containing the inverse of this matrix.
func (m *M4) Inverse() *M4 {
    // http://cg.info.hiroshima-cu.ac.jp/~miyazaki/knowledge/teche23.html
    det := m.Determinant()

    return &M4{
        a0: (m.b1*m.c2*m.d3 + m.b2*m.c3*m.d1 + m.b3*m.c1*m.d2 - m.b1*m.c3*m.d2 - m.b2*m.c1*m.d3 - m.b3*m.c2*m.d1) / det,
        a1: (m.a1*m.c3*m.d2 + m.a2*m.c1*m.d3 + m.a3*m.c2*m.d1 - m.a1*m.c2*m.d3 - m.a2*m.c3*m.d1 - m.a3*m.c1*m.d2) / det,
        a2: (m.a1*m.b2*m.d3 + m.a2*m.b3*m.d1 + m.a3*m.b1*m.d2 - m.a1*m.b3*m.d2 - m.a2*m.b1*m.d3 - m.a3*m.b2*m.d1) / det,
        a3: (m.a1*m.b3*m.c2 + m.a2*m.b1*m.c3 + m.a3*m.b2*m.c1 - m.a1*m.b2*m.c3 - m.a2*m.b3*m.c1 - m.a3*m.b1*m.c2) / det,
        b0: (m.b0*m.c3*m.d2 + m.b2*m.c0*m.d3 + m.b3*m.c2*m.d0 - m.b0*m.c2*m.d3 - m.b2*m.c3*m.d0 - m.b3*m.c0*m.d2) / det,
        b1: (m.a0*m.c2*m.d3 + m.a2*m.c3*m.d0 + m.a3*m.c0*m.d2 - m.a0*m.c3*m.d2 - m.a2*m.c0*m.d3 - m.a3*m.c2*m.d0) / det,
        b2: (m.a0*m.b3*m.d2 + m.a2*m.b0*m.d3 + m.a3*m.b2*m.d0 - m.a0*m.b2*m.d3 - m.a2*m.b3*m.d0 - m.a3*m.b0*m.d2) / det,
        b3: (m.a0*m.b2*m.c3 + m.a2*m.b3*m.c0 + m.a3*m.b0*m.c2 - m.a0*m.b3*m.c2 - m.a2*m.b0*m.c3 - m.a3*m.b2*m.c0) / det,
        c0: (m.b0*m.c1*m.d3 + m.b1*m.c3*m.d0 + m.b3*m.c0*m.d1 - m.b0*m.c3*m.d1 - m.b1*m.c0*m.d3 - m.b3*m.c1*m.d0) / det,
        c1: (m.a0*m.c3*m.d1 + m.a1*m.c0*m.d3 + m.a3*m.c1*m.d0 - m.a0*m.c1*m.d3 - m.a1*m.c3*m.d0 - m.a3*m.c0*m.d1) / det,
        c2: (m.a0*m.b1*m.d3 + m.a1*m.b3*m.d0 + m.a3*m.b0*m.d1 - m.a0*m.b3*m.d1 - m.a1*m.b0*m.d3 - m.a3*m.b1*m.d0) / det,
        c3: (m.a0*m.b3*m.c1 + m.a1*m.b0*m.c3 + m.a3*m.b1*m.c0 - m.a0*m.b1*m.c3 - m.a1*m.b3*m.c0 - m.a3*m.b0*m.c1) / det,
        d0: (m.b0*m.c2*m.d1 + m.b1*m.c0*m.d2 + m.b2*m.c1*m.d0 - m.b0*m.c1*m.d2 - m.b1*m.c2*m.d0 - m.b2*m.c0*m.d1) / det,
        d1: (m.a0*m.c1*m.d2 + m.a1*m.c2*m.d0 + m.a2*m.c0*m.d1 - m.a0*m.c2*m.d1 - m.a1*m.c0*m.d2 - m.a2*m.c1*m.d0) / det,
        d2: (m.a0*m.b2*m.d1 + m.a1*m.b0*m.d2 + m.a2*m.b1*m.d0 - m.a0*m.b1*m.d2 - m.a1*m.b2*m.d0 - m.a2*m.b0*m.d1) / det,
        d3: (m.a0*m.b1*m.c2 + m.a1*m.b2*m.c0 + m.a2*m.b0*m.c1 - m.a0*m.b2*m.c1 - m.a1*m.b0*m.c2 - m.a2*m.b1*m.c0) / det,
    }
}

type Triangle struct {
	// A single triangle, consisting of 3 vertices.
	v1, v2, v3 V4
}

func NewTriangle(points ...float64) *Triangle {
	return &Triangle{
		*NewV4(points[0], points[1], points[2]),
		*NewV4(points[3], points[4], points[5]),
		*NewV4(points[6], points[7], points[8])}
}

func (t Triangle) Clone() *Triangle {
	return &t
}

func (t *Triangle) Apply(m *M4) *Triangle {
	// Applies the specified transformation matrix to this triangle and returns itself.
	t.v1.MultiplyM(m)
	t.v2.MultiplyM(m)
	t.v3.MultiplyM(m)
	return t
}

func (t *Triangle) Normal() V4 {
	// Computes the triangle's normal vector.
	// The direction of the normal follows the triangle's winding direction:
	// counter-clockwise winding produces positive (z) normal, clockwise winding
	// produces negative (z) normals.
	v1 := t.v2.Subtract(t.v1)
	v2 := t.v3.Subtract(t.v1)
	return Cross(v1, v2)
}

type Model struct {
	// A model contains zero or more triangles.
	triangles []Triangle
}

func (m Model) Clone() *Model {
	m2 := Model{make([]Triangle, len(m.triangles))}
	for i, v := range m.triangles {
		m2.triangles[i] = v
	}
	return &m2
}

// Merge creates a new model that consist of the combination of this model and the supplied one.
func (m *Model) Merge(models ...Model) *Model {
	polys := len(m.triangles)
	for _, mod := range models {
		polys += len(mod.triangles)
	}
	m3 := Model{make([]Triangle, polys)}

	i := 0
	for _, mod := range append([]Model{*m}, models...) {
		for _, v := range mod.triangles {
			m3.triangles[i] = v
			i++
		}
	}
	return &m3
}

// Apply applies the specified transformation matrix to this model and returns itself.
func (m *Model) Apply(mat *M4) *Model {
	for i := 0; i < len(m.triangles); i++ {
		m.triangles[i].Apply(mat)
	}
	return m
}

func (m *Model) Move(x float64, y float64, z float64) *Model {
	// Translates the model.
	return m.Apply(TransM(NewV4(x, y, z)))
}

func (m *Model) Rot(ax float64, ay float64, az float64) *Model {
	// Rotates the model along the specified angles (in radians).
	return m.Apply(RotX(ax).Mul(RotY(ay).Mul(RotZ(az))))
}

func ScaleM(x float64, y float64, z float64) *M4 {
	// Creates a new scaling matrix with the specified magnitudes.
	return &M4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}
}

func TransM(v *V4) *M4 {
	// Creates a new translation matrix along the specified vector.
	return &M4{
		1, 0, 0, v.x,
		0, 1, 0, v.y,
		0, 0, 1, v.z,
		0, 0, 0, 1}
}

func RotX(a float64) *M4 {
	// Creates a new rotation matrix along the x-axis, with the specified angle in radians.
	return &M4{
		1, 0, 0, 0,
		0, math.Cos(a), -math.Sin(a), 0,
		0, math.Sin(a), math.Cos(a), 0,
		0, 0, 0, 1}
}

func RotY(a float64) *M4 {
	// Creates a new rotation matrix along the y-axis, with the specified angle in radians.
	return &M4{
		math.Cos(a), 0, math.Sin(a), 0,
		0, 1, 0, 0,
		-math.Sin(a), 0, math.Cos(a), 0,
		0, 0, 0, 1}
}

// RotZ creates a new rotation matrix along the z-axis, with the specified angle in radians.
func RotZ(a float64) *M4 {
	return &M4{
		math.Cos(a), -math.Sin(a), 0, 0,
		math.Sin(a), math.Cos(a), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Rot returns a rotation matrix that when applied to a vector, rotates the
// vector about the line through point `p` with direction vector `v` by the
// specified angle in radians.
func Rot(p *V4, v *V4, phi float64) *M4 {
    // https://sites.google.com/site/glennmurray/Home/rotation-matrices-and-formulas

    // Normalize the direction vector:
    l := v.Length()
    if l <= 0 {
        panic("Cannot rotate around vector of length zero")
    }
    x := v.x / l
    y := v.y / l
    z := v.z / l

    // Precompute intermediate values:
    x2 := x * x
    y2 := y * y
    z2 := z * z
    sp := math.Sin(phi)
    cp := math.Cos(phi)
    omcp := 1. - cp

    return &M4{
        a0: x2 + (y2 + z2) * cp,
        a1: x * y * omcp - z * sp,
        a2: x * z * omcp + y * sp,
        a3: (p.x * (y2 + z2) - x * (p.y * y + p.z * z)) * omcp + (p.y * z - p.z * y) * sp,
        b0: x * y * omcp + z * sp,
        b1: y2 + (x2 + z2) * cp,
        b2: y * z * omcp - x * sp,
        b3: (p.y * (x2 + z2) - y * (p.x * x + p.z * z)) * omcp + (p.z * x - p.x * z) * sp,
        c0: x * z * omcp - y * sp,
        c1: y * z * omcp + x * sp,
        c2: z2 + (x2 + y2) * cp,
        c3: (p.z * (x2 + y2) - z * (p.x * x + p.y * y)) * omcp + (p.x * y - p.y * x) * sp,
        d0: 0,
        d1: 0,
        d2: 0,
        d3: 1,
    }
}
