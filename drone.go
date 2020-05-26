package main

/*
Coordinate tuple (x,y)
 */
type Coordinate struct {
	X	uint64
	Y	uint64
}

/*
Drone struct with
MinPoint and MaxPoint defines drone borderline in (x,y)
CurrentPoint defines drone current position in (x,y)
Face defines drone direction
 */
type Drone struct {
	MinPoint     Coordinate
	MaxPoint     Coordinate
	CurrentPoint Coordinate
	Face         uint
}

/*
Get current drone position
based on CurrentPoint and Face value
return current(Coordinate), min(Coordinate), max(Coordinate), string
 */
func (drone *Drone) GetCurrentPosition() (Coordinate, Coordinate, Coordinate, string) {
	return drone.CurrentPoint, drone.MinPoint, drone.MaxPoint, DroneFace()[drone.Face]
}

/*
Change drone face direction
Input: integer 1-4
Output: bool
 */
func (drone *Drone) SetDroneFace(direction uint) bool {
	droneFace := DroneFace()
	if _, status := droneFace[direction]; status {
		drone.Face = direction
		return status
	} else {
		return status
	}
}

/*
Move drone from current position
Input: steps(uint64)
Output: bool, string
 */
func (drone *Drone) SetDronePosition(steps uint64) (bool, string) {
	switch drone.Face {
	case 1:
		if result := drone.CurrentPoint.Y + steps; result <= drone.MaxPoint.Y {
			drone.CurrentPoint.Y = result
			return true, ""
		} else {
			return false, "Drone cannot move outside of Max Y Point"
		}
	case 2:
		if steps <= drone.CurrentPoint.Y {
			drone.CurrentPoint.Y = drone.CurrentPoint.Y - steps
			return true, ""
		} else {
			return false, "Drone cannot move outside of Min Y Point"
		}
	case 3:
		if steps <= drone.CurrentPoint.X {
			drone.CurrentPoint.X = drone.CurrentPoint.X - steps
			return true, ""
		} else {
			return false, "Drone cannot move outside of Min X Point"
		}
	case 4:
		if result := drone.CurrentPoint.X + steps; result <= drone.MaxPoint.Y {
			drone.CurrentPoint.X = result
			return true, ""
		} else {
			return false, "Drone cannot move outside of Max X Point"
		}
	default:
		return false, "Undefined drone direction"
	}
}

/*
Generate available drone directions
Output: map[uint]string
 */
func DroneFace() map[uint]string {
	return map[uint]string{
		1: "UP",
		2: "DOWN",
		3: "LEFT",
		4: "RIGHT",
	}
}