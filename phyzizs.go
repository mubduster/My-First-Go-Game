package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const rectangleMass float32 = 4.0
const circleMass float32= 1.0 
const bouncy float32 = 0.7

func HandleCollisionBetweenRecAndCircle(circleCenter rl.Vector2, radius float32, rec rl.Rectangle, speedX, speedY, circleSpeedX, circleSpeedY float32) (float32 , float32, float32, float32, float32, float32, bool) {

	if !(rl.CheckCollisionCircleRec(circleCenter, radius, rec)) {return speedX, speedY, circleCenter.X, circleCenter.Y, circleSpeedX, circleSpeedY, false}
	
	closestXPoint := Clamp(circleCenter.X, rec.X, rec.X+rec.Width)
	closestYPoint := Clamp(circleCenter.Y, rec.Y, rec.Y + rec.Height)
	
	diffX := circleCenter.X - closestXPoint
	diffY := circleCenter.Y - closestYPoint

	distance := float32(rl.Vector2Length(rl.NewVector2(diffX, diffY)))

	newSpeedX := speedX
	newSpeedY := speedY
	newCircleSpeedX := circleSpeedX
	newCircleSpeedY := circleSpeedY

	overlap := radius - distance

	if distance == 0 {return speedX, speedY, circleCenter.X, circleCenter.Y, circleSpeedX, circleSpeedY, false}

	normalX := diffX / distance
	normalY := diffY / distance

	relativeVelocityX := circleSpeedX - speedX
	relativeVelocityY := circleSpeedY - speedY

	velocityAlongNormal := (relativeVelocityX * normalX) + (relativeVelocityY * normalY)

	if velocityAlongNormal < 0 {

		impulseScalar := (-(1 + bouncy) * velocityAlongNormal) / (1/rectangleMass + 1/circleMass)
		
		newSpeedX -= (1/rectangleMass) * impulseScalar * normalX
		newSpeedY -= (1/rectangleMass * impulseScalar * normalY)

		newCircleSpeedX += (1/circleMass) * impulseScalar * normalX	
		newCircleSpeedY += (1/circleMass) * impulseScalar * normalY

	}

	circleCenter.X += normalX * overlap
	circleCenter.Y += normalY * overlap

	return newSpeedX, newSpeedY, circleCenter.X, circleCenter.Y, newCircleSpeedX, newCircleSpeedY, true	
}

func Clamp (val, min, max float32) float32{
	if val < min {return min}
	if val > max {return max}
	return val
}

func Abs (val float32) float32{
	if val < 0 {return -val}
 	return val
}