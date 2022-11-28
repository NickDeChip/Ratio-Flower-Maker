package main

import (
	"fmt"
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	winScl         = 800
	screenCenter   = winScl / 2
	spokesRadius   = 10
	startingRadius = 10
)

var (
	spokes []rl.Vector2 = make([]rl.Vector2, 800)
	radius float32      = startingRadius
	ratio  float64      = 0
)

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(winScl, winScl, "ratio shapes")
	rl.SetTargetFPS(60)

	generateSpokes()
	for !rl.WindowShouldClose() {

		update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func update() {
	mouseWheelMove := rl.GetMouseWheelMove()
	if mouseWheelMove != 0 {
		radius = startingRadius
		ratio += float64(mouseWheelMove * 0.00025)
		generateSpokes()
	}
	if rl.IsKeyPressed(rl.KeyG) {
		radius = startingRadius
		ratio = math.Phi
		generateSpokes()
	}
	if rl.IsKeyPressed(rl.KeyP) {
		radius = startingRadius
		ratio = math.Pi
		generateSpokes()
	}
}

func generateSpokes() {
	deg := 360 * ratio
	degAfter := float64(0.0)
	for i := 0; i < len(spokes); i++ {
		if degAfter > 360 {
			degAfter = degAfter - 360
		}
		angle := ((math.Pi * float64(degAfter)) / 180.0)
		x := screenCenter + (radius * float32(math.Sin(angle)))
		y := screenCenter + (radius * float32(math.Cos(angle)))

		spokes[i] = rl.NewVector2(x, y)

		degAfter += deg
		d := radius * 2
		c := math.Pi * d
		c += 10
		r := c / (2 * math.Pi)
		radius = r
	}
}

func draw() {
	for i := 0; i < len(spokes); i++ {
		rl.DrawCircleLines(screenCenter, screenCenter, radius, rl.Black)
		rl.DrawCircleV(spokes[i], spokesRadius, rl.Purple)
		rl.DrawCircleLines(int32(spokes[i].X), int32(spokes[i].Y), spokesRadius, rl.Black)
		rl.DrawText(fmt.Sprintf("Ratio: %f", ratio), 20, 20, 28, rl.Gray)
	}
}
