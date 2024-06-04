package main

import (
	"fmt"
	"realChakrawarti/idiomatic/parking-receiver"
	game "realChakrawarti/idiomatic/player-receiver"
	idiomatic "realChakrawarti/idiomatic/receiver"
)

func receiverFuncs() {
	coords := idiomatic.Coordinate{X: 5, Y: 5}

	// Without pointer receiver function
	idiomatic.ShiftCoords(1, 1, &coords)
	fmt.Println(coords) // 6,6

	// With pointer receiver function
	coords.ShiftCoordsReceiver(2, 3)
	fmt.Println(coords) // 8,9

	// With value receiver function
	first := idiomatic.Coordinate{X: 2, Y: 2}
	second := idiomatic.Coordinate{X: 4, Y: 5}

	displacement := first.DistanceCoords(second)

	fmt.Println(displacement) // 2,3
}

func parkingDemo() {

	lot := make([]parking.Space, 4)
	lotSpaces := parking.ParkingLot{Spaces: lot}
	fmt.Println("Parking Lot (Before):", lotSpaces)
	lotSpaces.OccupySpace(1)
	parking.OccupySpace(&lotSpaces, 2)
	fmt.Println("Parking Lot (After occupying):", lotSpaces)
	lotSpaces.VacateSpace(1)
	fmt.Println("Parking Lot (After vacate):", lotSpaces)
}

func gameDemo() {
	hawkEye := game.NewPlayer("Hawk-eye", 100, 100)
	hawkEye.UpdateHealth(45)
	hawkEye.UpdateEnergy(95)
}

func main() {
	receiverFuncs()
	parkingDemo()
	gameDemo()
}
