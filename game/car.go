package game

type Direction int

const (
	TOP   Direction = iota
	DOWN            = iota
	LEFT            = iota
	RIGHT           = iota
)

type Car struct {
	Started, Visible bool
	X, Y             int // top-left corner
	Direction        Direction
	Length, Width    uint32
}

func NewCar(dir Direction) *Car {
	car := &Car{Started: false, Visible: false}
	if dir > RIGHT {
		dir = RIGHT
	}
	if dir < TOP {
		dir = TOP
	}
	car.Direction = dir
	return car
}
