package types

import "fmt"

// Position represents an X, Y coordinate on the grid
type Position struct {
	X, Y int
}

// String returns a string representation of the position
func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Add returns a new position by adding two positions together
func (p Position) Add(other Position) Position {
	return Position{X: p.X + other.X, Y: p.Y + other.Y}
}

// Equals checks if two positions are equal
func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

// Direction vectors for movement
var (
	Up    = Position{X: 0, Y: -1}
	Down  = Position{X: 0, Y: 1}
	Left  = Position{X: -1, Y: 0}
	Right = Position{X: 1, Y: 0}
)

