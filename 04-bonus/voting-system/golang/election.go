package votingsystem

import "sync"

type Election struct {
	ID         string
	Title      string
	status     ElectionStatus
	candidates map[string]*Candidate
	ballots    []Ballot
	voted      map[string]struct{}
	mu         sync.Mutex
}

func NewElection(id, title string) *Election {
	return &Election{
		ID:         id,
		Title:      title,
		status:     StatusCreated,
		candidates: make(map[string]*Candidate),
		voted:      make(map[string]struct{}),
	}
}

func (e *Election) Status() ElectionStatus {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.status
}

func (e *Election) AddCandidate(c *Candidate) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.status != StatusCreated {
		return ErrInvalidState
	}
	if _, ok := e.candidates[c.ID]; ok {
		return ErrAlreadyExists
	}
	e.candidates[c.ID] = c
	return nil
}

func (e *Election) Open() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.status != StatusCreated {
		return ErrInvalidState
	}
	if len(e.candidates) == 0 {
		return ErrInvalidState
	}
	e.status = StatusOpen
	return nil
}

func (e *Election) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.status != StatusOpen {
		return ErrInvalidState
	}
	e.status = StatusClosed
	return nil
}

func (e *Election) CastVote(voterID, candidateID string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.status != StatusOpen {
		return ErrInvalidState
	}
	if _, ok := e.candidates[candidateID]; !ok {
		return ErrUnknownCandidate
	}
	if _, ok := e.voted[voterID]; ok {
		return ErrAlreadyVoted
	}
	e.voted[voterID] = struct{}{}
	e.ballots = append(e.ballots, Ballot{VoterID: voterID, CandidateID: candidateID})
	return nil
}

func (e *Election) Ballots() []Ballot {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]Ballot, len(e.ballots))
	copy(out, e.ballots)
	return out
}

func (e *Election) Candidates() []*Candidate {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]*Candidate, 0, len(e.candidates))
	for _, c := range e.candidates {
		out = append(out, c)
	}
	return out
}
