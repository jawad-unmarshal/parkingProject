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

type Attendant struct {
	IsFull        bool
	parkingLots   []*ParkingLot
	availableLots []*ParkingLot
}

type ParkingLot struct {
	parkingSpace   map[*Vehicle]bool
	availableSlots int
	subscribers    []Subscriber
}

type PolicePerson struct {
	IsFull bool
}

func NewAttendant(parkingLots []*ParkingLot) *Attendant {
	var availableLots = make([]*ParkingLot, 0)
	RetAttendant := Attendant{parkingLots: parkingLots}
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
	return lot.availableSlots > 0
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

func NewParkingLot(availableSlots int, subscriberList []Subscriber) *ParkingLot {
	return &ParkingLot{availableSlots: availableSlots, subscribers: subscriberList, parkingSpace: make(map[*Vehicle]bool)}
}

func (o *Owner) NotifyIsFull(lot *ParkingLot) {
	o.IsFull = true
}

func (o *Owner) NotifyIsAvailable(lot *ParkingLot) {
	o.IsFull = false
}

func (p *PolicePerson) NotifyIsFull(lot *ParkingLot) {
	p.IsFull = true
}

func (p *PolicePerson) NotifyIsAvailable(lot *ParkingLot) {
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
	max := -1
	var maxLot *ParkingLot
	for _, lot := range a.availableLots {
		if lot.availableSlots > max {
			max = lot.availableSlots
			maxLot = lot
		}
	}

	return maxLot.Park(vehicle)

}

func (a *Attendant) UnPark(vehicle *Vehicle) error {
	for _, lot := range a.parkingLots {
		if lot.IsParked(vehicle) {
			return lot.UnPark(vehicle)
		}
	}
	return errors.New("Vehicle not found.")
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

func (lot *ParkingLot) NotifyAllSubs(functionName string) {
	for _, subscriber := range lot.subscribers {
		args := []reflect.Value{reflect.ValueOf(lot)}
		reflect.ValueOf(subscriber).MethodByName(functionName).Call(args)

	}
}

func (lot *ParkingLot) Park(vehicle *Vehicle) error {
	if lot.availableSlots == 0 {
		return errors.New("Parking is full. Cannot park vehicle")
	}
	if lot.parkingSpace[vehicle] == true {
		return errors.New("Cannot park already parked vehicle")
	}
	lot.parkingSpace[vehicle] = true
	lot.availableSlots--
	if lot.availableSlots == 0 {
		lot.NotifyAllSubs("NotifyIsFull")
		//P.NotifyAllSubsIsFull()
	}
	return nil
}

func (lot *ParkingLot) IsParked(vehicle *Vehicle) bool {
	return lot.parkingSpace[vehicle]
}

func (lot *ParkingLot) UnPark(vehicle *Vehicle) error {
	if lot.parkingSpace[vehicle] {
		lot.parkingSpace[vehicle] = false
		lot.availableSlots++
		if lot.availableSlots == 1 {
			lot.NotifyAllSubs("NotifyIsAvailable")
		}
		return nil
	}
	return errors.New("Vehicle has not been parked yet")
}

func (lot *ParkingLot) addSubscriber(attendant *Attendant) {
	lot.subscribers = append(lot.subscribers, attendant)
}
