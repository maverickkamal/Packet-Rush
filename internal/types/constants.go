package types

import "time"

// Game constants and configuration
const (
	// Grid dimensions
	DefaultGridWidth  = 80
	DefaultGridHeight = 24

	// Game timing
	DefaultGameSpeed    = 500 * time.Millisecond
	MinGameSpeed        = 50 * time.Millisecond
	SpeedIncreaseRate   = 10 * time.Millisecond
	PacketSpawnInterval = 2000 * time.Millisecond
	MinSpawnInterval    = 300 * time.Millisecond

	// Scoring
	ScorePerPacket = 10
	BonusPerCombo  = 5
	PenaltyPerMiss = -5

	// Game limits
	MaxLives           = 3
	MaxPacketsOnScreen = 50

	// UI constants
	InfoPanelHeight = 4
	GameAreaHeight  = DefaultGridHeight - InfoPanelHeight
)

// Map symbols
const (
	EmptySpace     = ' '
	Wall           = '#'
	Spawn          = 'S'
	Destination1   = '1'
	Destination2   = '2'
	Destination3   = '3'
	Destination4   = '4'
	Pipe           = '-'
	VerticalPipe   = '|'
	JunctionSymbol = '+'
)

// Key mappings for junction control
var JunctionKeys = map[rune]bool{
	'1': true, '2': true, '3': true, '4': true, '5': true,
	'6': true, '7': true, '8': true, '9': true, '0': true,
}

// Color definitions
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorBright  = "\033[1m"
)

// Game states
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateHighScore
)

// Difficulty levels
type Difficulty int

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
	DifficultyInsane
)

// GetSpawnIntervalForDifficulty returns the packet spawn interval for a given difficulty
func GetSpawnIntervalForDifficulty(diff Difficulty) time.Duration {
	switch diff {
	case DifficultyEasy:
		return 1500 * time.Millisecond
	case DifficultyMedium:
		return 1000 * time.Millisecond
	case DifficultyHard:
		return 700 * time.Millisecond
	case DifficultyInsane:
		return 400 * time.Millisecond
	default:
		return PacketSpawnInterval
	}
}

// GetGameSpeedForDifficulty returns the game speed for a given difficulty
func GetGameSpeedForDifficulty(diff Difficulty) time.Duration {
	switch diff {
	case DifficultyEasy:
		return 300 * time.Millisecond
	case DifficultyMedium:
		return 200 * time.Millisecond
	case DifficultyHard:
		return 150 * time.Millisecond
	case DifficultyInsane:
		return 100 * time.Millisecond
	default:
		return DefaultGameSpeed
	}
}
