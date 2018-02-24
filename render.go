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
	"github.com/andlabs/ui"
	"time"
	"math"
	"os"
)

type Renderer struct {
	a         *ui.Area
	dp        *ui.AreaDrawParams
	model     Model
    cameraMatrix M4
	projector Projector
	rotTime   float64 // seconds per rotation
}

func (r *Renderer) mainLoop() {
	for {
		r.a.QueueRedrawAll()
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
}

func (r *Renderer) drawModel(a *ui.Area, dp *ui.AreaDrawParams, model *Model) {

	path := ui.NewPath(ui.Winding)
	for _, t := range model.triangles {
		if Dot(t.v1, t.Normal()) < 0. {	// back-face culling
			point := r.projector.project(t.v1)
			path.NewFigure(point.x, point.y)

			point = r.projector.project(t.v2)
			path.LineTo(point.x, point.y)

			point = r.projector.project(t.v3)
			path.LineTo(point.x, point.y)
			path.CloseFigure()
		}
	}
	path.End()

	dp.Context.Stroke(path,
		&ui.Brush{A:1, Type:ui.Solid},
		&ui.StrokeParams{ui.FlatCap, ui.MiterJoin, 1, 2, nil, 1})
	path.Free()
}

func (r *Renderer) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	angle := (float64(time.Now().UnixNano() % (int64(r.rotTime * 1e9))) / 1e9) *
				((2 * math.Pi) / r.rotTime)
    mat := RotX(math.Pi/2.).Mul(RotY(rad(23.4))).Mul(RotZ(angle))
    model := r.model.Clone().Apply(mat)

    r.drawModel(a, dp, model.Apply(r.cameraMatrix.Inverse()))
}

func (r Renderer) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	return
}

func (r Renderer) MouseCrossed(a *ui.Area, left bool) {
	return
}

func (r Renderer) DragBroken(a *ui.Area) {
	return
}

func (r *Renderer) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
    step := .25

    if !ke.Up {
        tm := new(M4).SetIdentity()

        switch ke.Key {
        case int32('w'):
            tm = TransM(NewV4(0, 0, -step))
        case int32('s'):
            tm = TransM(NewV4(0, 0, step))
        case int32('a'):
            tm = TransM(NewV4(-step, 0, 0))
        case int32('d'):
            tm = TransM(NewV4(step, 0, 0))
        }

        switch ke.ExtKey {
        case ui.Left:
            tm = RotY(rad(step*4))
        case ui.Right:
            tm = RotY(rad(-step*4))
        }
        r.cameraMatrix.Mul(tm)
        return true
    }
	return
}

func rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Cube() (*Model) {
	top := &Model{[]Triangle{
		// counter-clockwise vertex winding
		*NewTriangle(.5, .5, .5,  -.5, .5, .5,  -.5, -.5, .5),
		*NewTriangle(-.5, -.5, .5,  .5, -.5, .5,  .5, .5, .5),
	}}

	cube := top.Merge(
		*top.Clone().Rot(rad(180), 0, 0),	// bottom
		*top.Clone().Rot(rad(90), 0, 0),	// north
		*top.Clone().Rot(rad(-90), 0, 0),	// south
		*top.Clone().Rot(0, rad(90), 0),	// west
		*top.Clone().Rot(0, rad(-90), 0),	// east
	)
	return cube
}

func main() {
	var model Model
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		model = *NewSTLReader(f).ReadModel(true)
		f.Close()
	} else {
		model = *Cube().Rot(math.Pi / 4, math.Pi / 4, math.Pi / 4)
	}

	err := ui.Main(func() {

		renderer := Renderer{
			a:    nil,
			dp:        nil,
			projector: *NewProjector(600, 52),
			model: model,
			cameraMatrix: *TransM(NewV4(0, 0, 2)),
			rotTime: 30,	// seconds per full rotation
		}
		canvas := ui.NewArea(&renderer)
		renderer.a = canvas

		box := ui.NewVerticalBox()
		box.Append(canvas, true)
		window := ui.NewWindow("Perspective Projection", renderer.projector.size, renderer.projector.size, false)
		window.SetMargined(false)
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()

		go func () {
			renderer.mainLoop()
		}()
	})
	if err != nil {
		panic(err)
	}
}
