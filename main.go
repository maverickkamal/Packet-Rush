package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Enhanced game with proper junctions and scoring
type GameModel struct {
	grid          []string
	packets       []Packet
	junctions     map[string]*Junction
	score         int
	lives         int
	gameTime      int
	lastSpawn     time.Time
	paused        bool
	gameOver      bool
	spawnInterval time.Duration
}

type Packet struct {
	x, y       int
	packetType rune // D, A, V, E
	dest       int  // 1-4
	dirX, dirY int
}

type Junction struct {
	x, y       int
	directions []Direction
	activeDir  int
	id         rune
}

type Direction struct {
	x, y int
}

var (
	Right = Direction{1, 0}
	Down  = Direction{0, 1}
	Up    = Direction{0, -1}
	Left  = Direction{-1, 0}
)

func NewGameModel() GameModel {
	grid := []string{
		"################################################################################",
		"#S----+----+----1                                                             #",
		"#     |    |                                                                  #",
		"#     |    +----2                                                             #",
		"#     |                                                                       #",
		"#     +----+----3                                                             #",
		"#          |                                                                  #",
		"#          +----4                                                             #",
		"#                                                                             #",
		"#  🚀 PACKET RUSH - Network Router Simulator                                 #",
		"#                                                                             #",
		"#  📖 HOW TO PLAY:                                                            #",
		"#  • Packets (D,A,V,E) spawn from S and need to reach ports 1-4             #",
		"#  • Each packet has a destination number                                     #",
		"#  • Use keys 1-4 to switch junction directions at + symbols                 #",
		"#  • Score points for correct deliveries, lose lives for mistakes!           #",
		"#                                                                             #",
		"#  🎮 CONTROLS: 1-4=Switch Junctions | Space=Pause | Q=Quit | D=Debug       #",
		"#                                                                             #",
		"################################################################################",
	}

	// Initialize junctions
	junctions := make(map[string]*Junction)
	junctions["6,1"] = &Junction{6, 1, []Direction{Right, Down}, 0, '1'}   // First junction
	junctions["11,1"] = &Junction{11, 1, []Direction{Right, Down}, 0, '2'} // Second junction
	junctions["6,5"] = &Junction{6, 5, []Direction{Right, Up}, 0, '3'}     // Third junction
	junctions["11,5"] = &Junction{11, 5, []Direction{Right, Down}, 0, '4'} // Fourth junction

	return GameModel{
		grid:          grid,
		packets:       make([]Packet, 0),
		junctions:     junctions,
		score:         0,
		lives:         3,
		gameTime:      0,
		lastSpawn:     time.Now(),
		paused:        false,
		gameOver:      false,
		spawnInterval: 3 * time.Second,
	}
}

func (m GameModel) Init() tea.Cmd {
	return tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
		return "tick"
	})
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit

		case " ":
			m.paused = !m.paused

		case "r":
			if m.gameOver {
				return NewGameModel(), NewGameModel().Init()
			}

		case "1", "2", "3", "4":
			// Switch junction
			for _, junction := range m.junctions {
				if junction.id == rune(key[0]) {
					junction.activeDir = (junction.activeDir + 1) % len(junction.directions)
					break
				}
			}
		}

	case string:
		if msg == "tick" && !m.paused && !m.gameOver {
			// Spawn new packets
			if time.Since(m.lastSpawn) > m.spawnInterval {
				packetTypes := []rune{'D', 'A', 'V', 'E'}
				newPacket := Packet{
					x:          1,
					y:          1,
					packetType: packetTypes[rand.Intn(len(packetTypes))],
					dest:       rand.Intn(4) + 1,
					dirX:       1,
					dirY:       0,
				}
				m.packets = append(m.packets, newPacket)
				m.lastSpawn = time.Now()

				// Increase difficulty over time
				if m.spawnInterval > 800*time.Millisecond {
					m.spawnInterval -= 100 * time.Millisecond
				}
			}

			// Move packets and handle collisions
			for i := len(m.packets) - 1; i >= 0; i-- {
				packet := &m.packets[i]

				// Move packet
				packet.x += packet.dirX
				packet.y += packet.dirY

				// Check for junctions
				jKey := fmt.Sprintf("%d,%d", packet.x, packet.y)
				if junction, exists := m.junctions[jKey]; exists {
					dir := junction.directions[junction.activeDir]
					packet.dirX = dir.x
					packet.dirY = dir.y
				}

				// Check for destinations
				if packet.y >= 0 && packet.y < len(m.grid) &&
					packet.x >= 0 && packet.x < len(m.grid[packet.y]) {
					char := rune(m.grid[packet.y][packet.x])
					if char >= '1' && char <= '4' {
						destNum := int(char - '0')
						if destNum == packet.dest {
							m.score += 10
						} else {
							m.lives--
							if m.lives <= 0 {
								m.gameOver = true
							}
						}
						// Remove packet
						m.packets = append(m.packets[:i], m.packets[i+1:]...)
						continue
					}
				}

				// Remove packets that go off-screen
				if packet.x > 78 || packet.y > 18 || packet.x < 0 || packet.y < 0 {
					m.packets = append(m.packets[:i], m.packets[i+1:]...)
					m.lives--
					if m.lives <= 0 {
						m.gameOver = true
					}
				}
			}

			m.gameTime++
		}
	}

	return m, tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
		return "tick"
	})
}

func (m GameModel) View() string {
	var builder strings.Builder

	// Create display grid
	display := make([][]rune, len(m.grid))
	for i, row := range m.grid {
		display[i] = []rune(row)
	}

	// Draw junction states
	for _, junction := range m.junctions {
		if junction.y >= 0 && junction.y < len(display) &&
			junction.x >= 0 && junction.x < len(display[junction.y]) {
			dir := junction.directions[junction.activeDir]
			if dir.x == 1 {
				display[junction.y][junction.x] = '>'
			} else if dir.y == 1 {
				display[junction.y][junction.x] = 'v'
			} else if dir.y == -1 {
				display[junction.y][junction.x] = '^'
			} else {
				display[junction.y][junction.x] = '<'
			}
		}
	}

	// Draw packets
	for _, packet := range m.packets {
		if packet.y >= 0 && packet.y < len(display) &&
			packet.x >= 0 && packet.x < len(display[packet.y]) {
			display[packet.y][packet.x] = packet.packetType
		}
	}

	// Render grid
	for _, row := range display {
		builder.WriteString(string(row) + "\n")
	}

	// Game stats
	builder.WriteString(fmt.Sprintf("Score: %d | Lives: %d | Packets: %d | Time: %ds\n",
		m.score, m.lives, len(m.packets), m.gameTime/5))

	// Junction states
	builder.WriteString("Junctions: ")
	for _, junction := range m.junctions {
		dir := junction.directions[junction.activeDir]
		symbol := "?"
		if dir.x == 1 {
			symbol = "→"
		} else if dir.y == 1 {
			symbol = "↓"
		} else if dir.y == -1 {
			symbol = "↑"
		}
		builder.WriteString(fmt.Sprintf("%c:%s ", junction.id, symbol))
	}
	builder.WriteString("\n")

	if m.paused {
		builder.WriteString("⏸️  PAUSED - Press SPACE to continue\n")
	}

	if m.gameOver {
		builder.WriteString("💥 KERNEL PANIC! Game Over - Press R to restart\n")
	}

	return builder.String()
}

func main() {
	fmt.Println("🚀 Starting Packet Rush...")

	model := NewGameModel()
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Thanks for playing Packet Rush!")
}
