package levels

import (
	"time"

	"github.com/maverickkamal/Packet-Rush/internal/types"
)

func NewGameModelForLevel(level int) *types.GameModel {
	levelData := GetLevelData(level)

	return &types.GameModel{
		Grid:          levelData.Grid,
		Packets:       make([]*types.Packet, 0),
		Junctions:     levelData.Junctions,
		Score:         0,
		Lives:         types.LivesPerLevel,
		Level:         level,
		GameTime:      0,
		LastSpawn:     time.Now(),
		Paused:        false,
		GameOver:      false,
		LevelComplete: false,
		SpawnInterval: levelData.SpawnInterval,
		TickSpeed:     types.InitialTickSpeed,
		CurrentGoal:   levelData.Goal,
		GoalProgress:  make([]rune, 0),
		TargetWord:    levelData.TargetWord,
	}
}

