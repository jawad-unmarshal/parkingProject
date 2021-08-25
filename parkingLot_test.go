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
		err := parkingLot.Park(vehicle)
		isParked := parkingLot.IsParked(vehicle)

		assert.True(t, isParked, err)
	})

	t.Run("Expect parked vehicle to be unparked", func(t *testing.T) {
		var vehicle = NewVehicle()
		owner := NewOwner()
		subList := make([]Subscriber, 1)
		subList[0] = owner
		parkingLot := NewParkingLot(2, subList)
		_ = parkingLot.Park(vehicle)
		_ = parkingLot.UnPark(vehicle)
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
		_ = ParkLot.Park(vehicle1)
		_ = ParkLot.Park(vehicle2)
		_ = ParkLot.Park(vehicle3)
		_ = ParkLot.UnPark(vehicle2)
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
			_ = parkingLot.Park(vehicle)
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
		_ = parkingLot.Park(vehicle1)
		vehicle2 := NewVehicle()
		_ = parkingLot.Park(vehicle2)

		_ = parkingLot.UnPark(vehicle1)

		result := owner.IsFull

		assert.Equal(t, false, result, "Parking lot should have space")
	})

	t.Run("Notify Policeman that slot is full", func(t *testing.T) {
		policeMen := NewPolicePerson()
		var subList = make([]Subscriber, 1)
		subList[0] = policeMen
		parkingLot := NewParkingLot(2, subList)

		vehicle1 := NewVehicle()
		_ = parkingLot.Park(vehicle1)
		vehicle2 := NewVehicle()
		_ = parkingLot.Park(vehicle2)

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
		strategy := Strategy(NewHighestAvailabilityStrategy())
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle1 := NewVehicle()

		parkResult := attendant.Park(vehicle1)

		if parkResult != nil {
			t.Fatal("attendant failed to park: ", parkResult)
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
		strategy := Strategy(NewHighestAvailabilityStrategy())
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle1 := NewVehicle()
		_ = attendant.Park(vehicle1)
		parkResult := attendant.UnPark(vehicle1)

		if parkResult != nil {
			t.Fatal("attendant failed to Unpark: ", parkResult)
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
		strategy := Strategy(NewHighestAvailabilityStrategy())
		attendant := NewAttendant(parkingLotList, strategy)
		//parkingLot.addSubscriber(attendant)
		vehicle1 := NewVehicle()
		_ = attendant.Park(vehicle1)

		result := attendant.IsFull

		assert.Equal(t, true, result, "attendant should be notified")

	})

	t.Run("Expect attendant to manage multiple parkinglots", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot1 := NewParkingLot(0, subList)
		parkingLot2 := NewParkingLot(3, subList)
		parkingLotList := make([]*ParkingLot, 2)
		parkingLotList[0] = parkingLot1
		parkingLotList[1] = parkingLot2
		strategy := Strategy(NewHighestAvailabilityStrategy())
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle := NewVehicle()
		_ = attendant.Park(vehicle)

		result := parkingLot2.IsParked(vehicle)

		assert.Equal(t, true, result, "vehicle should be parked in Lot 2")

	})

	t.Run("Expect attendant to direct vehicle to parking lot with highest free space", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot1 := NewParkingLot(2, subList)
		parkingLot2 := NewParkingLot(3, subList)
		parkingLotList := make([]*ParkingLot, 2)
		parkingLotList[0] = parkingLot1
		parkingLotList[1] = parkingLot2
		strategy := Strategy(NewHighestAvailabilityStrategy())
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle := NewVehicle()
		_ = attendant.Park(vehicle)

		result := parkingLot2.IsParked(vehicle)

		assert.Equal(t, true, result, "vehicle should be parked in lot 2")

	})

	t.Run("Expect attendant to direct vehicle to parking lot with highest free space", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot1 := NewParkingLot(5, subList)
		parkingLot2 := NewParkingLot(3, subList)
		strategy := Strategy(NewHighestAvailabilityStrategy())
		parkingLotList := make([]*ParkingLot, 2)
		parkingLotList[0] = parkingLot1
		parkingLotList[1] = parkingLot2
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle := NewVehicle()
		_ = attendant.Park(vehicle)

		result := parkingLot1.IsParked(vehicle)

		assert.Equal(t, true, result, "vehicle should be parked in lot 1")

	})

	t.Run("Expect attendant to direct vehicle to parking lot with highest capacity among the available", func(t *testing.T) {
		var subList = make([]Subscriber, 1)
		owner := NewOwner()
		subList[0] = owner
		parkingLot1 := NewParkingLot(2, subList)
		parkingLot2 := NewParkingLot(1, subList)
		var strategy = Strategy(NewHighestCapacityStrategy())
		parkingLotList := make([]*ParkingLot, 2)
		parkingLotList[0] = parkingLot1
		parkingLotList[1] = parkingLot2
		attendant := NewAttendant(parkingLotList, strategy)
		vehicle := NewVehicle()
		_ = attendant.Park(vehicle)

		result := parkingLot1.IsParked(vehicle)

		assert.Equal(t, true, result, "vehicle should be parked in lot 1")

	})

}
