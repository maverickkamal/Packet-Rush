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

// IsValid checks if the position is within the given bounds
func (p Position) IsValid(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// Direction vectors for movement
var (
	Up    = Position{X: 0, Y: -1}
	Down  = Position{X: 0, Y: 1}
	Left  = Position{X: -1, Y: 0}
	Right = Position{X: 1, Y: 0}
)

// GetDirectionFromChar converts a character to a direction
func GetDirectionFromChar(char rune) Position {
	switch char {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	default:
		return Position{0, 0}
	}
}

// GetCharFromDirection converts a direction to a character
func GetCharFromDirection(dir Position) rune {
	switch dir {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	default:
		return '?'
	}
}

