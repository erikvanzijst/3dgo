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