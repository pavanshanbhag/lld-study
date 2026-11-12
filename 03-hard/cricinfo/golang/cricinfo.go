package cricinfo

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	team1Players := []*Player{
		NewPlayer("P101", "Player 1", "Batsman"),
		NewPlayer("P102", "Player 2", "Bowler"),
		NewPlayer("P103", "Player 3", "All-rounder"),
	}
	team2Players := []*Player{
		NewPlayer("P201", "Player 4", "Batsman"),
		NewPlayer("P202", "Player 5", "Bowler"),
		NewPlayer("P203", "Player 6", "All-rounder"),
	}

	team1 := NewTeam("T1", "Team 1", team1Players)
	team2 := NewTeam("T2", "Team 2", team2Players)
	teams := []*Team{team1, team2}

	match := NewMatch("M001", "Match 1", "Venue 1", time.Now(), teams)
	cricinfoSystem := NewCricinfoSystem()
	cricinfoSystem.AddMatch(match)

	scorecardID := cricinfoSystem.CreateScorecard(match)
	cricinfoSystem.UpdateScore(scorecardID, "T1", 100)
	cricinfoSystem.UpdateScore(scorecardID, "T2", 75)

	innings1 := NewInnings("I1", "T1", "T2")
	innings2 := NewInnings("I2", "T2", "T1")

	over1 := NewOver(1)
	over1.AddBall(NewBall(1, "P202", "P101", "4"))
	over1.AddBall(NewBall(2, "P202", "P101", "6"))
	innings1.AddOver(over1)

	over2 := NewOver(2)
	over2.AddBall(NewBall(1, "P102", "P201", "1"))
	over2.AddBall(NewBall(2, "P102", "P201", "0"))
	innings1.AddOver(over2)

	cricinfoSystem.AddInnings(scorecardID, innings1)
	cricinfoSystem.AddInnings(scorecardID, innings2)

	displayScorecard(logger, cricinfoSystem.Scorecard(scorecardID))
}

func displayScorecard(logger *slog.Logger, scorecard *Scorecard) {
	logger.Info("scorecard", "id", scorecard.ID, "match", scorecard.Match.Title)
	for teamID, score := range scorecard.TeamScores {
		logger.Info("team score", "team", teamID, "score", score)
	}
	for _, innings := range scorecard.Innings {
		logger.Info("innings", "id", innings.ID, "batting", innings.BattingTeamID, "bowling", innings.BowlingTeamID)
		for _, over := range innings.Overs {
			for _, ball := range over.Balls {
				logger.Info("ball",
					"over", over.OverNumber,
					"ball", ball.BallNumber,
					"bowler", ball.Bowler,
					"batsman", ball.Batsman,
					"result", ball.Result,
				)
			}
		}
	}
}
