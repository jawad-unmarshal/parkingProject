package parkingProject

type Owner struct {
	IsFull bool
}

func NewOwner() *Owner {
	return &Owner{IsFull: false}
}

func (o *Owner) NotifyIsFull(*ParkingLot) {
	o.IsFull = true
}

func (o *Owner) NotifyIsAvailable(*ParkingLot) {
	o.IsFull = false
}
