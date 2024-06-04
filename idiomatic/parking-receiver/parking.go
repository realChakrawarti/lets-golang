package parking

type Space struct {
	occupied bool
}

type ParkingLot struct {
	Spaces []Space
}

func OccupySpace(lot *ParkingLot, spaceNum int) {
	lot.Spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) OccupySpace(spaceNum int) {
	lot.Spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) VacateSpace(spaceNum int) {
	lot.Spaces[spaceNum-1].occupied = false
}
