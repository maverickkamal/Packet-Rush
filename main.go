package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maverickkamal/Packet-Rush/internal/game"
)

func main() {
	fmt.Println("🚀 Starting Packet Rush...")

	coordinator := game.NewGameCoordinator()

	p := tea.NewProgram(coordinator)

	if _, err := p.Run(); err != nil {
		log.Printf("Error running game: %v", err)
		os.Exit(1)
	}

	fmt.Println("Thanks for playing Packet Rush! - Maverick Kamal")
}
