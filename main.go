package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)


const (
	screenWidth  = 800
	screenHeight = 600
	length = 60*2
	thickness = 3.0
	branchCount = 7
	branchAngle = float64(2.0*math.Pi/branchCount)
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Go")
	defer rl.CloseWindow()
	
	rl.SetTargetFPS(60)
	
	depth := 4
	
	center := rl.NewVector2(screenWidth/2, screenHeight/2)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		hue := float32(rl.GetTime()*5)
		rl.ClearBackground(rl.Color{R: 15, G: 13, B: 13, A: 255})
		drawFlakeRecursive(center, length, thickness, depth, hue)
		rl.EndDrawing()
	}
}

func drawFlakeRecursive(center rl.Vector2, length float64, thick float32, depth int, hue float32) {
	if depth <= 0 {
		return
	}
	for i := 0; i < branchCount; i ++ {
		to := rl.NewVector2(center.X + float32(length*math.Cos(branchAngle * float64(i))), center.Y + float32(length*math.Sin(branchAngle * float64(i))))
		rl.DrawLineEx(center, to, thick, rl.ColorFromHSV(hue, 1.0, 0.6))
		drawFlakeRecursive(to, length/2, thick/2, depth-1, (hue+10.0)*2)
	}
}
