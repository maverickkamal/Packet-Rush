package types

import "time"

const (
	ColorReset   = "\033[0m"
	ColorBlack   = "\033[30m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorBright  = "\033[1m"

	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgGold    = "\033[43m" 
)

const (
	InitialTickSpeed     = 800 * time.Millisecond
	MinTickSpeed         = 150 * time.Millisecond
	InitialSpawnInterval = 4 * time.Second
	MinSpawnInterval     = 1 * time.Second
)

const (
	CorrectLetterPoints = 20
	LevelCompleteBonus  = 100
	LivesPerLevel       = 2
	MaxLevel            = 10
)

