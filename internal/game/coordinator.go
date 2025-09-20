package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maverickkamal/Packet-Rush/internal/levels"
	"github.com/maverickkamal/Packet-Rush/internal/types"
)


type GameCoordinator struct {
	Model *types.GameModel
}


func NewGameCoordinator() *GameCoordinator {
	return &GameCoordinator{
		Model: levels.NewGameModelForLevel(1),
	}
}


func (gc *GameCoordinator) Init() tea.Cmd {
	return gc.Model.Init()
}


func (gc *GameCoordinator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := gc.Model.Update(msg)
	gc.Model = model.(*types.GameModel)

	if gc.Model.RestartRequested {
		gc.Model = levels.NewGameModelForLevel(1)
		gc.Model.RestartRequested = false
		return gc, gc.Model.Init()
	}

	if gc.Model.NextLevelRequested {
		oldScore := gc.Model.Score
		if gc.Model.Level < types.MaxLevel {
			gc.Model = levels.NewGameModelForLevel(gc.Model.Level + 1)
			gc.Model.Score = oldScore 
		} else {
			gc.Model = levels.NewGameModelForLevel(1)
			gc.Model.Score = oldScore 
		}
		gc.Model.NextLevelRequested = false
		return gc, gc.Model.Init()
	}

	return gc, cmd
}


func (gc *GameCoordinator) View() string {
	return gc.Model.View()
}
