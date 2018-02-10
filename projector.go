package main

import "math"

type Projector struct {
	aov float64
	size int
	plane float64
	scale float64
}

func NewProjector(resolution int, angleOfView float64) *Projector {
	aov := rad(angleOfView)
	plane := math.Tan(aov / 2.0)	// half the width of projection plane
	p := Projector{
		aov: aov,
		size: resolution,
		plane: plane,
		scale: float64(resolution) / (plane * 2.0),
	}
	return &p
}

func (p *Projector) project(v V4) V2 {
	// Given a vertex, computes its (x, y) pixel projection on the
	// projection plane, normalizes to NDC space, converts to raster space and
	// returns the screen pixel coordinates.
	return V2{
		x: ((v.x / -v.z) + p.plane) * p.scale,
		y: ((v.y / -v.z) + p.plane) * p.scale,
	}
}