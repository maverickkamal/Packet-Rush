package levels

import (
	"fmt"
	"strings"
	"time"

	"github.com/maverickkamal/Packet-Rush/internal/types"
)

type LevelData struct {
	Grid          []string
	Junctions     map[string]*types.Junction
	SpawnInterval time.Duration
	Goal          string
	TargetWord    string
}

func GetLevelData(level int) LevelData {
	switch level {
	case 1:
		return getLevelOne()
	case 2:
		return getLevelTwo()
	case 3:
		return getLevelThree()
	case 4:
		return getLevelFour()
	case 5:
		return getLevelFive()
	case 6:
		return getLevelSix()
	case 7:
		return getLevelSeven()
	case 8:
		return getLevelEight()
	case 9:
		return getLevelNine()
	case 10:
		return getLevelTen()
	default:
		return getAdvancedLevel(level)
	}
}

func getLevelOne() LevelData {
	grid := []string{
		"################################################################################",
		"#                                                                              #",
		"#  🚀 LEVEL 1: NETWORK BASICS - Spell 'GO' with packets!                     #",
		"#                                                                              #",
		"#S--------+----------+----------+----------G                                  #",
		"#         |          |          |                                            #",
		"#         |          |          |                                            #",
		"#         |          |          |                                            #",
		"#         |          |          |                                            #",
		"#         |          |          |                                            #",
		"#         |          |          O                                            #",
		"#         |          |                                                       #",
		"#         |          |                                                       #",
		"#         |          |                                                       #",
		"#         +----------+----------+----------1                                  #",
		"#                    |          |                                            #",
		"#                    |          |                                            #",
		"#                    |          2                                            #",
		"#                    |                                                       #",
		"#                    3                                                       #",
		"################################################################################",
	}

	junctions := make(map[string]*types.Junction)
	junctions["10,4"] = types.NewJunction(10, 4, []types.Position{types.Right, types.Down}, '1')
	junctions["21,4"] = types.NewJunction(21, 4, []types.Position{types.Right, types.Down}, '2')
	junctions["32,4"] = types.NewJunction(32, 4, []types.Position{types.Right, types.Down}, '3')
	junctions["10,14"] = types.NewJunction(10, 14, []types.Position{types.Right, types.Up}, '4')
	junctions["21,14"] = types.NewJunction(21, 14, []types.Position{types.Right, types.Down}, '5')
	junctions["32,14"] = types.NewJunction(32, 14, []types.Position{types.Right, types.Down}, '6')

	return LevelData{
		Grid:          grid,
		Junctions:     junctions,
		SpawnInterval: 4 * time.Second,
		Goal:          "Route G packets to 'G' and O packets to 'O' to spell 'GO'!",
		TargetWord:    "GO",
	}
}

func getLevelTwo() LevelData {
	grid := []string{
		"################################################################################",
		"#                                                                              #",
		"#  🚀 LEVEL 2: EXTENDED NETWORK - Spell 'HI' with longer paths!              #",
		"#                                                                              #",
		"#S-----+-----+-----+-----+-----+-----+-----+-----H                          #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     |                                 #",
		"#      |     |     |     |     |     |     I                                 #",
		"#      |     |     |     |     |     |                                       #",
		"#      |     |     |     |     |     |                                       #",
		"#      +-----+-----+-----+-----+-----+-----+-----1                          #",
		"#            |     |     |     |     |     |                                 #",
		"#            |     |     |     |     |     2                                 #",
		"#            |     |     |     |     3                                       #",
		"################################################################################",
	}

	junctions := make(map[string]*types.Junction)
	junctions["7,4"] = types.NewJunction(7, 4, []types.Position{types.Right, types.Down}, '1')
	junctions["13,4"] = types.NewJunction(13, 4, []types.Position{types.Right, types.Down}, '2')
	junctions["19,4"] = types.NewJunction(19, 4, []types.Position{types.Right, types.Down}, '3')
	junctions["25,4"] = types.NewJunction(25, 4, []types.Position{types.Right, types.Down}, '4')
	junctions["31,4"] = types.NewJunction(31, 4, []types.Position{types.Right, types.Down}, '5')
	junctions["37,4"] = types.NewJunction(37, 4, []types.Position{types.Right, types.Down}, '6')
	junctions["43,4"] = types.NewJunction(43, 4, []types.Position{types.Right, types.Down}, '7')

	return LevelData{
		Grid:          grid,
		Junctions:     junctions,
		SpawnInterval: 3500 * time.Millisecond,
		Goal:          "Route H packets to 'H' and I packets to 'I' to spell 'HI'!",
		TargetWord:    "HI",
	}
}

func getLevelThree() LevelData {
	grid := []string{
		"################################################################################",
		"#                                                                              #",
		"#  🚀 LEVEL 3: COMPLEX ROUTING - Spell 'WIN' with challenging paths!         #",
		"#                                                                              #",
		"#S--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--W                     #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                        #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                        #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                        #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  I                        #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                           #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                           #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  N                           #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                              #",
		"#   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                              #",
		"#   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--1                     #",
		"#      |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |                        #",
		"#      |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  2                        #",
		"#      |  |  |  |  |  |  |  |  |  |  |  |  |  |  3                           #",
		"#      |  |  |  |  |  |  |  |  |  |  |  |  |  4                              #",
		"################################################################################",
	}

	junctions := make(map[string]*types.Junction)
	for i, x := range []int{4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34, 37, 40, 43, 46, 49, 52} {
		junctions[fmt.Sprintf("%d,4", x)] = types.NewJunction(x, 4, []types.Position{types.Right, types.Down}, rune('1'+i))
	}

	return LevelData{
		Grid:          grid,
		Junctions:     junctions,
		SpawnInterval: 3 * time.Second,
		Goal:          "Route W, I, N packets to spell 'WIN'! More junctions = more complexity!",
		TargetWord:    "WIN",
	}
}

func getLevelFour() LevelData {
	return LevelData{
		Grid:          createStandardGrid(4, "CODE"),
		Junctions:     createStandardJunctions(6),
		SpawnInterval: 2500 * time.Millisecond,
		Goal:          "Multi-path routing! Spell 'CODE' with increasing complexity!",
		TargetWord:    "CODE",
	}
}

func getLevelFive() LevelData {
	return LevelData{
		Grid:          createStandardGrid(5, "RUSH"),
		Junctions:     createStandardJunctions(8),
		SpawnInterval: 2 * time.Second,
		Goal:          "Master level! Route R-U-S-H packets through the network maze!",
		TargetWord:    "RUSH",
	}
}

func getLevelSix() LevelData {
	return LevelData{
		Grid:          createStandardGrid(6, "EXPERT"),
		Junctions:     createStandardJunctions(10),
		SpawnInterval: 1800 * time.Millisecond,
		Goal:          "Expert level! Spell 'EXPERT' with precision routing!",
		TargetWord:    "EXPERT",
	}
}

func getLevelSeven() LevelData {
	return LevelData{
		Grid:          createStandardGrid(7, "GENIUS"),
		Junctions:     createStandardJunctions(12),
		SpawnInterval: 1500 * time.Millisecond,
		Goal:          "Genius level! Route 'GENIUS' packets through complex paths!",
		TargetWord:    "GENIUS",
	}
}

func getLevelEight() LevelData {
	return LevelData{
		Grid:          createStandardGrid(8, "MASTER"),
		Junctions:     createStandardJunctions(14),
		SpawnInterval: 1300 * time.Millisecond,
		Goal:          "Master level! Spell 'MASTER' with network mastery!",
		TargetWord:    "MASTER",
	}
}

func getLevelNine() LevelData {
	return LevelData{
		Grid:          createStandardGrid(9, "LEGEND"),
		Junctions:     createStandardJunctions(16),
		SpawnInterval: 1100 * time.Millisecond,
		Goal:          "Legendary routing! Spell 'LEGEND' like a true network hero!",
		TargetWord:    "LEGEND",
	}
}

func getLevelTen() LevelData {
	return LevelData{
		Grid:          createStandardGrid(10, "CHAMPION"),
		Junctions:     createStandardJunctions(20),
		SpawnInterval: 1 * time.Second,
		Goal:          "FINAL LEVEL! Spell 'CHAMPION' - prove you're the ultimate router!",
		TargetWord:    "CHAMPION",
	}
}

func getAdvancedLevel(level int) LevelData {
	words := []string{"ULTIMATE", "SUPREME", "INFINITE", "ETERNAL", "COSMIC"}
	word := words[(level-11)%len(words)]

	return LevelData{
		Grid:          createStandardGrid(level, word),
		Junctions:     createStandardJunctions(20 + (level-10)*2),
		SpawnInterval: 800 * time.Millisecond,
		Goal:          fmt.Sprintf("Advanced Level %d! Spell '%s' with ultimate skill!", level, word),
		TargetWord:    word,
	}
}

func createStandardGrid(level int, word string) []string {
	grid := []string{
		"################################################################################",
		"#                                                                              #",
		fmt.Sprintf("#  🚀 LEVEL %d: ADVANCED NETWORK - Spell '%s'!                            #", level, word),
		"#                                                                              #",
	}

	wordRunes := []rune(word)
	letterPositions := make(map[rune]string)

	for i, letter := range wordRunes {
		if i < 4 {
			y := 4 + i*4
			letterPositions[letter] = fmt.Sprintf("%d,%d", 70, y)
		} else {
			x := 10 + (i-4)*10
			letterPositions[letter] = fmt.Sprintf("%d,%d", x, 16)
		}
	}

	for i := 0; i < 16; i++ {
		line := "#S"
		for j := 0; j < 35; j++ {
			if j%3 == 0 {
				line += "-+"
			} else {
				line += "--"
			}
		}
	
		line += strings.Repeat("-", 5)
		if i == 4 && len(wordRunes) > 0 {
			line += string(wordRunes[0])
		} else if i == 8 && len(wordRunes) > 1 {
			line += string(wordRunes[1])
		} else if i == 12 && len(wordRunes) > 2 {
			line += string(wordRunes[2])
		} else if i == 16 && len(wordRunes) > 3 {
			line += string(wordRunes[3])
		} else {
			line += " "
		}

		line += strings.Repeat(" ", 80-len(line)-1) + "#"
		grid = append(grid, line)
	}

	grid = append(grid, "################################################################################")
	return grid
}

func createStandardJunctions(count int) map[string]*types.Junction {
	junctions := make(map[string]*types.Junction)

	for i := 0; i < count; i++ {
		x := 3 + i*4
		y := 4 + (i%4)*3
		if x < 70 {
			directions := []types.Position{types.Right, types.Down}
			if i%3 == 0 {
				directions = append(directions, types.Up)
			}
			junctions[fmt.Sprintf("%d,%d", x, y)] = types.NewJunction(x, y, directions, rune('1'+i%9))
		}
	}

	return junctions
}

