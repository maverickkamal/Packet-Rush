# 🚀 Packet Rush - Network Router Simulator

A thrilling terminal-based puzzle game where you manage network packet routing through an increasingly complex network infrastructure!

[![Go Version](https://img.shields.io/badge/go-1.25.1-blue.svg)](https://golang.org)
[![Built with Bubble Tea](https://img.shields.io/badge/built%20with-Bubble%20Tea-purple.svg)](https://github.com/charmbracelet/bubbletea)
[![Competition](https://img.shields.io/badge/TerminalCraft-2024-gold.svg)](https://terminalcraft.com)

## 🎮 Game Overview

Take control of a network router and guide data packets to their destinations by switching junction routes in real-time. Each level challenges you to spell words by routing letter packets through an increasingly complex network maze!

### ✨ Key Features

- **🌈 Beautiful Colored Interface** - Immersive terminal UI with full color support
- **📈 Progressive Difficulty** - 10 challenging levels from simple to master-level complexity  
- **🎯 Word-Spelling Objectives** - Route packets to spell words like "GO", "HI", "WIN", "CODE", "RUSH"
- **⚡ Dynamic Speed** - Starts slow and gradually increases for mounting tension
- **🎛️ Interactive Controls** - Switch up to 9 junctions with responsive keyboard controls
- **🏆 Scoring System** - Earn points for correct routing, manage limited lives
- **⏸️ Pause & Resume** - Strategic planning with built-in pause functionality

## 🚀 Quick Start

### Prerequisites
- Go 1.25.1 or higher
- Terminal with color support

### Installation & Run
```bash
# Clone and navigate to project
cd packet-rush

# Build the game
go build -o packet-rush.exe .

# Run and enjoy!
./packet-rush.exe
```

## 🎯 Game Mechanics

### Core Gameplay
- **Letter Packets**: Route packets (G, O, H, I, W, etc.) to matching letter destinations
- **Junction Control**: Switch network junctions using keys 1-9 to direct packet flow
- **Word Completion**: Successfully route all letters to spell the target word
- **Level Progression**: Complete levels to unlock increasingly complex networks

### Difficulty Progression
- **Level 1**: Simple "GO" - Learn the basics with 2 letters
- **Level 2**: Extended "HI" - Longer network paths  
- **Level 3**: Complex "WIN" - Multi-junction routing
- **Level 4**: Advanced "CODE" - Master-level complexity
- **Level 5**: Ultimate "RUSH" - Maximum challenge
- **Levels 6-10**: Expert words like "EXPERT", "GENIUS", "MASTER", "LEGEND", "CHAMPION"

### Lives & Scoring
- **2 Lives per Level** - Limited chances to complete each challenge
- **20 Points** - For each correctly routed packet
- **100 Bonus Points** - For completing a level
- **Progressive Speed** - Packet movement and spawning accelerate over time

## 🎮 Controls

| Key | Action |
|-----|--------|
| **1-9** | Switch junction directions (watch the arrows!) |
| **SPACE** | Pause/Resume game |
| **R** | Restart current level or advance to next level |
| **Q** | Quit game |

## 🗺️ Visual Guide

```
🚀 PACKET RUSH - Network Router Simulator 🚀

################################################################################
#  🚀 LEVEL 1: NETWORK BASICS - Spell 'GO' with packets!                     #
#                                                                              #
#S--------G-->--------O->---------O>----------G                              #
#         |          |          |                                            #
#         |          |          |                                            #
#         |          |          O                                            #
#         >---------->---------->----------1                                  #
################################################################################
║ Level: 1 │ Score: 140 │ Lives: 2 │ Packets: 3 │ Time: 17s ║
🎯 Goal: Route G packets to 'G' and O packets to 'O' to spell 'GO'!
📝 Progress: G (need 1 more: GO)
Junctions: 1:→ 2:→ 3:→ 4:→ 5:→ 6:→
Active Packets: G O G
```

### Legend
- **S** = Spawn point where packets appear
- **G, O, H, I, etc.** = Letter destination ports
- **→ ↓ ↑ ←** = Junction direction indicators
- **1-9** = Junction control numbers
- **G, O, etc.** = Letter packets moving through network

## 🏗️ Technical Architecture

### Clean Modular Design
```
packet-rush/
├── main.go                     # Clean entry point (26 lines)
├── internal/
│   ├── types/                  # Core game logic & rendering
│   │   ├── game_model.go       # Game state management
│   │   ├── game_methods.go     # Game logic & controls
│   │   ├── view_methods.go     # UI rendering & colors
│   │   ├── packet.go           # Packet behavior
│   │   ├── junction.go         # Junction switching logic
│   │   └── constants.go        # Game constants & colors
│   ├── game/
│   │   └── coordinator.go      # Level transition coordinator
│   └── levels/
│       ├── level_data.go       # All 10 level definitions
│       └── game_factory.go     # Level initialization
├── go.mod                      # Go module definition
├── GAMEPLAY_GUIDE.md          # Detailed gameplay instructions
└── README.md                   # This file
```

### Built With
- **Go 1.25.1** - High-performance, concurrent game engine
- **Bubble Tea** - Elegant terminal UI framework
- **Clean Architecture** - Modular, maintainable codebase

## 📚 Need Help?

For detailed gameplay instructions, strategies, and pro tips, check out our comprehensive **[GAMEPLAY_GUIDE.md](GAMEPLAY_GUIDE.md)** - your complete guide to mastering network packet routing!

## 🎉 Why Packet Rush?

- **🧠 Educational**: Learn network routing concepts through gameplay
- **🎮 Addictive**: Simple rules, complex strategy, endless replayability  
- **⚡ Performance**: Smooth 60fps gameplay with Go's concurrency
- **🌐 Universal**: Runs on any terminal - Windows, Mac, Linux
- **🎨 Beautiful**: Full-color interface with polished visual design
- **🏆 Competitive**: Perfect for speedrun challenges and high scores

## 🌟 Game Highlights

> *"A masterful blend of network administration simulation and puzzle gaming!"*

- **Zero Dependencies** - Single binary execution
- **Cross-Platform** - Works everywhere Go runs  
- **Professional Quality** - Production-ready codebase
- **Unique Concept** - Network router simulation gaming
- **Progressive Challenge** - From beginner to expert levels

---

**Ready to become the ultimate network administrator?** 

Build the game and start routing packets through your network empire! 🚀

*Built with ❤️ by Maverick Kamal*