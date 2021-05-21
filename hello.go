package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	fnt := rl.LoadFont("Exo.ttf")
	//rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		
		rl.DrawTextEx(fnt, "Congrats! ТЫ created your first window!", rl.NewVector2(float32(190), float32(200)), float32(32), float32(2), rl.LightGray)
		rl.DrawText("Congrats! ТЫ created your first window!", 190, 230, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}