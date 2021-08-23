package parkingProject

import (
	"errors"
	"reflect"
)

type Subscriber interface {
	NotifyIsFull()
	NotifyIsAvailable()
}

type Vehicle struct {
	i int
}
type Owner struct {
	IsFull bool
	Attendant
}

type Attendant struct {
	*ParkingLot
}

func NewAttendant(parkingLot *ParkingLot) *Attendant {
	return &Attendant{ParkingLot: parkingLot}
}

func (a Attendant) Park(vehicle *Vehicle) error {
	lot := a.ParkingLot
	return lot.Park(vehicle)

}

type PolicePerson struct {
	IsFull bool
}

func NewPolicePerson() *PolicePerson {
	return &PolicePerson{IsFull: false}
}

func NewOwner() *Owner {
	return &Owner{IsFull: false}
}

type ParkingLot struct {
	parkingSpace   []*Vehicle
	availableSlots int
	subscribers    []Subscriber
}

func NewParkingLot(availableSlots int, subscriberList []Subscriber) *ParkingLot {
	return &ParkingLot{availableSlots: availableSlots, subscribers: subscriberList}
}
func (o *Owner) NotifyIsFull() {
	o.IsFull = true
}

func (o *Owner) NotifyIsAvailable() {
	o.IsFull = false
}

func (p *PolicePerson) NotifyIsFull() {
	p.IsFull = true
}

func (p *PolicePerson) NotifyIsAvailable() {
	p.IsFull = false
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

//func (P ParkingLot) NotifyAllSubsIsFull() {
//	for _, subscriber := range P.subscribers {
//		subscriber.notifyIsFull()
//
//	}
//}
//func (P ParkingLot) NotifyAllSubsIsAvailable() {
//	for _, subscriber := range P.subscribers {
//		subscriber.NotifyIsAvailable()
//
//	}
//}

func (P ParkingLot) NotifyAllSubs(functionName string) {
	for _, subscriber := range P.subscribers {
		reflect.ValueOf(subscriber).MethodByName(functionName).Call(nil)

	}

}

func (P *ParkingLot) Park(vehicle *Vehicle) error {
	if P.availableSlots == 0 {
		return errors.New("Parking is full. Cannot park vehicle")
	}
	P.parkingSpace = append(P.parkingSpace, vehicle)
	P.availableSlots--
	if P.availableSlots == 0 {
		P.NotifyAllSubs("NotifyIsFull")
		//P.NotifyAllSubsIsFull()
	}
	return nil
}

func (P *ParkingLot) IsParked(vehicle *Vehicle) bool {
	for _, parkedVehicle := range P.parkingSpace {
		if parkedVehicle == vehicle {
			return true
		}
	}
	return false
}

func (P *ParkingLot) UnPark(vehicle *Vehicle) error {
	for i, parkedVehicle := range P.parkingSpace {
		if parkedVehicle != vehicle {
			continue
		}

		P.unparkVehicleAt(i)
		P.availableSlots++
		if P.availableSlots == 1 {
			P.NotifyAllSubs("NotifyIsAvailable")
		}
		return nil
	}
	return errors.New("Vehicle has not been parked yet")
}

func (P *ParkingLot) unparkVehicleAt(i int) {
	P.parkingSpace[i] = P.parkingSpace[len(P.parkingSpace)-1]
	P.parkingSpace = P.parkingSpace[:len(P.parkingSpace)-1]
}
