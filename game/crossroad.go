package game

type CrossroadType uint8

const (
	CR_FULL CrossroadType = iota
	CR_RIGHT
	CR_TOP
	CR_LEFT
	CR_BOTTOM
)

type Crossroad struct {
	Left, Right, Top, Bottom *Crossroad
	Type CrossroadType
}
