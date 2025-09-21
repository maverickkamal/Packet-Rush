package types

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *GameModel) handleKeyInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	switch key {
	case "ctrl+c", "q":
		return m, tea.Quit

	case " ":
		m.Paused = !m.Paused
		if !m.Paused && !m.GameOver && !m.LevelComplete {
			return m, tea.Batch(
				tea.Tick(m.TickSpeed, func(t time.Time) tea.Msg {
					return TickMsg{Time: t}
				}),
				tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
					return SpawnMsg{Time: t}
				}),
			)
		}
		return m, nil

	case "r":
		if m.GameOver {
			m.RestartRequested = true
			return m, nil
		}
		if m.LevelComplete {
			m.NextLevelRequested = true
			return m, nil
		}
		return m, nil
	}

	if len(key) == 1 {
		keyRune := rune(key[0])
		if keyRune >= '1' && keyRune <= '9' {
			m.switchJunction(keyRune)
		}
	}

	return m, nil
}

func (m *GameModel) handleGameTick(_ TickMsg) (tea.Model, tea.Cmd) {
	// Don't process ticks if paused, game over, or level complete
	if m.Paused || m.GameOver || m.LevelComplete {
		return m, nil
	}

	m.movePackets()

	m.processPacketCollisions()

	m.cleanupPackets()

	m.GameTime++

	// Check game over conditions (after all processing)
	m.checkGameOver()

	m.TickSpeed = InitialTickSpeed - time.Duration(m.GameTime)*time.Millisecond
	if m.TickSpeed < MinTickSpeed {
		m.TickSpeed = MinTickSpeed
	}

	// Dynamic tick speed - starts slow, gets faster
	if !m.GameOver && !m.LevelComplete {
		return m, tea.Tick(m.TickSpeed, func(t time.Time) tea.Msg {
			return TickMsg{Time: t}
		})
	}

	return m, nil
}

func (m *GameModel) handlePacketSpawn(_ SpawnMsg) (tea.Model, tea.Cmd) {
	if m.Paused || m.GameOver || m.LevelComplete {
		return m, nil
	}

	if len(m.Packets) >= 10 {
		return m, tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		})
	}

	if time.Since(m.LastSpawn) > m.SpawnInterval {
		newPacket := NewRandomPacket(1, 4, m.TargetWord)
		m.Packets = append(m.Packets, newPacket)
		m.LastSpawn = time.Now()

		if m.SpawnInterval > MinSpawnInterval {
			m.SpawnInterval -= 50 * time.Millisecond
		}
	}

	if !m.GameOver && !m.LevelComplete {
		return m, tea.Tick(m.SpawnInterval, func(t time.Time) tea.Msg {
			return SpawnMsg{Time: t}
		})
	}

	return m, nil
}

func (m *GameModel) switchJunction(key rune) {
	for _, junction := range m.Junctions {
		if junction.ID == key {
			junction.SwitchRoute()
			break
		}
	}
}

func (m *GameModel) movePackets() {
	for _, packet := range m.Packets {
		oldX, oldY := packet.X, packet.Y

		packet.Move()

		if packet.X != oldX || packet.Y != oldY {
			m.processPacketAtPosition(packet)
		}
	}
}

func (m *GameModel) processPacketAtPosition(packet *Packet) {
	char := m.GetCharAt(packet.X, packet.Y)

	jKey := fmt.Sprintf("%d,%d", packet.X, packet.Y)
	if junction, exists := m.Junctions[jKey]; exists {
		dir := junction.GetActiveDirection()
		packet.SetDirection(dir.X, dir.Y)
	}

	if char == '#' {
		packet.SetDirection(0, 0) // Stop the packet
	}
}

func (m *GameModel) processPacketCollisions() {
	for i := len(m.Packets) - 1; i >= 0; i-- {
		packet := m.Packets[i]
		char := m.GetCharAt(packet.X, packet.Y)

		if char == packet.PacketType {
			m.Score += CorrectLetterPoints
			m.GoalProgress = append(m.GoalProgress, packet.PacketType)

			if len(m.GoalProgress) >= len(m.TargetWord) {
				m.LevelComplete = true
				m.Score += LevelCompleteBonus
				return
			}

			m.removePacket(i)
			continue
		} else if IsLetterDestination(char) && char != packet.PacketType {
			m.Lives--
			m.removePacket(i)
			continue
		}
	}
}

func (m *GameModel) cleanupPackets() {
	for i := len(m.Packets) - 1; i >= 0; i-- {
		packet := m.Packets[i]

		if !m.IsValidPosition(packet.X, packet.Y) || m.GetCharAt(packet.X, packet.Y) == '#' {
			m.removePacket(i)
			m.Lives--
		}
	}
}

func (m *GameModel) removePacket(index int) {
	if index >= 0 && index < len(m.Packets) {
		m.Packets[index] = m.Packets[len(m.Packets)-1]
		m.Packets = m.Packets[:len(m.Packets)-1]
	}
}

func (m *GameModel) checkGameOver() {
	if m.Lives <= 0 && !m.LevelComplete {
		m.GameOver = true
	}
}
