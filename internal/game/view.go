package game

import (
	"fmt"
	"strings"

	"github.com/your-username/packet-rush/internal/types"
)

// renderGame renders the main game view
func (m *GameModel) renderGame() string {
	var builder strings.Builder

	// Game title
	builder.WriteString(types.ColorCyan + types.ColorBright + "🚀 PACKET RUSH - Network Router Simulator 🚀" + types.ColorReset + "\n\n")

	// Create a copy of the level map for rendering
	gameGrid := make([][]rune, len(m.LevelMap))
	for i, row := range m.LevelMap {
		gameGrid[i] = []rune(row)
	}

	// Draw junctions with their current state
	for _, junction := range m.Junctions {
		if m.isValidPosition(junction.Pos) {
			gameGrid[junction.Pos.Y][junction.Pos.X] = junction.GetSymbol()
		}
	}

	// Draw packets on the grid
	for _, packet := range m.Packets {
		if m.isValidPosition(packet.Pos) {
			gameGrid[packet.Pos.Y][packet.Pos.X] = packet.GetChar()
		}
	}

	// Render the grid with colors
	for y, row := range gameGrid {
		for x, char := range row {
			pos := types.Position{X: x, Y: y}
			coloredChar := m.getColoredChar(char, pos)
			builder.WriteString(coloredChar)
		}
		builder.WriteString("\n")
	}

	// Add game UI below the grid
	builder.WriteString(m.renderUI())

	// Add debug info if enabled
	if m.DebugMode {
		builder.WriteString(m.renderDebugInfo())
	}

	return builder.String()
}

// renderUI renders the game status UI
func (m *GameModel) renderUI() string {
	var builder strings.Builder

	// Separator line
	builder.WriteString(strings.Repeat("=", 80) + "\n")

	// Game stats
	builder.WriteString(fmt.Sprintf("Score: %d | Lives: %d | Level: %d | Packets: %d",
		m.Score, m.Lives, m.Level, len(m.Packets)))

	if m.ComboCount > 0 {
		builder.WriteString(fmt.Sprintf(" | Combo: %dx", m.ComboCount))
	}
	builder.WriteString("\n")

	// Junction controls
	builder.WriteString("Junction Controls: ")
	count := 0
	for _, junction := range m.Junctions {
		if count > 0 {
			builder.WriteString(" | ")
		}
		builder.WriteString(fmt.Sprintf("%c: %c", junction.ID, junction.GetSymbol()))
		count++
		if count >= 5 { // Limit display to prevent overflow
			break
		}
	}
	builder.WriteString("\n")

	// How to play instructions
	if len(m.Packets) == 0 && m.TotalPackets == 0 {
		builder.WriteString(types.ColorYellow + "📖 HOW TO PLAY: Packets (D,A,V,E) spawn from 'S' and need to reach ports 1-4.\n")
		builder.WriteString("Use keys 1-9 to switch junction directions (+ symbols) and route packets correctly!" + types.ColorReset + "\n")
	}

	// Controls help
	builder.WriteString("Controls: [SPACE] Pause | [Q] Quit | [1-9] Switch Junctions | [D] Debug")
	if m.IsGameOver {
		builder.WriteString(" | [R] Restart")
	}
	builder.WriteString("\n")

	return builder.String()
}

// renderGameOver renders the game over screen
func (m *GameModel) renderGameOver() string {
	var builder strings.Builder

	// Center the game over message
	width := 80
	height := 24

	for i := 0; i < height/2-3; i++ {
		builder.WriteString("\n")
	}

	// Game Over title
	title := "KERNEL PANIC!"
	padding := (width - len(title)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(types.ColorRed + types.ColorBright + title + types.ColorReset + "\n\n")

	// Final score
	scoreText := fmt.Sprintf("Final Score: %d", m.Score)
	padding = (width - len(scoreText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(types.ColorYellow + scoreText + types.ColorReset + "\n\n")

	// Level reached
	levelText := fmt.Sprintf("Level Reached: %d", m.Level)
	padding = (width - len(levelText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(types.ColorCyan + levelText + types.ColorReset + "\n\n")

	// Packets processed
	packetsText := fmt.Sprintf("Packets Processed: %d", m.TotalPackets)
	padding = (width - len(packetsText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(types.ColorGreen + packetsText + types.ColorReset + "\n\n")

	// Controls
	controlsText := "Press [R] to restart or [Q] to quit"
	padding = (width - len(controlsText)) / 2
	builder.WriteString(strings.Repeat(" ", padding))
	builder.WriteString(types.ColorWhite + controlsText + types.ColorReset + "\n")

	return builder.String()
}

// renderPaused renders the pause screen
func (m *GameModel) renderPaused() string {
	var builder strings.Builder

	// Show the current game state but dimmed
	gameView := m.renderGame()

	// Add pause overlay
	lines := strings.Split(gameView, "\n")
	height := len(lines)

	for i, line := range lines {
		if i == height/2 {
			// Center pause message
			pauseMsg := "=== PAUSED ==="
			width := len(line)
			if width > 0 {
				padding := (width - len(pauseMsg)) / 2
				if padding > 0 {
					line = line[:padding] + types.ColorYellow + types.ColorBright + pauseMsg + types.ColorReset + line[padding+len(pauseMsg):]
				}
			}
		}
		builder.WriteString(line + "\n")
	}

	return builder.String()
}

// renderDebugInfo renders debug information
func (m *GameModel) renderDebugInfo() string {
	var builder strings.Builder

	builder.WriteString("\n" + strings.Repeat("-", 80) + "\n")
	builder.WriteString("DEBUG INFO:\n")
	builder.WriteString(fmt.Sprintf("Game Speed: %v | Spawn Interval: %v\n", m.GameSpeed, m.SpawnInterval))
	builder.WriteString(fmt.Sprintf("Active Packets: %d | Total Spawned: %d\n", len(m.Packets), m.TotalPackets))
	builder.WriteString(fmt.Sprintf("Junctions: %d | Combo Count: %d\n", len(m.Junctions), m.ComboCount))

	// Show junction states
	builder.WriteString("Junction States: ")
	for _, junction := range m.Junctions {
		builder.WriteString(fmt.Sprintf("%c:%d/%d ", junction.ID, junction.ActiveRoute, len(junction.Routes)))
	}
	builder.WriteString("\n")

	// Show packet details
	if len(m.Packets) > 0 {
		builder.WriteString("Packets: ")
		for i, packet := range m.Packets {
			if i < 5 { // Show only first 5 packets
				builder.WriteString(fmt.Sprintf("%c@%s ", packet.GetChar(), packet.Pos))
			}
		}
		if len(m.Packets) > 5 {
			builder.WriteString(fmt.Sprintf("... (+%d more)", len(m.Packets)-5))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// getColoredChar returns a colored version of a character based on its type
func (m *GameModel) getColoredChar(char rune, pos types.Position) string {
	// Check if this position has a packet
	for _, packet := range m.Packets {
		if packet.Pos.Equals(pos) {
			return m.getPacketColor(packet) + string(char) + types.ColorReset
		}
	}

	// Color based on character type
	switch char {
	case types.Wall:
		return types.ColorBlue + string(char) + types.ColorReset
	case types.Spawn:
		return types.ColorGreen + types.ColorBright + string(char) + types.ColorReset
	case types.Destination1, types.Destination2, types.Destination3, types.Destination4:
		return types.ColorYellow + types.ColorBright + string(char) + types.ColorReset
	case types.Pipe, types.VerticalPipe:
		return types.ColorWhite + string(char) + types.ColorReset
	case '^', 'v', '<', '>', '+':
		// Junction symbols
		return types.ColorMagenta + types.ColorBright + string(char) + types.ColorReset
	default:
		return string(char)
	}
}

// getPacketColor returns the appropriate color for a packet
func (m *GameModel) getPacketColor(packet *types.Packet) string {
	switch packet.Type {
	case types.DataPacket:
		return types.ColorCyan + types.ColorBright
	case types.AudioPacket:
		return types.ColorGreen + types.ColorBright
	case types.VideoPacket:
		return types.ColorYellow + types.ColorBright
	case types.EmailPacket:
		return types.ColorMagenta + types.ColorBright
	default:
		return types.ColorWhite
	}
}
