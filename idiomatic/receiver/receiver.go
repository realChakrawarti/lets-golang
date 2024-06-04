package idiomatic

type Coordinate struct {
	X, Y int
}

// Pointer Receiver

func ShiftCoords(x, y int, coords *Coordinate) {
	coords.X += x
	coords.Y += y
}

// Can modify a struct
func (coords *Coordinate) ShiftCoordsReceiver(x, y int) {
	coords.X += x
	coords.Y += y
}

// Value Reciever

// Cannot modify a struct, that's why need to return a new one
func (c Coordinate) DistanceCoords(other Coordinate) Coordinate {
	return Coordinate{
		other.X - c.X, other.Y - c.Y,
	}
}
