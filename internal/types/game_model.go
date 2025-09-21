package types

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
	Grid      []string
	Packets   []*Packet
	Junctions map[string]*Junction

	Score         int
	Lives         int
	Level         int
	GameTime      int
	LastSpawn     time.Time
	Paused        bool
	GameOver      bool
	LevelComplete bool

	SpawnInterval time.Duration
	TickSpeed     time.Duration

	CurrentGoal  string
	GoalProgress []rune
	TargetWord   string

	RestartRequested   bool
	NextLevelRequested bool
}

type TickMsg struct {
	Time time.Time
}

type SpawnMsg struct {
	Time time.Time
}

func NewGameModel() *GameModel {
	return &GameModel{
		Level:     1,
		Lives:     LivesPerLevel,
		TickSpeed: InitialTickSpeed,
	}
}

func (m *GameModel) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(m.TickSpeed, func(t time.Time) tea.Msg {
			return TickMsg{Time: t}
		}),
		tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		}),
	)
}

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyInput(msg)

	case TickMsg:
		return m.handleGameTick(msg)

	case SpawnMsg:
		return m.handlePacketSpawn(msg)

	case tea.WindowSizeMsg:
		return m, nil
	}

	return m, nil
}

func (m *GameModel) IsValidPosition(x, y int) bool {
	return y >= 0 && y < len(m.Grid) && x >= 0 && x < len(m.Grid[y])
}

func (m *GameModel) GetCharAt(x, y int) rune {
	if !m.IsValidPosition(x, y) {
		return '#'
	}
	return rune(m.Grid[y][x])
}
