package main

import (
	"testing"
)

func TestDrone_SetDroneFace(t *testing.T) {
	drone := Drone{
		MinPoint:     Coordinate{0,0},
		MaxPoint:     Coordinate{10,10},
		CurrentPoint: Coordinate{0,0},
		Face:         1,
	}

	direction := DroneFace()

	drone.SetDroneFace(2)
	if direction[2] != direction[drone.Face] {
		t.Error("Drone face incorrect")
	}
}

func TestDrone_SetDronePosition(t *testing.T) {
	drone := Drone{
		MinPoint:     Coordinate{0,0},
		MaxPoint:     Coordinate{10,10},
		CurrentPoint: Coordinate{5,5},
		Face:         2,
	}

	//test for false condition
	//expected result = false
	result, _ := drone.SetDronePosition(10)
	if result {
		t.Error("Drone failed negative test")
	}

	//test for true condition
	currentCoordinate, minCoordinate, _, _ := drone.GetCurrentPosition()
	result, _ = drone.SetDronePosition(currentCoordinate.Y)
	if !result {
		t.Error("Drone failed positive test")
	}
	if drone.CurrentPoint.Y != minCoordinate.Y {
		t.Error("Coordinate value not match")
	}
}