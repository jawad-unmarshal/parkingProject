package parkingProject

type Vehicle struct {
}

type ParkingLot struct {
	ParkingSpace []*Vehicle
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func (P *ParkingLot) Park(vehicle *Vehicle) {
	P.ParkingSpace = append(P.ParkingSpace, vehicle)
}

func (P *ParkingLot) IsParked(vehicle *Vehicle) bool {
	for _, parkedVehicle := range P.ParkingSpace {
		if parkedVehicle == vehicle {
			return true
		}
	}
	return false
}

func (P *ParkingLot) UnPark(vehicle *Vehicle) {
	for i, parkedVehicle := range P.ParkingSpace {
		if parkedVehicle == vehicle {
			P.ParkingSpace = append(P.ParkingSpace[:i], P.ParkingSpace[1+i:]...)
		}
	}

}
