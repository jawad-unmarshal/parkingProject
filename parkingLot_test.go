package parkingProject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPark(t *testing.T) {

	t.Run("Expecting vehicle to be parked", func(t *testing.T) {
		var vehicle = NewVehicle()
		var ParkLot ParkingLot
		ParkLot.Park(vehicle)
		isParked := ParkLot.IsParked(vehicle)

		assert.True(t, isParked)
	})

	t.Run("Expect parked vehicle to be unparked", func(t *testing.T) {
		var vehicle = NewVehicle()
		var ParkLot ParkingLot
		ParkLot.Park(vehicle)
		ParkLot.UnPark(vehicle)
		isParked := ParkLot.IsParked(vehicle)

		assert.False(t, isParked)
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
		isParked := ParkLot.IsParked(vehicle1)

		assert.False(t, isParked)
	})

}
