package parkingProject

import "errors"

type Attendant struct {
	IsFull        bool
	parkingLots   []*ParkingLot
	availableLots []*ParkingLot
	strategy      Strategy
}

func NewAttendant(parkingLots []*ParkingLot, strategy Strategy) *Attendant {
	var availableLots = make([]*ParkingLot, 0)
	RetAttendant := Attendant{parkingLots: parkingLots, strategy: strategy}
	for _, lot := range parkingLots {
		lot.addSubscriber(&RetAttendant)
		if lot.isAvailable() {
			availableLots = append(availableLots, lot)
		}
	}
	RetAttendant.availableLots = availableLots
	return &RetAttendant
}

func (a *Attendant) Park(vehicle *Vehicle) error {
	if len(a.availableLots) == 0 {
		return errors.New("all Parking lots are full")
	}
	chooseStrategy := a.strategy
	maxLot := chooseStrategy.chooseLot(a.availableLots)

	return maxLot.Park(vehicle)

}

func (a *Attendant) UnPark(vehicle *Vehicle) error {
	for _, lot := range a.parkingLots {
		if lot.IsParked(vehicle) {
			return lot.UnPark(vehicle)
		}
	}
	return errors.New("vehicle not found")
}

func (a *Attendant) NotifyIsFull(lotToRemove *ParkingLot) {
	a.IsFull = true
	availList := a.availableLots
	for i, parkingLot := range availList {
		if parkingLot == lotToRemove {
			availList[i] = availList[len(availList)-1]
			availList = availList[:len(availList)-1]
			break
		}
	}

}

func (a *Attendant) NotifyIsAvailable(lotToAdd *ParkingLot) {
	a.IsFull = false
	a.availableLots = append(a.availableLots, lotToAdd)
}
