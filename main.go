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
var R uint8 = 0
var G uint8 = 0
var B uint8 = 0
var wallX bool
var wallY bool
var gravity float32 = 0.0
var grav bool = true

func main(){
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Basic Window")

	screenWidth, screenHeight := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())

	fmt.Printf("ScreenWidth: %f, ScreenHeight: %f.",screenWidth, screenHeight)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	halfScreenWidth, halfScreenHeight := int32(screenWidth/2), int32(screenHeight/2)
	
	x := float32(screenWidth - 200)
	y := float32(screenHeight - 200)
	
	for !rl.WindowShouldClose() {
		
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
		if rl.IsKeyDown(rl.KeyDown) && y+200 < screenHeight{
			speedY += 1
		}
		if rl.IsKeyDown(rl.KeyRight) && x+300 < screenWidth{
			speedX += 1
		}
		if rl.IsKeyDown(rl.KeyLeft) && x > 0{
			speedX -= 1
		}

		R ,G ,B = R+shift, G+shift, B+shift
		
		if x+300 > screenWidth {
			x = screenWidth - 300
			speedX = speedX * -1 
		}else if x < 0 {
			x = 0
			speedX = speedX * -1
		}
		if y+200 > screenHeight{
			y = screenHeight - 200
			speedY = speedY * -1
		}else if y < 0 {
			y = 0
			speedY = speedY * -1
		}

		if y+200 <= screenHeight && int32(speedY) >= -1{
			grav = true
		}else {
			grav = false
		}
		if rl.IsKeyDown(rl.KeyUp){
			grav = false
		}

		if !rl.IsKeyDown(rl.KeyUp) && gravity < 2.0 && grav {
			gravity += 0.1
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

		if hueShift >360.0 {
			dir = -dir
		}

		rl.BeginDrawing()
		
		rl.ClearBackground(rl.NewColor(R, G, B, 225))
		rl.DrawRectangle(int32(x), int32(y), 300, 200, rl.ColorFromHSV(hueShift,0.5,0.7))
		rl.DrawText("Bouncy ractnagle", int32(x)+15, int32(y+80), 30, rl.ColorFromHSV(hueShift/2,1.0,1.0))
		rl.DrawText(fmt.Sprintf("Screen Width: %0.1f, Screen Height: %0.1f. halfScreenWidth: %0.1f, halfScreenHeight: %0.1f", screenWidth, screenHeight, float32(halfScreenWidth), float32(halfScreenHeight)), 10 , 40, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Use arrow keys to move the rectangle!"), 10, 90, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Speed X : %0.1f\nSpeed Y: %0.1f \nGravity: %0.1f\nGrav: %t", speedX, speedY, gravity, grav), 10, 120, 30, rl.GetColor(0x00FFFF77))
		rl.DrawFPS(10, 10)
		
		rl.EndDrawing()
	}	
}

