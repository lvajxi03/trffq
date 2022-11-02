package game

import (
	"github.com/lvajxi03/bdlist"
)

type LaneType int

const (
	LANE_NS LaneType = iota
	LANE_WE
	LANE_S
	LANE_W
	LANE_N
	LANE_E
)

const (
	CAR_MARGIN = 5
)

type Lane struct {
	Terminated               bool
	Cars *bdlist.BDList
	Cross1, Cross2 *Crossroad
	Length uint32
	Capacity uint32
	Type LaneType
}

func NewLane(lt LaneType) *Lane {
	lane := &Lane{
		Type: lt,
		Cars: bdlist.New(),
	}
	return lane
}

func (lane *Lane) AppendCar(car *Car) bool {
	if lane != nil && car != nil {
		if lane.Length < lane.Capacity + car.Length + CAR_MARGIN {
			// TODO
		}
	}
	return false
}

func (lane *Lane) PopCar() *Car {
	if lane != nil {
		
		car, _ := lane.Cars.RemoveAt(0)
		return car.(*Car)
	}
	return nil
}
