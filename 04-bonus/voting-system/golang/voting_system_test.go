package votingsystem

import (
	"errors"
	"sync"
	"testing"
)

func TestVotingHappyPath(t *testing.T) {
	t.Parallel()
	vs := NewVotingSystem()
	if _, err := vs.CreateElection("E1", "Mayor"); err != nil {
		t.Fatal(err)
	}
	_ = vs.AddCandidate("E1", "C1", "Alice")
	_ = vs.AddCandidate("E1", "C2", "Bob")
	_, _ = vs.RegisterVoter("V1", "A")
	_, _ = vs.RegisterVoter("V2", "B")
	_, _ = vs.RegisterVoter("V3", "C")

	if err := vs.OpenElection("E1"); err != nil {
		t.Fatal(err)
	}
	_ = vs.CastVote("E1", "V1", "C1")
	_ = vs.CastVote("E1", "V2", "C1")
	_ = vs.CastVote("E1", "V3", "C2")
	_ = vs.CloseElection("E1")

	r, err := vs.Results("E1")
	if err != nil {
		t.Fatal(err)
	}
	if r.WinnerID != "C1" || r.TotalVotes != 3 {
		t.Fatalf("got %+v", r)
	}
}

func TestDuplicateVote(t *testing.T) {
	t.Parallel()
	vs := NewVotingSystem()
	_, _ = vs.CreateElection("E1", "Mayor")
	_ = vs.AddCandidate("E1", "C1", "Alice")
	_, _ = vs.RegisterVoter("V1", "A")
	_ = vs.OpenElection("E1")
	_ = vs.CastVote("E1", "V1", "C1")
	if err := vs.CastVote("E1", "V1", "C1"); !errors.Is(err, ErrAlreadyVoted) {
		t.Fatalf("err = %v", err)
	}
}

func TestCannotVoteWhenClosed(t *testing.T) {
	t.Parallel()
	vs := NewVotingSystem()
	_, _ = vs.CreateElection("E1", "Mayor")
	_ = vs.AddCandidate("E1", "C1", "Alice")
	_, _ = vs.RegisterVoter("V1", "A")
	_ = vs.OpenElection("E1")
	_ = vs.CloseElection("E1")
	if err := vs.CastVote("E1", "V1", "C1"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("err = %v", err)
	}
}

func TestConcurrentVotes(t *testing.T) {
	t.Parallel()
	vs := NewVotingSystem()
	_, _ = vs.CreateElection("E1", "Mayor")
	_ = vs.AddCandidate("E1", "C1", "Alice")
	_ = vs.OpenElection("E1")

	const n = 50
	for i := range n {
		id := string(rune('A'+i%26)) + string(rune('0'+i/26))
		_, _ = vs.RegisterVoter(id, id)
	}

	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := string(rune('A'+i%26)) + string(rune('0'+i/26))
			_ = vs.CastVote("E1", id, "C1")
		}(i)
	}
	wg.Wait()
	_ = vs.CloseElection("E1")
	r, err := vs.Results("E1")
	if err != nil {
		t.Fatal(err)
	}
	if r.TotalVotes != n {
		t.Fatalf("total = %d want %d", r.TotalVotes, n)
	}
}
