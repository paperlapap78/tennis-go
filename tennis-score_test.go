package tennis

import "testing"
import "fmt"

func TestGame(t *testing.T) {
	var gameTests = []struct {
		scoresOne int
		scoresTwo int
		expected  string
	}{
		{1, 0, "0-0, 15-0"},
		{2, 0, "0-0, 30-0"},
		{3, 0, "0-0, 40-0"},
		{0, 1, "0-0, 0-15"},
		{2, 2, "0-0, 30-30"},
		{4, 0, "1-0, 0-0"},
		{0, 4, "0-1, 0-0"},
		{2, 4, "0-1, 0-0"},
		{3, 4, "0-0, Advantage Becker"},
		{4, 4, "0-0, Deuce"},
		{4, 3, "0-0, Advantage Lendl"},
		{5, 0, "1-0, 15-0"},
	}

	for _, gt := range gameTests {
		playerOne := &PlayerScores{"Lendl", 0, 0}
		playerTwo := &PlayerScores{"Becker", 0, 0}
		testGame := Match{playerOne, playerTwo}

		maxPoints := max(gt.scoresOne, gt.scoresTwo)
		for a := 1; a <= maxPoints; a++ {
			if a <= gt.scoresOne {
				testGame.PointWonBy(playerOne)
			}
			if a <= gt.scoresTwo {
				testGame.PointWonBy(playerTwo)
			}
		}
		actual := testGame.Score()
		if actual != gt.expected {
			t.Error("expected " + gt.expected + ", got: " + actual)
		}
		fmt.Println(actual)
	}
}

func TestMatchSet(t *testing.T) {
	var matchSetTests = []struct {
		gamesWonOne int
		gamesWonTwo int
		scoresOne   int
		scoresTwo   int
		expected    string
	}{
		{4, 5, 0, 4, "4-6, Match win Becker"},
		{5, 4, 4, 0, "6-4, Match win Lendl"},
		{5, 5, 4, 0, "6-5, 0-0"},
		{5, 5, 0, 4, "5-6, 0-0"},
		{5, 6, 0, 4, "5-7, Match win Becker"},
		{6, 5, 4, 0, "7-5, Match win Lendl"},
		{6, 5, 0, 4, "6-6, 0-0 Tie Break"},
		{6, 6, 0, 5, "6-6, 0-5 Tie Break"},
		{6, 6, 0, 7, "6-7, Match win Becker"},
		{6, 6, 6, 7, "6-6, 6-7 Tie Break"},
		{6, 6, 6, 8, "6-7, Match win Becker"},
		{6, 6, 9, 7, "7-6, Match win Lendl"},
	}

	for _, test := range matchSetTests {
		playerOne := &PlayerScores{"Lendl", 0, test.gamesWonOne}
		playerTwo := &PlayerScores{"Becker", 0, test.gamesWonTwo}
		testMatch := Match{playerOne, playerTwo}

		maxPoints := max(test.scoresOne, test.scoresTwo)
		for a := 1; a <= maxPoints; a++ {
			if a <= test.scoresOne {
				testMatch.PointWonBy(playerOne)
			}
			if a <= test.scoresTwo {
				testMatch.PointWonBy(playerTwo)
			}
		}

		actual := testMatch.Score()
		if actual != test.expected {
			t.Error("expected " + test.expected + ", got: " + actual)
		}
		fmt.Println(actual)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
