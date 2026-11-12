package votingsystem

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	vs := NewVotingSystem()

	if _, err := vs.CreateElection("E1", "City Mayor 2026"); err != nil {
		logger.Error("create election", "err", err)
		return
	}
	_ = vs.AddCandidate("E1", "C1", "Alice")
	_ = vs.AddCandidate("E1", "C2", "Bob")

	_, _ = vs.RegisterVoter("V1", "Nina")
	_, _ = vs.RegisterVoter("V2", "Omar")
	_, _ = vs.RegisterVoter("V3", "Priya")

	if err := vs.OpenElection("E1"); err != nil {
		logger.Error("open", "err", err)
		return
	}

	_ = vs.CastVote("E1", "V1", "C1")
	_ = vs.CastVote("E1", "V2", "C2")
	_ = vs.CastVote("E1", "V3", "C1")

	if err := vs.CastVote("E1", "V1", "C2"); err != nil {
		logger.Info("duplicate vote rejected", "err", err)
	}

	_ = vs.CloseElection("E1")
	result, err := vs.Results("E1")
	if err != nil {
		logger.Error("results", "err", err)
		return
	}
	logger.Info("tally", "winner", result.WinnerName, "total", result.TotalVotes, "counts", result.Counts)
}
