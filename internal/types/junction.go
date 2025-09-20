package types

import "fmt"

type Junction struct {
	X, Y       int
	Directions []Position
	ActiveDir  int
	ID         rune
}

func NewJunction(x, y int, directions []Position, id rune) *Junction {
	return &Junction{
		X:          x,
		Y:          y,
		Directions: directions,
		ActiveDir:  0,
		ID:         id,
	}
}

func (j *Junction) SwitchRoute() {
	if len(j.Directions) > 0 {
		j.ActiveDir = (j.ActiveDir + 1) % len(j.Directions)
	}
}

func (j *Junction) GetActiveDirection() Position {
	if len(j.Directions) == 0 {
		return Position{0, 0}
	}
	return j.Directions[j.ActiveDir]
}

func (j *Junction) GetSymbol() rune {
	activeDir := j.GetActiveDirection()
	switch activeDir {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	default:
		return '+'
	}
}

func (j *Junction) String() string {
	return fmt.Sprintf("Junction{ID: %c, Pos: (%d,%d), Routes: %d, Active: %d}",
		j.ID, j.X, j.Y, len(j.Directions), j.ActiveDir)
}

