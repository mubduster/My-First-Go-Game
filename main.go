package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// CJ
var speedX float32 
var speedY float32 
var circleSpeedX float32 
var circleSpeedY float32 
var newCircleSpeedX float32
var newCircleSpeedY float32
var shift uint8 = 1
var shiftDir uint8 = 1
var hueShift float32 
var dir float32 = 0.4
var R uint8 = 224
var G uint8 = 224
var B uint8 = 224
var wallX bool
var wallY bool
var grav bool = true
// var circleGrav bool = true
var bR uint8 = 0
var bG uint8 = 0
var bB uint8 = 0
var gR uint8 = 128
var gG uint8 = 128
var gB uint8 = 128
var shiftBorderColor uint8 = 1
var shiftGridColor uint8 = 1
var shiftBorderDir uint8 = 1
var shiftGridDir uint8 = 1
var maxSpeed float32 = 5000
var acceleration float32 = 1500
var drag float32 = 1100
var circleDrag float32 = 250
var gravity float32
var gravityForce float32 = 50
var maxGravity float32 = 3000
var newSpeedX float32
var newSpeedY float32
var didBounce bool = false
var gravityCircle float32 


func main(){
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized) // Game window customization
	rl.InitWindow(0, 0, "Basic Window") // Open new game window
	
	defer rl.CloseWindow() // Close window when closed or Esc pressed

	rl.SetTargetFPS(60)

	const worldWidth float32 = 6500
	const worldHeight float32 = 6500

	const gridSpacing float32 = 200
	const padding float32 = 1000
	
	x := worldWidth - 200
	y := worldHeight - 200
	
	var preX float32 = x
	var preY float32 = y

	rectangleTexture := rl.LoadTexture("./textures/player.png")
	defer rl.UnloadTexture(rectangleTexture)
	
	stageCircle := rl.NewVector2(worldWidth/2 , worldHeight - 100)
	radius := float32(100)

	
	for !rl.WindowShouldClose() {
		// screen setup ------------------------------------------------------------------------------------------------------------------------------------------
		screenWidth, screenHeight := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
		halfScreenWidth, halfScreenHeight := int32(screenWidth/2), int32(screenHeight/2)
		//--------------------------------------------------------------------------------------------------------------------------------------------------------
		
		dT := rl.GetFrameTime()

		cursorAreaCircle := rl.GetMousePosition()
		
		// rectangle drag ----------------------------------------------------------------------------------------------------------------------------------------
		if !(rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) && !(rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)) {
			if speedX > 0 {
				speedX -= drag * dT
				if speedX < 0 {
					speedX = 0
				}
			}else if speedX < 0 {
				speedX += drag * dT
				if speedX > 0 {
					speedX = 0
				}
			}
		}
		if !(rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW)) && !(rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS)) {
			if speedY > 0 && !grav{
				speedY -= drag * dT
				if speedY < 0 {
					speedY = 0
				}
			}else if speedY < 0 {
				speedY += drag * dT
				if speedY > 0 {
					speedY = 0
				}
			}
		}
		//---------------------------------------------------------------------------------------------------------------------------------------------------------
		// circle drag --------------------------------------------------------------------------------------------------------------------------------------------
		if circleSpeedX > 0 {
			circleSpeedX -= circleDrag * dT
			if circleSpeedX < 0 {
				circleSpeedX = 0
			}
		}else if circleSpeedX < 0 {
			circleSpeedX += circleDrag * dT
			if circleSpeedX > 0 {
				circleSpeedX = 0
			}
		}
		if circleSpeedY > 0{
			circleSpeedY -= circleDrag * dT
			if circleSpeedY < 0 {
				circleSpeedY = 0
			}
		}else if circleSpeedY < 0 {
			circleSpeedY += circleDrag * dT
			if circleSpeedY > 0 {
				circleSpeedY = 0
			}
		}
		//---------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// movement controls + max movement speed -----------------------------------------------------------------------------------------------------------------
		if (rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW)) && y > 0 && int32(speedY) >= -int32(maxSpeed){
			speedY -= acceleration * dT
		}
		if (rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS)) && y+200 < worldHeight && int32(speedY) <= int32(maxSpeed){
			speedY += acceleration * dT
		}
		if (rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) && x+300 < worldWidth && int32(speedX) <= int32(maxSpeed){
			speedX += acceleration * dT
		}
		if (rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)) && x > 0 && int32(speedX) >= -int32(maxSpeed){
			speedX -= acceleration * dT
		}
		//----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// color shifting ------------------------------------------------------------------------------------------------------------------------------------------
		R ,G ,B = R+shiftDir, G+shiftDir, B+shiftDir
		bR, bG, bB = bR+shiftBorderDir, bG+shiftBorderDir, bB+shiftBorderDir
		gR, gG, gB = gR+shiftGridDir, gG+shiftGridDir, gB+shiftGridDir
		hueShift += dir * dT

		if R >= 225 && G >= 225 && B >= 225 || R <= 0 && G <= 0 && B <= 0 {
			shift += -shiftDir
		}
		if bR >= 225 && bG >= 225 && bB >= 225 || bR <= 0 && bG <= 0 && bB <= 0 {
			shiftBorderColor += -shiftBorderDir
		}
		if gR >= 225 && gG >= 225 && gB >= 225 || gR <= 0 && gG <= 0 && gB <= 0 {
			shiftBorderColor += -shiftBorderDir
		}

		if hueShift >= 360.0 || hueShift <= 0 {
			dir = dir * -dir
		}
		//----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// border collision rectangle ------------------------------------------------------------------------------------------------------------------------------
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
		//----------------------------------------------------------------------------------------------------------------------------------------------------------
		// border collision circle ---------------------------------------------------------------------------------------------------------------------------------
		if stageCircle.X + radius > worldWidth {
			stageCircle.X = worldWidth - radius
			circleSpeedX = circleSpeedX * -1
		}else if stageCircle.X - radius < 0 {
			stageCircle.X = radius
			circleSpeedX = circleSpeedX * -1
		}else if stageCircle.Y + radius > worldHeight {
			stageCircle.Y = worldHeight - radius	
			circleSpeedY = circleSpeedY * -1
		}else if stageCircle.Y - radius < 0 {
			stageCircle.Y = radius
			circleSpeedY = circleSpeedY * -1
		}
		//----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// gravity rectangle ---------------------------------------------------------------------------------------------------------------------------------------
		if (y+20 <= worldHeight && speedY >= -0.5) && !((x + 300 > stageCircle.X - 10 && x < stageCircle.X + 10) && (y + 200 > stageCircle.Y - radius - 0.4 && y + 200 < stageCircle.Y -radius +0.4)){
			grav = true
		}else {
			grav = false
		}
		if rl.IsKeyDown(rl.KeyUp){
			grav = false
		}
		
		if !rl.IsKeyDown(rl.KeyUp) && gravity < maxGravity && grav && speedY < maxSpeed{
			gravity += gravityForce
			speedY += gravity * dT
		}else if gravity > (maxGravity-1) && grav && speedY < maxSpeed{
			speedY += gravity * dT
		}else if !grav{
			gravity = 0.0
		}
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		// gravity circle -------------------------------------------------------------------------------------------------------------------------------------------
		if !((stageCircle.X + radius > x && stageCircle.X - radius < x + 300) && (stageCircle.Y + radius > y - 10 && stageCircle.Y + radius < y + 10)){
			if gravityCircle < maxGravity && circleSpeedY < maxSpeed && stageCircle.Y + radius > worldHeight - 1 {
				gravityCircle += gravityForce
				circleSpeedY += gravityCircle * dT
			}else if gravityCircle > (maxGravity-1) && speedY < maxSpeed{
				circleSpeedY += gravityCircle * dT
			}
		}
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// camera setup ---------------------------------------------------------------------------------------------------------------------------------------------
		boxPos := rl.NewVector2(x + 150 ,y + 100)
		cameraOffset := rl.NewVector2(screenWidth/2, screenHeight/2)
		
		camera := rl.Camera2D{
			Target: boxPos,
			Offset: cameraOffset,
			Rotation: 0.0,
			Zoom: 0.8,
		}
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// border vetor setup ---------------------------------------------------------------------------------------------------------------------------------------
		borderStart := rl.NewVector2(-10,-10)
		borderEnd := rl.NewVector2(worldWidth+30,worldHeight+30)
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------

		// move player according to speed ---------------------------------------------------------------------------------------------------------------------------
		x += speedX * dT
		y += speedY * dT
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// move stage circle according to speed ---------------------------------------------------------------------------------------------------------------------
		stageCircle.X += circleSpeedX * dT
		stageCircle.Y += circleSpeedY * dT
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------

		// player Texture -------------------------------------------------------------------------------------------------------------------------------------------
		textureSource := rl.NewRectangle(0, 0, float32(rectangleTexture.Width), float32(rectangleTexture.Height))
		player := rl.NewRectangle(x, y, 300, 200)
		origin := rl.NewVector2(0, 0)
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		
		// collision check ------------------------------------------------------------------------------------------------------------------------------------------
		newSpeedX, newSpeedY, stageCircle.X, stageCircle.Y, newCircleSpeedX, newCircleSpeedY, didBounce= HandleCollisionBetweenRecAndCircle(stageCircle, radius, player, speedX, speedY, circleSpeedX, circleSpeedY)

		if didBounce {
			speedX = newSpeedX
			speedY = newSpeedY
			circleSpeedX = newCircleSpeedX
			circleSpeedY = newCircleSpeedY
			
			x = preX
			y = preY
		}
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
		

		// render ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.BeginDrawing() // start frame 
		
		// render objects here needed to be with the camera----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.ClearBackground(rl.NewColor(R, G, B, 225))

		Border := rl.NewRectangle(borderStart.X, borderStart.Y, borderEnd.X, borderEnd.Y)
		
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	
		

		rl.BeginMode2D(camera) // camera setup + start frame
		//render objects here to stay with the stage-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		
		for gridV := -padding; gridV <= worldWidth+padding; gridV += gridSpacing {
			lineStart := rl.NewVector2(gridV, -padding)
			lineEnd := rl.NewVector2(gridV, worldHeight+padding)
			rl.DrawLineEx(lineStart, lineEnd, 5.0, rl.NewColor(gR, gG, gB, 255))
		}
		for gridH := -padding; gridH <= worldWidth+padding; gridH += gridSpacing {
			lineStart := rl.NewVector2(-padding, gridH)
			lineEnd	:= rl.NewVector2(worldWidth+padding, gridH)
			rl.DrawLineEx(lineStart, lineEnd, 5.0, rl.NewColor(gR, gG, gB, 255))
		}

		rl.DrawRectangle(int32(preX), int32(preY), 300, 200, rl.GetColor(0x000044ff))
		rl.DrawRectangle(int32(x), int32(y), 300, 200, rl.ColorFromHSV(hueShift,0.0,0.0))
		
		rl.DrawTexturePro(rectangleTexture, textureSource, player, origin, 0.0, rl.White)
		
		rl.DrawText("Bouncy ractnagle", int32(x)+15, int32(y+80), 30, rl.ColorFromHSV(hueShift/2,-1.0,-1.0))

		rl.DrawCircleV(stageCircle, radius, rl.Red)

		rl.DrawRectangleLinesEx(Border, 20.0, rl.NewColor(bR, bG, bB, 255))
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		
		rl.EndMode2D() // camera end frame

		// render objects here needed to be with the camera----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.DrawText(fmt.Sprintf("Screen Width: %0.1f, Screen Height: %0.1f. halfScreenWidth: %0.1f, halfScreenHeight: %0.1f", screenWidth, screenHeight, float32(halfScreenWidth), float32(halfScreenHeight)), 10 , 40, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Use arrow keys to move the rectangle!"), 10, 90, 30, rl.GetColor(0x00FFFF77))
		rl.DrawText(fmt.Sprintf("Speed X : %0.1f\nSpeed Y: %0.1f\nX: %0.1f\nY: %0.1f\nGravity: %0.1f\nGrav: %t:", speedX, speedY, x+150, y+100,gravity, grav), 10, 120, 30, rl.GetColor(0x00FFFF77))
		rl.DrawFPS(10, 10)
		rl.DrawCircleV(cursorAreaCircle, 20, rl.GetColor(0xffff0044))
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.EndDrawing() // end frame
		//------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


		// trailing rectangle movement ------------------------------------------------------------------------------------------------------------------------------
		preX = x
		preY = y
		//-----------------------------------------------------------------------------------------------------------------------------------------------------------
	}	
}

