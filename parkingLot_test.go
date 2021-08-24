package parkingProject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPark(t *testing.T) {

	t.Run("Expecting vehicle to be parked", func(t *testing.T) {
		var vehicle = NewVehicle()
		owner := NewOwner()
		subList := make([]Subscriber, 1)
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		parkingLot.Park(vehicle)
		isParked := parkingLot.IsParked(vehicle)

		assert.True(t, isParked)
	})

	t.Run("Expect parked vehicle to be unparked", func(t *testing.T) {
		var vehicle = NewVehicle()
		owner := NewOwner()
		subList := make([]Subscriber, 1)
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		parkingLot.Park(vehicle)
		parkingLot.UnPark(vehicle)
		isParked := parkingLot.IsParked(vehicle)

		assert.False(t, isParked)
	})

	t.Run("Expect error when trying to unpark a vehicle that is not parked", func(t *testing.T) {
		var vehicle = NewVehicle()
		var ParkLot ParkingLot
		result := ParkLot.UnPark(vehicle)

		if result == nil {
			t.Fatalf("Should not be able to unpark a vehicle that is not parked")
		}
	})

	t.Run("Expect first parked vehicle to be unparked", func(t *testing.T) {
		var vehicle1 = NewVehicle()
		var vehicle2 = NewVehicle()
		var vehicle3 = NewVehicle()
		var ParkLot ParkingLot
		ParkLot.Park(vehicle1)
		ParkLot.Park(vehicle2)
		ParkLot.Park(vehicle3)
		ParkLot.UnPark(vehicle2)
		isParked := ParkLot.IsParked(vehicle2)

		assert.False(t, isParked)
	})

	t.Run("Expect parkinglot with capacity of 2 to be full after 2 vehicles are parked", func(t *testing.T) {
		owner := NewOwner()
		subList := make([]Subscriber, 1)
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		for i := 0; i < 2; i++ {
			vehicle := NewVehicle()
			parkingLot.Park(vehicle)
		}

		result := owner.IsFull

		assert.True(t, result, "Parking lot should be full")
	})

	t.Run("Expect to notify owner when slots become available", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)

		vehicle1 := NewVehicle()
		parkingLot.Park(vehicle1)
		vehicle2 := NewVehicle()
		parkingLot.Park(vehicle2)

		parkingLot.UnPark(vehicle1)

		result := owner.IsFull

		assert.Equal(t, false, result, "Parking lot should have space")
	})

	t.Run("Notify Policeman that slot is full", func(t *testing.T) {
		policeMen := NewPolicePerson()
		var subList = make([]Subscriber, 1)
		subList[0] = policeMen
		parkingLot := NewParkingLot(2, subList)

		vehicle1 := NewVehicle()
		parkingLot.Park(vehicle1)
		vehicle2 := NewVehicle()
		parkingLot.Park(vehicle2)

		result := policeMen.IsFull
		assert.Equal(t, true, result, "Parking lot should be full")

	})

	t.Run("Expect an attendant to park vehicle", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		parkingLotList := make([]*ParkingLot, 1)
		parkingLotList[0] = parkingLot
		attendant := NewAttendant(parkingLotList)
		vehicle1 := NewVehicle()

		parkResult := attendant.Park(vehicle1)

		if parkResult != nil {
			t.Fatal("attendant failed to park: ", parkResult)
			t.FailNow()
		}

		isParked := parkingLot.IsParked(vehicle1)

		assert.Equal(t, true, isParked, "Vehicle should be parked")

	})

	t.Run("Expect an attendant to Unpark vehicle", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		parkingLotList := make([]*ParkingLot, 1)
		parkingLotList[0] = parkingLot
		attendant := NewAttendant(parkingLotList)
		vehicle1 := NewVehicle()
		attendant.Park(vehicle1)
		parkResult := attendant.UnPark(vehicle1)

		if parkResult != nil {
			t.Fatal("attendant failed to Unpark: ", parkResult)
			t.FailNow()
		}

		isParked := parkingLot.IsParked(vehicle1)

		assert.Equal(t, false, isParked, "Vehicle should be Unparked")

	})

	t.Run("Expect attendant to receive notification", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot := NewParkingLot(1, subList)
		parkingLotList := make([]*ParkingLot, 1)
		parkingLotList[0] = parkingLot
		attendant := NewAttendant(parkingLotList)
		parkingLot.addSubscriber(attendant)

		vehicle1 := NewVehicle()
		attendant.Park(vehicle1)

		result := attendant.IsFull

		assert.Equal(t, true, result, "attendant should be notified")

	})

}
