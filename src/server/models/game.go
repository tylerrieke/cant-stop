package models

import (
	"math/rand"
)

const MIN_PLAYERS = 2
const MAX_PLAYERS = 4
const MAX_CURRENT_COLUMNS = 3
const WINNING_SCORE = 3

type Phase string

const TURN_START = "TURN_START"
const SELECTION = "SELECTION"
const ROLL_OR_STOP = "ROLL_OR_STOP"
const BUST = "BUST"
const GAME_OVER = "GAME_OVER"

var PlayerColors = [MAX_PLAYERS]string{
	"blue",
	"red",
	"yellow",
	"green",
}

var ColumnMaximums = map[int]int{
	2:  3,
	3:  5,
	4:  7,
	5:  9,
	6:  11,
	7:  13,
	8:  11,
	9:  9,
	10: 7,
	11: 5,
	12: 3,
}

type Player struct {
	Color        string
	ColumnValues map[int]int
}

type Game struct {
	CurrentPlayerIndex int
	CurrentTurnColumns map[int]int
	TurnPhase          string
	Players            []Player
	Dice               []int
	Options            [][][]int
	WinningPlayerIndex int
}

func (g *Game) canRoll() bool {
	return g.TurnPhase == TURN_START || g.TurnPhase == ROLL_OR_STOP
}

func (g *Game) Roll() *Game {
	if !g.canRoll() {
		panic("game is not in a phase for rolling")
	}

	g.Dice = []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}

	println("New Dice")

	g.updateOptions()

	println("Options Updated " + g.TurnPhase)

	return g
}

func (g *Game) SelectOption(option [][]int) *Game {
	if g.TurnPhase != SELECTION || !containsValueSet(g.Options, option) {
		panic("not a valid option")
	}
	for _, pair := range option {
		var value = pair[0] + pair[1]
		g.CurrentTurnColumns[value] = g.optionValueCount(value) + 1
	}
	g.Options = [][][]int{}
	g.TurnPhase = ROLL_OR_STOP
	return g
}

func (g *Game) Stop() *Game {
	if g.TurnPhase != ROLL_OR_STOP {
		panic("not a valid state to stop")
	}
	for value, count := range g.CurrentTurnColumns {
		g.Players[g.CurrentPlayerIndex].ColumnValues[value] = count
	}
	g = g.evaluateGameOver()
	if g.TurnPhase == GAME_OVER {
		return g
	} else {
		return g.NextPlayerTurn()
	}
}

func (g *Game) NextPlayerTurn() *Game {
	if g.TurnPhase != ROLL_OR_STOP && g.TurnPhase != BUST {
		panic("not a valid state to switch players")
	}
	g.CurrentTurnColumns = map[int]int{}
	g.CurrentPlayerIndex = (g.CurrentPlayerIndex + 1) % len(g.Players)
	g.TurnPhase = TURN_START
	return g
}

func (g *Game) updateOptions() *Game {
	var potentialPairs = [][][]int{
		{{g.Dice[0], g.Dice[1]}, {g.Dice[2], g.Dice[3]}},
		{{g.Dice[0], g.Dice[2]}, {g.Dice[1], g.Dice[3]}},
		{{g.Dice[0], g.Dice[3]}, {g.Dice[1], g.Dice[2]}},
	}

	var options = make([][][]int, 0)

	for _, pairSet := range potentialPairs {
		var first = pairSet[0][0] + pairSet[0][1]
		var second = pairSet[1][0] + pairSet[1][1]
		var firstValid = g.isValidOption(first)
		var secondValid = g.isValidOption(second)
		if firstValid && secondValid {
			if first == second {
				if g.optionValueCount(first) <= ColumnMaximums[first]+2 {
					options = appendIfNotPresent(options, [][]int{pairSet[0], pairSet[1]})
				} else {
					options = appendIfNotPresent(options, [][]int{pairSet[0]})
				}
			} else {
				if len(g.CurrentTurnColumns) == (MAX_CURRENT_COLUMNS-1) && g.CurrentTurnColumns[first] == 0 && g.CurrentTurnColumns[second] == 0 {
					options = appendIfNotPresent(options, [][]int{pairSet[0]})
					options = appendIfNotPresent(options, [][]int{pairSet[1]})
				} else {
					options = appendIfNotPresent(options, [][]int{pairSet[0], pairSet[1]})
				}
			}
		} else if firstValid {
			options = appendIfNotPresent(options, [][]int{pairSet[0]})
		} else if secondValid {
			options = appendIfNotPresent(options, [][]int{pairSet[1]})
		}
	}

	g.Options = options

	if len(options) == 0 {
		println("Setting to BUST")
		g.TurnPhase = BUST
	} else {
		println("Setting to SELECTION")
		g.TurnPhase = SELECTION
	}

	println("Before return = " + g.TurnPhase)

	return g
}

func (g *Game) optionValueCount(option int) int {
	return max(g.Players[g.CurrentPlayerIndex].ColumnValues[option], g.CurrentTurnColumns[option])
}

func (g *Game) isValueReached(option int) bool {
	for _, player := range g.Players {
		if player.ColumnValues[option] == ColumnMaximums[option] {
			return true
		}
	}
	return false
}

func (g *Game) isValidOption(option int) bool {
	var previousValue = g.optionValueCount(option)
	if previousValue >= ColumnMaximums[option] || g.isValueReached(option) {
		return false
	}

	return g.CurrentTurnColumns[option] > 0 || len(g.CurrentTurnColumns) < MAX_CURRENT_COLUMNS
}

func containsValueSet(options [][][]int, pairs [][]int) bool {
	var first = pairs[0][0] + pairs[0][1]
	var second = 0
	if len(pairs) > 1 {
		second = pairs[1][0] + pairs[1][1]
	}
	for _, pairSet := range options {
		var firstTest = pairSet[0][0] + pairSet[0][1]
		var secondTest = 0
		if len(pairSet) > 1 {
			secondTest = pairSet[1][0] + pairSet[1][1]
		}
		if (firstTest == first && secondTest == second) || (firstTest == second && secondTest == first) {
			return true
		}
	}
	return false
}

func appendIfNotPresent(options [][][]int, pairs [][]int) [][][]int {
	if !containsValueSet(options, pairs) {
		return append(options, pairs)
	} else {
		return options
	}
}

func (g *Game) clearOtherPlayersFromColumn(playerIndex int, column int) *Game {
	for index, player := range g.Players {
		if index != playerIndex {
			delete(player.ColumnValues, column)
		}
	}
	return g
}

func (g *Game) evaluatePlayerScore(playerIndex int) int {
	var score = 0
	for column, value := range g.Players[playerIndex].ColumnValues {
		if ColumnMaximums[column] == value {
			score++
			g.clearOtherPlayersFromColumn(playerIndex, column)
		}
	}
	return score
}

func (g *Game) evaluateGameOver() *Game {
	for index, _ := range g.Players {
		if g.evaluatePlayerScore(index) >= WINNING_SCORE {
			g.WinningPlayerIndex = index
			g.TurnPhase = GAME_OVER
		}
	}
	return g
}
