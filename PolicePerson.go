package parkingProject

type PolicePerson struct {
	IsFull bool
}

func NewPolicePerson() *PolicePerson {
	return &PolicePerson{IsFull: false}
}

func (p *PolicePerson) NotifyIsFull(*ParkingLot) {
	p.IsFull = true
}

func (p *PolicePerson) NotifyIsAvailable(*ParkingLot) {
	p.IsFull = false
}
