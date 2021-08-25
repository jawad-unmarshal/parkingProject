package parkingProject

import (
	"errors"
	"reflect"
)

type (
	Subscriber interface {
		NotifyIsFull(lot *ParkingLot)
		NotifyIsAvailable(lot *ParkingLot)
	}
)

type Vehicle struct {
	i int
}
type Owner struct {
	IsFull bool
	//Attendant
}

type Strategy interface {
	chooseLot(availableLot []*ParkingLot) *ParkingLot
}

type Attendant struct {
	IsFull        bool
	parkingLots   []*ParkingLot
	availableLots []*ParkingLot
	strategy      Strategy
}

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

type ParkingLot struct {
	parkingSpace map[*Vehicle]bool
	Capacity     int
	subscribers  []Subscriber
}

type PolicePerson struct {
	IsFull bool
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
func (lot ParkingLot) isAvailable() bool {
	return lot.getAvailableSlots() > 0
}

func NewPolicePerson() *PolicePerson {
	return &PolicePerson{IsFull: false}
}

func NewOwner() *Owner {
	return &Owner{IsFull: false}
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func NewParkingLot(capacity int, subscriberList []Subscriber) *ParkingLot {
	return &ParkingLot{Capacity: capacity, subscribers: subscriberList, parkingSpace: make(map[*Vehicle]bool)}
}

func (o *Owner) NotifyIsFull(*ParkingLot) {
	o.IsFull = true
}

func (o *Owner) NotifyIsAvailable(*ParkingLot) {
	o.IsFull = false
}

func (p *PolicePerson) NotifyIsFull(*ParkingLot) {
	p.IsFull = true
}

func (p *PolicePerson) NotifyIsAvailable(*ParkingLot) {
	p.IsFull = false
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

func (a *Attendant) Park(vehicle *Vehicle) error {
	if len(a.availableLots) == 0 {
		return errors.New("all Parking lots are full")
	}
	chooseStrategy := a.strategy
	maxLot := chooseStrategy.chooseLot(a.availableLots)

	return maxLot.Park(vehicle)

}

func (lot ParkingLot) getAvailableSlots() int {
	return lot.Capacity - len(lot.parkingSpace)
}

func (a *Attendant) UnPark(vehicle *Vehicle) error {
	for _, lot := range a.parkingLots {
		if lot.IsParked(vehicle) {
			return lot.UnPark(vehicle)
		}
	}
	return errors.New("Vehicle not found.")
}

func (lot *ParkingLot) NotifyAllSubs(functionName string) {
	for _, subscriber := range lot.subscribers {
		args := []reflect.Value{reflect.ValueOf(lot)}
		reflect.ValueOf(subscriber).MethodByName(functionName).Call(args)

	}
}

func (lot *ParkingLot) Park(vehicle *Vehicle) error {
	if len(lot.parkingSpace) == lot.Capacity {
		return errors.New("Parking is full. Cannot park vehicle")
	}
	if lot.parkingSpace[vehicle] == true {
		return errors.New("Cannot park already parked vehicle")
	}
	lot.parkingSpace[vehicle] = true
	if lot.Capacity == len(lot.parkingSpace) {
		lot.NotifyAllSubs("NotifyIsFull")
	}
	return nil
}

func (lot *ParkingLot) IsParked(vehicle *Vehicle) bool {
	return lot.parkingSpace[vehicle]
}

func (lot *ParkingLot) UnPark(vehicle *Vehicle) error {
	if lot.parkingSpace[vehicle] {
		//lot.parkingSpace[vehicle] = false
		delete(lot.parkingSpace, vehicle)
		if len(lot.parkingSpace) == 1 {
			lot.NotifyAllSubs("NotifyIsAvailable")
		}
		return nil
	}
	return errors.New("Vehicle has not been parked yet")
}

func (lot *ParkingLot) addSubscriber(attendant *Attendant) {
	lot.subscribers = append(lot.subscribers, attendant)
}
