package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/your-username/packet-rush/internal/types"
)

// GameModel holds the entire state of the game
// This implements the tea.Model interface for Bubble Tea
type GameModel struct {
	// Game state
	LevelMap  []string                           // The static grid layout
	Packets   []*types.Packet                    // Active packets on the grid
	Junctions map[types.Position]*types.Junction // Map for quick lookup

	// Game status
	Score      int
	Lives      int
	Level      int
	IsGameOver bool
	IsPaused   bool
	GameState  types.GameState

	// Timing
	GameSpeed     time.Duration
	SpawnInterval time.Duration
	LastSpawn     time.Time

	// Grid dimensions
	Width  int
	Height int

	// Game mechanics
	Difficulty   types.Difficulty
	ComboCount   int
	TotalPackets int

	// Debug info
	DebugMode bool
}

// NewGameModel creates a new game model with default values
func NewGameModel() *GameModel {
	return &GameModel{
		LevelMap:      getDefaultLevel(),
		Packets:       make([]*types.Packet, 0),
		Junctions:     make(map[types.Position]*types.Junction),
		Score:         0,
		Lives:         types.MaxLives,
		Level:         1,
		IsGameOver:    false,
		IsPaused:      false,
		GameState:     types.StatePlaying,
		GameSpeed:     types.DefaultGameSpeed,
		SpawnInterval: types.PacketSpawnInterval,
		LastSpawn:     time.Now(),
		Width:         types.DefaultGridWidth,
		Height:        types.DefaultGridHeight,
		Difficulty:    types.DifficultyEasy,
		ComboCount:    0,
		TotalPackets:  0,
		DebugMode:     false,
	}
}

// Init initializes the game model (required by tea.Model interface)
func (m *GameModel) Init() tea.Cmd {
	// Initialize junctions from the level map
	m.initializeJunctions()

	// Return a command to start the game tick
	return tea.Batch(
		tea.Tick(m.GameSpeed, func(t time.Time) tea.Msg {
			return TickMsg{Time: t}
		}),
		tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		}),
	)
}

// Update handles all game events and state transitions (required by tea.Model interface)
func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyInput(msg)

	case TickMsg:
		return m.handleGameTick(msg)

	case SpawnMsg:
		return m.handlePacketSpawn(msg)

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	}

	return m, nil
}

// View renders the game state (required by tea.Model interface)
func (m *GameModel) View() string {
	if m.GameState == types.StateGameOver || m.IsGameOver {
		return m.renderGameOver()
	}

	if m.IsPaused {
		return m.renderPaused()
	}

	return m.renderGame()
}

// initializeJunctions scans the level map and creates junction objects
func (m *GameModel) initializeJunctions() {
	junctionID := '1'

	for y, row := range m.LevelMap {
		for x, char := range row {
			if char == types.JunctionSymbol {
				pos := types.Position{X: x, Y: y}

				// Determine available routes by checking adjacent cells
				routes := m.getAvailableRoutes(pos)

				// Create junction with available routes
				junction := types.NewJunction(pos, routes, junctionID)
				m.Junctions[pos] = junction

				// Increment junction ID
				if junctionID < '9' {
					junctionID++
				}
			}
		}
	}
}

// getAvailableRoutes determines valid exit directions from a junction
func (m *GameModel) getAvailableRoutes(pos types.Position) []types.Position {
	routes := []types.Position{}
	directions := []types.Position{types.Up, types.Down, types.Left, types.Right}

	for _, dir := range directions {
		newPos := pos.Add(dir)
		if m.isValidPosition(newPos) {
			char := m.getCharAt(newPos)
			// Check if this direction leads to a valid path
			if char != types.Wall && char != types.EmptySpace {
				routes = append(routes, dir)
			}
		}
	}

	return routes
}

// isValidPosition checks if a position is within the grid bounds
func (m *GameModel) isValidPosition(pos types.Position) bool {
	return pos.Y >= 0 && pos.Y < len(m.LevelMap) &&
		pos.X >= 0 && pos.X < len(m.LevelMap[pos.Y])
}

// getCharAt returns the character at a specific position
func (m *GameModel) getCharAt(pos types.Position) rune {
	if !m.isValidPosition(pos) {
		return types.Wall // Treat out-of-bounds as walls
	}
	return rune(m.LevelMap[pos.Y][pos.X])
}

// getDefaultLevel returns a simple test level
func getDefaultLevel() []string {
	return []string{
		"################################################################################",
		"#S----+----+----1                                                             #",
		"#     |    |                                                                  #",
		"#     |    +----2                                                             #",
		"#     |                                                                       #",
		"#     +----+----3                                                             #",
		"#          |                                                                  #",
		"#          +----4                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"#                                                                             #",
		"################################################################################",
	}
}

// Message types for Bubble Tea
type TickMsg struct {
	Time time.Time
}

type SpawnMsg struct {
	Time time.Time
}
