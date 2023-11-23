// Bloodmage Engine
// Copyright (C) 2023 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package engine

import (
	"math"

	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	// Time in seconds since last frame
	DeltaTime      float64
	bkg            sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	title          string
	width          int32   = 800
	widthF64       float64 = float64(width)
	halfWidthF64   float64 = widthF64 / 2
	height         int32   = 800
	heightF64      float64 = float64(height)
	halfHeightF64  float64 = heightF64 / 2
	centerX        int32   = width / 2
	centerY        int32   = height / 2
	screenDist     float64 = 0.5
	renderer       *sdl.Renderer
	window         *sdl.Window
	frameStartTime uint64
	keyStates      = sdl.GetKeyboardState()
	running        bool
	cursorLocked   bool
)

const (
	targetFrameTime uint64 = 1000 / 60
)

func Start(t string) {
	if window != nil {
		log.Fatal("window already started")
	}
	title = t
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")
	var err error
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	windowFlags := uint32(sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
	if options.Fullscreen {
		windowFlags |= sdl.WINDOW_FULLSCREEN
	}
	window, err = sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		windowFlags)
	if err != nil {
		log.Fatal(err)
	}

	rendererFlags := uint32(sdl.RENDERER_ACCELERATED)
	if options.Vsync {
		rendererFlags |= sdl.RENDERER_PRESENTVSYNC
	}
	renderer, err = sdl.CreateRenderer(window, -1, rendererFlags)
	if err != nil {
		log.Fatal(err)
	}
	running = true
	updateWindowSize()
	frameStartTime = sdl.GetTicks64()
}

func Window() *sdl.Window {
    return window
}

func updateWindowSize() {
	width, height = window.GetSize()
	widthF64 = float64(width)
	halfWidthF64 = widthF64 / 2
	heightF64 = float64(height)
	halfWidthF64 = widthF64 / 2
	screenDist = halfWidthF64 / math.Tan(halfFov)

	centerX = width / 2
	centerY = height / 2

	numOfRays = width / options.PixelScale
	deltaAngle = fov / (widthF64 / float64(options.PixelScale))
	scale = width / int32(numOfRays)
}

func Stop() {
	running = false
	_ = window.Destroy()
	_ = renderer.Destroy()
	sdl.Quit()
}

func input() {
	keyStates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			running = false
		case sdl.WINDOWEVENT:
			updateWindowSize()
		}
	}
	getMouseInput()
}

func beginRender() {
	if err := renderer.SetDrawColor(bkg.R, bkg.G, bkg.B, bkg.A); err != nil {
		log.Error(err)
	}
	if err := renderer.Clear(); err != nil {
		log.Error(err)
	}
}

func Running() bool {
	now := sdl.GetTicks64()
	DeltaTime = float64(now-frameStartTime) / 1000.0
	frameStartTime = now

	input()
	beginRender()
	return running
}

func Present() {
	renderer.Present()
	frameTime := sdl.GetTicks64() - frameStartTime
	if frameTime < targetFrameTime {
		sdl.Delay(uint32(targetFrameTime - frameTime))
	}
}
