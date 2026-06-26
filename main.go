package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// CJ
var speedX float32 = 0
var speedY float32 = 0
var shift uint8 = 1
var hueShift float32 = 0.0
var dir float32 = 0.4
var R uint8 = 224
var G uint8 = 224
var B uint8 = 224
var wallX bool
var wallY bool
var gravity float32 = 0.0
var grav bool = true
var bR uint8 = 0
var bG uint8 = 0
var bB uint8 = 0
var gR uint8 = 128
var gG uint8 = 128
var gB uint8 = 128
var shiftBorderColor uint8 = 1
var shiftGridColor uint8 = 1


func main(){
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Basic Window")
	
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	
	const worldWidth float32 = 6000
	const worldHeight float32 = 6000

	const gridSpacing float32 = 200
	const padding float32 = 1000
	
	x := worldWidth - 200
	y := worldHeight - 200

	var preX float32 = x
	var preY float32 = y
	
	for !rl.WindowShouldClose() {
		screenWidth, screenHeight := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
		halfScreenWidth, halfScreenHeight := int32(screenWidth/2), int32(screenHeight/2)

		
		hueShift += dir

		if !rl.IsKeyDown(rl.KeyRight) && !rl.IsKeyDown(rl.KeyLeft) {
			if speedX > 0 {
				speedX -= 1
			}else if speedX < 0 {
				speedX += 1
			}
		}
		if !rl.IsKeyDown(rl.KeyUp) && !rl.IsKeyDown(rl.KeyDown) {
			if speedY > 0 && !grav{
				speedY -= 1
			}else if speedY < 0 {
				speedY += 1
			}
		}
	
		if rl.IsKeyDown(rl.KeyUp) && y > 0{
			speedY -= 1
		}
		if rl.IsKeyDown(rl.KeyDown) && y+200 < worldHeight{
			speedY += 1
		}
		if rl.IsKeyDown(rl.KeyRight) && x+300 < worldWidth{
			speedX += 1
		}
		if rl.IsKeyDown(rl.KeyLeft) && x > 0{
			speedX -= 1
		}

		R ,G ,B = R+shift, G+shift, B+shift
		bR, bG, bB = bR+shiftBorderColor, bG+shiftBorderColor, bB+shiftBorderColor
		gR, gG, gB = gR+shiftGridColor, gG+shiftGridColor, gB+shiftGridColor
		
		// border collision-----------------------------------------
		if x+300 > worldWidth {
			x = worldWidth - 300
			speedX = speedX * -1 
		}else if x < 0 {
			x = 0
			speedX = speedX * -1
		}
		if y+200 > worldHeight{
			y = worldHeight - 200
			speedY = speedY * -0.7
		}else if y < 0 {
			y = 0
			speedY = speedY * -1
		}
		//-----------------------------------------------------------

		if y+200 <= worldHeight && int32(speedY) >= -1{
			grav = true
		}else {
			grav = false
		}
		if rl.IsKeyDown(rl.KeyUp){
			grav = false
		}

		if !rl.IsKeyDown(rl.KeyUp) && gravity < 3.0 && grav {
			gravity += 0.025
			speedY += gravity 
		}else if gravity > 2.9 && grav{
			speedY += gravity
		}else if !grav{
			gravity = 0.0
		}

		x += speedX
		y += speedY
		
		if R >= 225 && G >= 225 && B >= 225 || R <= 0 && G <= 0 && B <= 0 {
			shift = -shift
		}
		if bR >= 225 && bG >= 225 && bB >= 225 || bR <= 0 && bG <= 0 && bB <= 0 {
			shiftBorderColor = -shiftBorderColor
		}
		if gR >= 225 && gG >= 225 && gB >= 225 || gR <= 0 && gG <= 0 && gB <= 0 {
			shiftBorderColor = -shiftBorderColor
		}

		if hueShift >= 360.0 || hueShift <= 0 {
			dir = dir * -dir
		}

		boxPos := rl.NewVector2(x,y)
		cameraOffset := rl.NewVector2(screenWidth/2, screenHeight/2)

		camera := rl.Camera2D{
			Target: boxPos,
			Offset: cameraOffset,
			Rotation: 0.0,
			Zoom: 1.0,
		}

		borderStart := rl.NewVector2(0,0)
		borderEnd := rl.NewVector2(worldWidth,worldHeight)

		rl.BeginDrawing()
		// render objects here needed to be with the camera----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.ClearBackground(rl.NewColor(R, G, B, 225))

		Border := rl.NewRectangle(borderStart.X, borderStart.Y, borderEnd.X, borderEnd.Y)
		
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	
		rl.BeginMode2D(camera)
		//render objects here to stay with the stage-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

		for gridV := -padding; gridV <= worldWidth+padding; gridV += gridSpacing {
			lineStart := rl.NewVector2(gridV, -padding)
			lineEnd := rl.NewVector2(gridV, worldHeight+padding)
			rl.DrawLineEx(lineStart, lineEnd, 5.0, rl.NewColor(gR, gG, gB, 225))
		}
		for gridH := -padding; gridH <= worldWidth+padding; gridH += gridSpacing {
			lineStart := rl.NewVector2(-padding, gridH)
			lineEnd	:= rl.NewVector2(worldWidth+padding, gridH)
			rl.DrawLineEx(lineStart, lineEnd, 5.0, rl.NewColor(gR, gG, gB, 225))
		}

		rl.DrawRectangle(int32(preX), int32(preY), 300, 200, rl.ColorFromHSV(hueShift,0.2,0.7))
		rl.DrawRectangle(int32(x), int32(y), 300, 200, rl.ColorFromHSV(hueShift,0.5,0.7))
		rl.DrawText("Bouncy ractnagle", int32(x)+15, int32(y+80), 30, rl.ColorFromHSV(hueShift/2,1.0,1.0))

		rl.DrawRectangleLinesEx(Border, 10.0, rl.NewColor(bR, bG, bB, 225))
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.EndMode2D()
		// render objects here needed to be with the camera----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.DrawText(fmt.Sprintf("Screen Width: %0.1f, Screen Height: %0.1f. halfScreenWidth: %0.1f, halfScreenHeight: %0.1f", screenWidth, screenHeight, float32(halfScreenWidth), float32(halfScreenHeight)), 10 , 40, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Use arrow keys to move the rectangle!"), 10, 90, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Speed X : %0.1f\nSpeed Y: %0.1f\nX: %0.1f\nY: %0.1f\nGravity: %0.1f\nGrav: %t:", speedX, speedY, x+150, y+100,gravity, grav), 10, 120, 30, rl.GetColor(0x00FFFF77))
		rl.DrawFPS(10, 10)

		
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.EndDrawing()

		preX = x
		preY = y
	}	
}

