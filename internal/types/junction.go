package types

import "fmt"

// Junction represents a switchable point on the map
type Junction struct {
	Pos         Position   // Position on the grid
	Routes      []Position // Possible exit directions
	ActiveRoute int        // Index of currently active route
	ID          rune       // Junction identifier for keyboard control (1-9)
	Symbol      rune       // Visual symbol for the junction
}

// NewJunction creates a new junction with the given parameters
func NewJunction(pos Position, routes []Position, id rune) *Junction {
	return &Junction{
		Pos:         pos,
		Routes:      routes,
		ActiveRoute: 0, // Start with first route
		ID:          id,
		Symbol:      JunctionSymbol, // Default junction symbol
	}
}

// SwitchRoute changes to the next available route
func (j *Junction) SwitchRoute() {
	if len(j.Routes) > 0 {
		j.ActiveRoute = (j.ActiveRoute + 1) % len(j.Routes)
	}
}

// SetRoute sets the active route to a specific index
func (j *Junction) SetRoute(routeIndex int) {
	if routeIndex >= 0 && routeIndex < len(j.Routes) {
		j.ActiveRoute = routeIndex
	}
}

// GetActiveDirection returns the currently active direction
func (j *Junction) GetActiveDirection() Position {
	if len(j.Routes) == 0 {
		return Position{0, 0}
	}
	return j.Routes[j.ActiveRoute]
}

// GetRouteDirection returns a specific route direction by index
func (j *Junction) GetRouteDirection(index int) Position {
	if index >= 0 && index < len(j.Routes) {
		return j.Routes[index]
	}
	return Position{0, 0}
}

// GetSymbol returns the visual representation of the junction
func (j *Junction) GetSymbol() rune {
	// Show different symbols based on active direction for visual feedback
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

// String returns a string representation of the junction
func (j *Junction) String() string {
	return fmt.Sprintf("Junction{ID: %c, Pos: %s, Routes: %d, Active: %d}",
		j.ID, j.Pos, len(j.Routes), j.ActiveRoute)
}

// CanRoute checks if the junction has any routes available
func (j *Junction) CanRoute() bool {
	return len(j.Routes) > 0
}

// GetRouteCount returns the number of available routes
func (j *Junction) GetRouteCount() int {
	return len(j.Routes)
}

// CreateTJunction creates a T-junction with three routes (excluding the input direction)
func CreateTJunction(pos Position, id rune, inputDir Position) *Junction {
	routes := []Position{}

	// Add all directions except the opposite of input direction
	directions := []Position{Up, Down, Left, Right}
	oppositeDir := Position{-inputDir.X, -inputDir.Y}

	for _, dir := range directions {
		if !dir.Equals(oppositeDir) {
			routes = append(routes, dir)
		}
	}

	return NewJunction(pos, routes, id)
}

// CreateCrossJunction creates a cross junction with four routes
func CreateCrossJunction(pos Position, id rune) *Junction {
	routes := []Position{Up, Down, Left, Right}
	return NewJunction(pos, routes, id)
}
