package main

import (
	"fmt"
	"runtime"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Printf("Hello %s/%s\n", runtime.GOOS, runtime.GOARCH)
	rl.InitWindow(300, 300, "[TIMER]")
	rl.SetTargetFPS(60)
	//rl.SetWindowState(rl.FlagWindowAlwaysRun)

	// Time stuff
	var work, rest int = 1 * 60, 1 * 60
	var current int = work
	working := true
	pause := true
	timerSecond := 60

	// Sound stuff
	rl.InitAudioDevice()
	ring := rl.LoadSound("ring.mp3")
	volume := float32(0.2)

	// Text stuff
	var startBtnStatus string

	var vec rl.Vector2
	var clock string // := fmt.Sprintf("%2d:%2d", work/60, work%60)
	font := rl.GetFontDefault()

	var vecMonkey rl.Vector2
	fontMonkey := rl.LoadFont("Segment7-4Gml.otf")

	workColor := rl.Red
	restColor := rl.Green
	baseColor := rl.Black

	for !rl.WindowShouldClose() {
		// Update
		clock = fmt.Sprintf("%02d:%02d", current/60, current%60)
		vec = rl.Vector2{300/2 - rl.MeasureTextEx(font, clock, 42, 2).X/2, 50}
		vecMonkey = rl.Vector2{300/2 - rl.MeasureTextEx(fontMonkey, clock, 32, 1).X/2, 120}

		if !pause {
			if working {
				startBtnStatus = "[WORKING]"
				baseColor = workColor
				if current <= 0 {
					current = rest
					working = !working
					pause = true
					rl.PlaySound(ring)
				}
			} else if !working {
				startBtnStatus = "[REST]"
				baseColor = restColor
				if current <= 0 {
					current = work
					working = !working
					pause = true
					rl.PlaySound(ring)
				}
			}
			if current > 0 {
				if timerSecond > 0 {
					timerSecond -= 1
				} else {
					timerSecond = 60
					current -= 1
				}
			}
		} else {
			startBtnStatus = "[PAUSE]"
		}

		rl.SetMasterVolume(volume)
		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawTextEx(font, clock, vec, 42, 2, baseColor)
		rl.DrawTextEx(fontMonkey, clock, vecMonkey, 32, 1, baseColor)

		rl.DrawText("VOLUME", 10, 210, 20, rl.Gray)
		volume = rg.SliderBar(rl.NewRectangle(110, 210, 180, 20), volume, 0.0, 1.0)
		startStop := rg.Button(rl.NewRectangle(10, 240, 280, 50), startBtnStatus)

		if startStop {
			pause = !pause
		}

		//  Drag and Drop stuff
		//// INIT
		//dragging := false
		//var mx, my float32 = -1, -1
		//var mx, my int = -1, -1
		//// STEP
		// if dragging {
		// 	x := mx + rl.GetMousePosition().X
		// 	y := my + rl.GetMousePosition().Y
		// 	// x := mx + int(rl.GetMouseX())
		// 	// y := my + int(rl.GetMouseY())
		// 	fmt.Println(rl.GetWindowPosition(), rl.GetMousePosition(), x, y)
		// 	rl.SetWindowPosition(int(x), int(y))
		// 	if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
		// 		dragging = false
		// 	}
		// } else {
		// 	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		// 		dragging = true
		// 		mx = rl.GetWindowPosition().X - rl.GetMousePosition().X
		// 		my = rl.GetWindowPosition().Y - rl.GetMousePosition().Y
		// 		// mx = int(rl.GetWindowPosition().X) - int(rl.GetMouseX())
		// 		// my = int(rl.GetWindowPosition().Y) - int(rl.GetMouseY())
		// 	}
		// }
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
