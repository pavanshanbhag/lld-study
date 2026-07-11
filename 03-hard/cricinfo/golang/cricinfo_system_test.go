package cricinfo

import (
	"testing"
	"time"
)

func TestCricinfoSystemAddMatchAndScorecard(t *testing.T) {
	t.Parallel()

	system := NewCricinfoSystem()
	team1 := NewTeam("T1", "Team 1", []*Player{NewPlayer("P1", "Player 1", "Batsman")})
	team2 := NewTeam("T2", "Team 2", []*Player{NewPlayer("P2", "Player 2", "Bowler")})
	match := NewMatch("M1", "Final", "Stadium", time.Now(), []*Team{team1, team2})

	system.AddMatch(match)
	if system.GetMatch("M1") == nil {
		t.Fatal("expected match to be registered")
	}

	scorecardID := system.CreateScorecard(match)
	system.UpdateScore(scorecardID, "T1", 120)

	scorecard := system.GetScorecard(scorecardID)
	if scorecard == nil || scorecard.TeamScores["T1"] != 120 {
		t.Fatalf("scorecard = %+v, want T1 score 120", scorecard)
	}
}
