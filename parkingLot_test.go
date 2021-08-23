package parkingProject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPark(t *testing.T) {

	t.Run("Expecting vehicle to be parked", func(t *testing.T) {
		var vehicle = NewVehicle("")

		Park(*vehicle)
		isParked := IsParked(*vehicle)

		assert.True(t, isParked)
	})

	t.Run("Expect vehicle to be unparked", func(t *testing.T) {
		var vehicle = NewVehicle("")

		UnPark(*vehicle)
		isParked := IsParked(*vehicle)

		assert.False(t, isParked)
	})

}
