package types

import (
	"fmt"
	"strings"
)

func (m *GameModel) View() string {
	
	if m.LevelComplete {
		return m.renderLevelComplete()
	}

	if m.GameOver {
		return m.renderGameOver()
	}

	return m.renderGame()
}


func (m *GameModel) renderGame() string {
	var builder strings.Builder

	
	builder.WriteString(ColorCyan + ColorBright + "🚀 PACKET RUSH - Network Router Simulator 🚀" + ColorReset + "\n\n")

	
	display := make([][]rune, len(m.Grid))
	for i, row := range m.Grid {
		display[i] = []rune(row)
	}

	
	for _, junction := range m.Junctions {
		if m.IsValidPosition(junction.X, junction.Y) {
			display[junction.Y][junction.X] = junction.GetSymbol()
		}
	}

	
	for _, packet := range m.Packets {
		if m.IsValidPosition(packet.X, packet.Y) {
			display[packet.Y][packet.X] = packet.GetChar()
		}
	}

	
	for y, row := range display {
		for x, char := range row {
			coloredChar := getColoredChar(char, x, y, m)
			builder.WriteString(coloredChar)
		}
		builder.WriteString("\n")
	}

	
	builder.WriteString("╔══════════════════════════════════════════════════════════════════════════════╗\n")
	builder.WriteString(fmt.Sprintf("║ "+ColorYellow+"Level: %d"+ColorReset+" │ "+
		ColorYellow+"Score: %d"+ColorReset+" │ "+
		ColorRed+"Lives: %d"+ColorReset+" │ "+
		ColorGreen+"Packets: %d"+ColorReset+" │ "+
		ColorCyan+"Time: %ds"+ColorReset+" ║\n",
		m.Level, m.Score, m.Lives, len(m.Packets), m.GameTime/5))
	builder.WriteString("╚══════════════════════════════════════════════════════════════════════════════╝\n")

	
	builder.WriteString(ColorMagenta + "🎯 Goal: " + ColorReset + m.CurrentGoal + "\n")
	builder.WriteString(ColorCyan + "📝 Progress: " + ColorReset)
	for _, letter := range m.GoalProgress {
		builder.WriteString(ColorGreen + ColorBright + string(letter) + ColorReset + " ")
	}
	needed := len(m.TargetWord) - len(m.GoalProgress)
	if needed > 0 {
		builder.WriteString(ColorWhite + fmt.Sprintf("(need %d more: %s)", needed, m.TargetWord) + ColorReset)
	}
	builder.WriteString("\n")

	
	builder.WriteString(ColorMagenta + "Junctions: " + ColorReset)
	for _, junction := range m.Junctions {
		dir := junction.GetActiveDirection()
		symbol := "?"
		color := ColorWhite
		switch dir {
		case Right:
			symbol = "→"
			color = ColorGreen
		case Down:
			symbol = "↓"
			color = ColorBlue
		case Up:
			symbol = "↑"
			color = ColorRed
		case Left:
			symbol = "←"
			color = ColorYellow
		}
		builder.WriteString(fmt.Sprintf(ColorBright+"%c"+ColorReset+":"+color+"%s"+ColorReset+" ", junction.ID, symbol))
	}
	builder.WriteString("\n")

	
	if len(m.Packets) > 0 {
		builder.WriteString(ColorYellow + "Active Packets: " + ColorReset)
		for i, packet := range m.Packets {
			if i < 5 { 
				color := packet.GetColor()
				builder.WriteString(fmt.Sprintf(color+"%c"+ColorReset+" ", packet.PacketType))
			}
		}
		if len(m.Packets) > 5 {
			builder.WriteString(fmt.Sprintf("+"+ColorWhite+"%d more..."+ColorReset, len(m.Packets)-5))
		}
		builder.WriteString("\n")
	}

	if m.Paused {
		builder.WriteString(BgYellow + ColorBlack + " ⏸️  PAUSED - Press SPACE to continue " + ColorReset + "\n")
	}

		
	builder.WriteString(ColorWhite + "Controls: " + ColorGreen + "[1-9]" + ColorWhite + " Switch Junctions │ " +
		ColorYellow + "[SPACE]" + ColorWhite + " Pause │ " +
		ColorRed + "[Q]" + ColorWhite + " Quit" + ColorReset + "\n")

	return builder.String()
}


func (m *GameModel) renderLevelComplete() string {
	var builder strings.Builder

	
	gameView := m.renderGame()

	
	lines := strings.Split(gameView, "\n")
	height := len(lines)

	for i, line := range lines {
		switch i {
case height/2 - 2:
			if m.Level < MaxLevel {
				msg := "🎉 LEVEL COMPLETE! 🎉"
				width := 80
				padding := (width - len(msg)) / 2
				if padding > 0 {
					line = strings.Repeat(" ", padding) + BgGreen + ColorWhite + ColorBright + msg + ColorReset + strings.Repeat(" ", padding)
				}
			} else {
				msg := "🏆 ALL LEVELS COMPLETE! YOU'RE A NETWORK MASTER! 🏆"
				width := 80
				padding := (width - len(msg)) / 2
				if padding > 0 {
					line = strings.Repeat(" ", padding) + BgGold + ColorBlack + ColorBright + msg + ColorReset + strings.Repeat(" ", padding)
				}
			}
		case height / 2:
			msg := "Press R for next level"
			if m.Level >= MaxLevel {
				msg = "Press R to restart"
			}
			width := 80
			padding := (width - len(msg)) / 2
			if padding > 0 {
				line = strings.Repeat(" ", padding) + ColorYellow + ColorBright + msg + ColorReset + strings.Repeat(" ", padding)
			}
		}
		builder.WriteString(line + "\n")
	}

	return builder.String()
}


func (m *GameModel) renderGameOver() string {
	var builder strings.Builder

	width := 80
	height := 24

	for i := 0; i < height/2-3; i++ {
		builder.WriteString("\n")
	}

	title := "KERNEL PANIC!"
	padding := (width - len(title)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(ColorRed + ColorBright + title + ColorReset + "\n\n")

	scoreText := fmt.Sprintf("Final Score: %d", m.Score)
	padding = (width - len(scoreText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(ColorYellow + scoreText + ColorReset + "\n\n")

	levelText := fmt.Sprintf("Level Reached: %d", m.Level)
	padding = (width - len(levelText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(ColorCyan + levelText + ColorReset + "\n\n")

	controlsText := "Press [R] to restart or [Q] to quit"
	padding = (width - len(controlsText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(ColorWhite + controlsText + ColorReset + "\n")

	return builder.String()
}

func getColoredChar(char rune, x, y int, m *GameModel) string {
	for _, packet := range m.Packets {
		if packet.X == x && packet.Y == y {
			color := packet.GetColor()
			return color + ColorBright + string(char) + ColorReset
		}
	}

	
	for _, junction := range m.Junctions {
		if junction.X == x && junction.Y == y {
			return ColorMagenta + ColorBright + string(char) + ColorReset
		}
	}


	switch char {
	case '#': 
		return ColorBlue + string(char) + ColorReset
	case 'S': 
		return BgGreen + ColorWhite + string(char) + ColorReset
	case '1', '2', '3', '4': 
		return BgYellow + ColorBlue + ColorBright + string(char) + ColorReset
	case '-', '|': 
		return ColorWhite + string(char) + ColorReset
	case '+': 
		return ColorMagenta + ColorBright + string(char) + ColorReset
	default:
		
		if IsLetterDestination(char) {
			return BgMagenta + ColorWhite + ColorBright + string(char) + ColorReset
		}
		return string(char)
	}
}
