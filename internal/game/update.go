package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/your-username/packet-rush/internal/types"
)

// handleKeyInput processes keyboard input
func (m *GameModel) handleKeyInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case " ": // Spacebar toggles pause
		m.IsPaused = !m.IsPaused
		if !m.IsPaused {
			// Resume game with tick commands
			return m, tea.Batch(
				tea.Tick(m.GameSpeed, func(t time.Time) tea.Msg {
					return TickMsg{Time: t}
				}),
				tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
					return SpawnMsg{Time: t}
				}),
			)
		}
		return m, nil

	case "r": // Restart game
		if m.IsGameOver {
			newModel := NewGameModel()
			return newModel, newModel.Init()
		}
		return m, nil

	case "d": // Toggle debug mode
		m.DebugMode = !m.DebugMode
		return m, nil
	}

	// Handle junction switching (keys 1-9)
	if len(msg.String()) == 1 {
		key := rune(msg.String()[0])
		if types.JunctionKeys[key] {
			m.switchJunction(key)
			// Enable debug mode temporarily to show junction switched
			if !m.DebugMode {
				m.DebugMode = true
			}
		}
	}

	return m, nil
}

// handleGameTick processes the main game logic updates
func (m *GameModel) handleGameTick(msg TickMsg) (tea.Model, tea.Cmd) {
	if m.IsPaused || m.IsGameOver {
		return m, nil
	}

	// Move all packets
	m.movePackets()

	// Check for collisions and destinations
	m.processPacketCollisions()

	// Remove packets that are off-screen or invalid
	m.cleanupPackets()

	// Check game over conditions
	m.checkGameOver()

	// Increase difficulty over time
	m.increaseDifficulty()

	// Schedule next tick
	return m, tea.Tick(m.GameSpeed, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

// handlePacketSpawn creates new packets at spawn points
func (m *GameModel) handlePacketSpawn(msg SpawnMsg) (tea.Model, tea.Cmd) {
	if m.IsPaused || m.IsGameOver {
		// Still schedule next spawn even if paused
		return m, tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		})
	}

	// Limit the number of packets on screen
	if len(m.Packets) >= types.MaxPacketsOnScreen {
		return m, tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		})
	}

	// Find spawn points and create packets
	spawned := false
	for y, row := range m.LevelMap {
		for x, char := range row {
			if char == types.Spawn {
				spawnPos := types.Position{X: x, Y: y}

				// Create a new packet moving right (default direction)
				packet := types.NewPacket(spawnPos, types.Right)
				m.Packets = append(m.Packets, packet)
				m.TotalPackets++
				spawned = true
				break // Only spawn one packet per tick
			}
		}
		if spawned {
			break
		}
	}

	// Schedule next spawn
	return m, tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
		return SpawnMsg{Time: t}
	})
}

// switchJunction toggles the active route of a junction
func (m *GameModel) switchJunction(key rune) {
	for _, junction := range m.Junctions {
		if junction.ID == key {
			junction.SwitchRoute()
			break
		}
	}
}

// movePackets updates the position of all packets
func (m *GameModel) movePackets() {
	for _, packet := range m.Packets {
		// Store old position for collision detection
		oldPos := packet.Pos

		// Move the packet
		packet.Move()

		// Check if packet moved to a new cell
		if !packet.Pos.Equals(oldPos) {
			m.processPacketAtPosition(packet)
		}
	}
}

// processPacketAtPosition handles packet behavior at its current position
func (m *GameModel) processPacketAtPosition(packet *types.Packet) {
	char := m.getCharAt(packet.Pos)

	switch char {
	case types.JunctionSymbol:
		// Check if there's a junction at this position
		if junction, exists := m.Junctions[packet.Pos]; exists {
			// Change packet direction based on junction's active route
			newDir := junction.GetActiveDirection()
			packet.SetDirection(newDir)
		}

	case types.Wall:
		// Packet hit a wall - mark for removal or bounce back
		packet.SetDirection(types.Position{X: 0, Y: 0}) // Stop the packet

	case types.Destination1, types.Destination2, types.Destination3, types.Destination4:
		// Check if packet reached correct destination
		if packet.IsAtDestination(char) {
			m.Score += types.ScorePerPacket
			m.ComboCount++
			// Packet will be removed in cleanup
		} else {
			// Wrong destination - penalty
			m.Score += types.PenaltyPerMiss
			m.Lives--
			m.ComboCount = 0
		}
	}
}

// processPacketCollisions handles packet interactions and destination checking
func (m *GameModel) processPacketCollisions() {
	for i := len(m.Packets) - 1; i >= 0; i-- {
		packet := m.Packets[i]
		char := m.getCharAt(packet.Pos)

		// Check if packet reached a destination
		switch char {
		case types.Destination1, types.Destination2, types.Destination3, types.Destination4:
			if packet.IsAtDestination(char) {
				// Correct destination - add to score and remove packet
				m.Score += types.ScorePerPacket + (m.ComboCount * types.BonusPerCombo)
				m.ComboCount++
				m.removePacket(i)
			} else {
				// Wrong destination - penalty and remove packet
				m.Score += types.PenaltyPerMiss
				m.Lives--
				m.ComboCount = 0
				m.removePacket(i)
			}
		}
	}
}

// cleanupPackets removes packets that are off-screen or invalid
func (m *GameModel) cleanupPackets() {
	for i := len(m.Packets) - 1; i >= 0; i-- {
		packet := m.Packets[i]

		// Remove packets that are out of bounds
		if !m.isValidPosition(packet.Pos) {
			m.removePacket(i)
			m.Lives-- // Penalty for losing a packet
			m.ComboCount = 0
			continue
		}

		// Remove packets that hit walls
		char := m.getCharAt(packet.Pos)
		if char == types.Wall {
			m.removePacket(i)
			m.Lives-- // Penalty for hitting a wall
			m.ComboCount = 0
		}
	}
}

// removePacket removes a packet from the game at the specified index
func (m *GameModel) removePacket(index int) {
	if index >= 0 && index < len(m.Packets) {
		// Remove packet by swapping with last element and truncating
		m.Packets[index] = m.Packets[len(m.Packets)-1]
		m.Packets = m.Packets[:len(m.Packets)-1]
	}
}

// checkGameOver determines if the game should end
func (m *GameModel) checkGameOver() {
	if m.Lives <= 0 {
		m.IsGameOver = true
		m.GameState = types.StateGameOver
	}
}

// increaseDifficulty adjusts game speed and spawn rate based on score
func (m *GameModel) increaseDifficulty() {
	// Increase difficulty every 100 points
	newLevel := m.Score/100 + 1
	if newLevel > m.Level {
		m.Level = newLevel

		// Increase game speed (make it faster)
		if m.GameSpeed > types.MinGameSpeed {
			m.GameSpeed -= types.SpeedIncreaseRate
		}

		// Increase spawn rate (spawn more frequently)
		if m.SpawnInterval > types.MinSpawnInterval {
			m.SpawnInterval -= 50 * time.Millisecond
		}
	}
}
