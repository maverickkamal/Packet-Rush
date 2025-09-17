# 🚀 Packet Rush

A fast-paced terminal-based puzzle game where you play as a network router managing packet flow!

## 🎮 Game Concept

Direct chaotic data packets to their correct destination ports by switching network junctions. As the game progresses, speed and complexity increase until your network experiences a "kernel panic" (game over).

## 🏗️ Implementation Status

### ✅ Phase 1: Foundation (COMPLETED)
- **Core Data Structures**: Position, Packet, Junction, GameModel
- **Project Structure**: Clean modular architecture with `internal/` organization
- **Basic Types**: PacketType, GameState, Difficulty levels
- **Constants & Configuration**: Game timing, scoring, grid dimensions
- **Bubble Tea Integration**: Base model implementing `tea.Model` interface

### 🚧 Phase 2: Game Engine (IN PROGRESS)
- **Game Loop**: Tick-based updates and packet movement
- **Rendering System**: Terminal UI with colors and animations  
- **Input Handling**: Junction switching via keyboard (1-9)
- **Game Logic**: Collision detection, scoring, lives system

### 📋 Phase 3: Polish (PENDING)
- **Advanced Features**: Multiple levels, difficulty scaling
- **Visual Effects**: Smooth animations, color-coded packets
- **Game Modes**: Different difficulty levels
- **Persistence**: High score tracking

## 🚀 Current Features

- **Modular Architecture**: Well-organized Go code with clear separation of concerns
- **Terminal-Native**: Designed specifically for terminal environments
- **Concurrent Design**: Uses goroutines for smooth gameplay
- **Extensible**: Easy to add new levels, packet types, and features

## 🎯 Game Mechanics

- **Packets**: Different types (Data, Audio, Video, Email) with unique destinations
- **Junctions**: Switchable routing points controlled by keyboard (keys 1-9)
- **Scoring**: Points for correct routing, penalties for mistakes
- **Lives System**: Limited chances before "kernel panic"
- **Progressive Difficulty**: Speed increases as score grows

## 🔧 Building & Running

```bash
# Build the game
go build -o packet-rush.exe .

# Run the game
./packet-rush.exe
```

## 🎮 Controls (When Implemented)

- **1-9**: Switch junction routes
- **SPACE**: Pause/Resume
- **Q**: Quit game
- **R**: Restart (when game over)
- **D**: Toggle debug mode

## 📁 Project Structure

```
packet-rush/
├── main.go                 # Entry point
├── go.mod                  # Go module
├── palms.json             # Competition metadata
├── internal/              
│   ├── game/              # Core game logic
│   ├── types/             # Type definitions  
│   ├── levels/            # Level management
│   └── ui/                # UI components
└── assets/                # Game assets
    └── levels/            # Level files
```

## 🎉 Why This Game Rocks

- **Original Concept**: Network router simulation is perfect for terminal gaming
- **Performance**: Go's speed ensures smooth 60fps gameplay
- **Accessibility**: Runs on any terminal without dependencies
- **Addictive**: Simple rules, complex strategy, high replayability

---

*Built with ❤️ using Go and Bubble Tea for TerminalCraft Competition*

