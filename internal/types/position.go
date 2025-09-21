package types

import "fmt"

type Position struct {
	X, Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Position) Add(other Position) Position {
	return Position{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

var (
	Up    = Position{X: 0, Y: -1}
	Down  = Position{X: 0, Y: 1}
	Left  = Position{X: -1, Y: 0}
	Right = Position{X: 1, Y: 0}
)

