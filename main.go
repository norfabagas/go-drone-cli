package main

import (
	"fmt"
	"strconv"
)

/*
Use fmt.Scanln to read input from console
and convert it to type uint64
 */
func InputUint() uint64 {
	var input string
	var output uint64
	fmt.Scanln(&input)
	if val, err := strconv.ParseUint(input, 10, 32); err != nil || val < 1 {
		fmt.Println(err)
		fmt.Print("Try another value: ")
		InputUint()
	} else {
		output = val
	}
	return output
}

/*
Generate message for displaying
drone location
input current(Coordinate), min(Coordinate), max(Coordinate), direction(string)
 */
func PrintDroneStatus(current Coordinate, min Coordinate, max Coordinate, direction string) {
	fmt.Printf("Current Position: (%d, %d)\n", current.X, current.Y)
	fmt.Printf("Min Coordinate: (%d, %d)\n", min.X, min.Y)
	fmt.Printf("Max Coordinate: (%d, %d)\n", max.X, max.Y)
	fmt.Printf("The drone is pointing to: %s\n", direction)
}

func main() {
		fmt.Print("Set Drone Max X point: ")
		pointX := InputUint()

		fmt.Print("Set Drone Max Y point: ")
		pointY := InputUint()

	drone := Drone{
		MinPoint:     Coordinate{0,0},
		MaxPoint:     Coordinate{pointX, pointY},
		CurrentPoint: Coordinate{0,0},
		Face:         1,
	}
	fmt.Println("Drone defined")

	menu:
		fmt.Println("====================")
		fmt.Println("Type menu available:")
		fmt.Println("" +
			"- exit\n" +
			"- move (Move drone)\n" +
			"- direction (Change drone direction)\n" +
			"- location (Get current drone location)")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "exit":
			fmt.Println("Exit...")
			return
		case "direction":
			direction:
				fmt.Print("Input direction: ", DroneFace())
				direction := InputUint()
				if direction > 4 {
					goto direction
				}
				moveStatus := drone.SetDroneFace(uint(direction))
				if moveStatus {
					fmt.Println("Change drone position")
					PrintDroneStatus(drone.GetCurrentPosition())
				} else {
					fmt.Println("Undefined drone direction")
				}
				goto menu
		case "move":
			fmt.Print("Input number of steps: ")
			steps := InputUint()
			status, message := drone.SetDronePosition(steps)
			if status {
				fmt.Println("Drone position changed")
				PrintDroneStatus(drone.GetCurrentPosition())
			} else {
				fmt.Println(message)
			}
			goto menu
		case "location":
			PrintDroneStatus(drone.GetCurrentPosition())
			goto menu
		default:
			fmt.Println("Input undefined")
			goto menu
		}
}