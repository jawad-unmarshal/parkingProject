package parkingProject

type HighestCapacityStrategy struct {
	i int
}

func NewHighestCapacityStrategy() HighestCapacityStrategy {
	return HighestCapacityStrategy{}
}
func (s HighestCapacityStrategy) chooseLot(availableLots []*ParkingLot) *ParkingLot {
	max := -1
	var maxLot *ParkingLot
	for _, lot := range availableLots {
		if lot.Capacity > max {
			max = lot.Capacity
			maxLot = lot
		}
	}
	return maxLot

}

type HighestAvailabilityStrategy struct {
	i int
}

func NewHighestAvailabilityStrategy() HighestAvailabilityStrategy {
	return HighestAvailabilityStrategy{}
}

func (s HighestAvailabilityStrategy) chooseLot(availableLots []*ParkingLot) *ParkingLot {
	max := -1
	var maxLot *ParkingLot
	for _, lot := range availableLots {
		if lot.getAvailableSlots() > max {
			max = lot.getAvailableSlots()
			maxLot = lot
		}
	}
	return maxLot

}
