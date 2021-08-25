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
	Strategy interface {
		chooseLot(availableLot []*ParkingLot) *ParkingLot
	}
)

type Vehicle struct {
	i int
}

type ParkingLot struct {
	parkingSpace map[*Vehicle]bool
	Capacity     int
	subscribers  []Subscriber
}

func (lot ParkingLot) isAvailable() bool {
	return lot.getAvailableSlots() > 0
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func NewParkingLot(capacity int, subscriberList []Subscriber) *ParkingLot {
	return &ParkingLot{Capacity: capacity, subscribers: subscriberList, parkingSpace: make(map[*Vehicle]bool)}
}

func (lot ParkingLot) getAvailableSlots() int {
	return lot.Capacity - len(lot.parkingSpace)
}

func (lot *ParkingLot) NotifyAllSubs(functionName string) {
	for _, subscriber := range lot.subscribers {
		args := []reflect.Value{reflect.ValueOf(lot)}
		reflect.ValueOf(subscriber).MethodByName(functionName).Call(args)

	}
}

func (lot *ParkingLot) Park(vehicle *Vehicle) error {
	if len(lot.parkingSpace) == lot.Capacity {
		return errors.New("parking is full. Cannot park vehicle")
	}
	if lot.parkingSpace[vehicle] == true {
		return errors.New("cannot park already parked vehicle")
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
	return errors.New("vehicle has not been parked yet")
}

func (lot *ParkingLot) addSubscriber(attendant *Attendant) {
	lot.subscribers = append(lot.subscribers, attendant)
}
