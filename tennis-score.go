package tennis

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerScores struct {
	name     string
	points   int
	gamesWon int
}

type Match struct {
	playerOne *PlayerScores
	playerTwo *PlayerScores
}

func (match Match) PointWonBy(player *PlayerScores) {
	player.points++
	score := match.GameScore()
	if isTieBreak(match) {
		score = match.TieBreakScore()
	}
	if strings.HasPrefix(score, "Win") {
		player.gamesWon++
		match.playerOne.points = 0
		match.playerTwo.points = 0
	}
}

func (match Match) GameScore() string {
	score := ""
	tempScore := 0

	if match.playerOne.points == match.playerTwo.points {
		switch match.playerOne.points {
		case 0:
			score = "0-0"
		case 1:
			score = "15-15"
		case 2:
			score = "30-30"
		default:
			score = "Deuce"
		}
	} else if match.playerOne.points >= 4 || match.playerTwo.points >= 4 {
		lead := match.playerOne.points - match.playerTwo.points
		if lead == 1 {
			score = "Advantage " + match.playerOne.name
		} else if lead == -1 {
			score = "Advantage " + match.playerTwo.name
		} else if lead >= 2 {
			score = "Win " + match.playerOne.name
		} else {
			score = "Win " + match.playerTwo.name
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = match.playerOne.points
			} else {
				score += "-"
				tempScore = match.playerTwo.points
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

func (match Match) TieBreakScore() string {
	score := ""
	lead := match.playerOne.points - match.playerTwo.points
	if match.playerOne.points >= 7 && lead >= 2 {
		score = "Win " + match.playerOne.name
	} else if match.playerTwo.points >= 7 && lead <= -2 {
		score = "Win " + match.playerTwo.name
	} else {
		score = strconv.Itoa(match.playerOne.points) + "-" + strconv.Itoa(match.playerTwo.points)
	}

	return score
}

func (match Match) Score() string {
	score := ""
	lead := match.playerOne.gamesWon - match.playerTwo.gamesWon

	if (match.playerOne.gamesWon >= 6 && lead >= 2) || (match.playerOne.gamesWon == 7 && lead == 1) {
		score = "Match win " + match.playerOne.name
	} else if (match.playerTwo.gamesWon >= 6 && lead <= -2) || (match.playerTwo.gamesWon == 7 && lead == -1) {
		score = "Match win " + match.playerTwo.name
	} else if isTieBreak(match) {
		score = match.TieBreakScore() + " Tie Break"
	} else {
		score = match.GameScore()
	}

	return strconv.Itoa(match.playerOne.gamesWon) + "-" + strconv.Itoa(match.playerTwo.gamesWon) + ", " + score
}

func isTieBreak(match Match) bool {
	return match.playerOne.gamesWon == 6 && match.playerTwo.gamesWon == 6
}

func main() {
	fmt.Println("hello")
}
