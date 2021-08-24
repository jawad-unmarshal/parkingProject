package parkingProject

type Vehicle struct {
	Id int
}

type ParkingLot struct {
	ParkingSpace []*Vehicle
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func (parkLot *ParkingLot) Park(vehicle *Vehicle) {
	parkLot.ParkingSpace = append(parkLot.ParkingSpace, vehicle)
}

func (parkLot *ParkingLot) IsParked(vehicle *Vehicle) bool {
	for _, parkedVehicle := range parkLot.ParkingSpace {
		if parkedVehicle == vehicle {
			return true
		}
	}
	return false
}

func (parkLot *ParkingLot) UnPark(vehicle *Vehicle) {
	for i, parkedVehicle := range parkLot.ParkingSpace {
		if parkedVehicle == vehicle {
			parkLot.ParkingSpace = append(parkLot.ParkingSpace[:i], parkLot.ParkingSpace[i+1:]...)
		}
	}

}
