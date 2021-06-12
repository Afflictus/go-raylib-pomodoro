package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - text writing anim")

	rl.SetTargetFPS(120)

	var vsp, hsp int32 = 1, 1
	positionX := screenWidth / 2
	positionY := screenHeight / 2
	for !rl.WindowShouldClose() {
		// Update
		if rl.IsKeyDown(rl.KeyD) {
			positionX += hsp
		}
		if rl.IsKeyDown(rl.KeyA) {
			positionX -= hsp
		}
		if rl.IsKeyDown(rl.KeyW) {
			positionY -= vsp
		}
		if rl.IsKeyDown(rl.KeyS) {
			positionY += vsp
		}
		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(positionX, positionY, 50, 50, rl.Gray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func clamp()
