package main

import (
	"fmt"
	"strconv"
)

type Game struct {
	Player1 *PlayerPoints
	Player2 *PlayerPoints
}

type PlayerPoints struct {
	Name  string
	Points int
}

func (game Game) PointWonBy(player *Player) {
	player.Points += 1
}

func (game Game) Score() string {
	score := ""
	tempScore := 0

	if game.Player1.Points == game.Player2.Points {
		switch game.Player1.Points {
		case 0:
			score = "0-0"
		case 1:
			score = "15-15"
		case 2:
			score = "30-30"
		default:
			score = "Deuce"
		}
	} else if game.Player1.Points >= 4 || game.Player2.Points >= 4 {
		lead := game.Player1.Points - game.Player2.Points
		if lead == 1 {
			score = "Advantage " + game.Player1.Name
		} else if lead == -1 {
			score = "Advantage " + game.Player2.Name
		} else if lead >= 2 {
			score = "Win  " + game.Player1.Name
		} else {
			score = "Win  " + game.Player2.Name
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.Player1.Points
			} else {
				score += "-"
				tempScore = game.Player2.Points
			}
			switch tempScore {
			case 0:
				score += "0"
			case 1:
				score += "15"
			case 2:
				score += "30"
			case 3:
				score += "40"
			}
		}
	}
	return score
}

func (game Game) printScore() string {
	a := game.Player1.Name + " " + strconv.Itoa(game.Player1.Points) + " - " + strconv.Itoa(game.Player2.Points) + " " + game.Player2.Name
	return a
}

func main() {

	var player1 = &Player{"Lendl", 0}
	var player2 = &Player{"Becker", 0}
	var game = Game{player1, player2}

	game.PointWonBy(player1)
	fmt.Println(game.printScore())
	fmt.Println(game.Score())
}
