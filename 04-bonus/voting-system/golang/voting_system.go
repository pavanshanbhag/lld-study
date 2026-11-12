package votingsystem

import "sync"

type VotingSystem struct {
	elections map[string]*Election
	voters    map[string]*Voter
	tally     TallyStrategy
	mu        sync.RWMutex
}

func NewVotingSystem() *VotingSystem {
	return &VotingSystem{
		elections: make(map[string]*Election),
		voters:    make(map[string]*Voter),
		tally:     NewPluralityTally(),
	}
}

func (vs *VotingSystem) CreateElection(id, title string) (*Election, error) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	if _, ok := vs.elections[id]; ok {
		return nil, ErrAlreadyExists
	}
	e := NewElection(id, title)
	vs.elections[id] = e
	return e, nil
}

func (vs *VotingSystem) RegisterVoter(id, name string) (*Voter, error) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	if _, ok := vs.voters[id]; ok {
		return nil, ErrAlreadyExists
	}
	v := NewVoter(id, name)
	vs.voters[id] = v
	return v, nil
}

func (vs *VotingSystem) AddCandidate(electionID, candidateID, name string) error {
	e, err := vs.election(electionID)
	if err != nil {
		return err
	}
	return e.AddCandidate(NewCandidate(candidateID, name))
}

func (vs *VotingSystem) OpenElection(electionID string) error {
	e, err := vs.election(electionID)
	if err != nil {
		return err
	}
	return e.Open()
}

func (vs *VotingSystem) CloseElection(electionID string) error {
	e, err := vs.election(electionID)
	if err != nil {
		return err
	}
	return e.Close()
}

func (vs *VotingSystem) CastVote(electionID, voterID, candidateID string) error {
	vs.mu.RLock()
	_, ok := vs.voters[voterID]
	vs.mu.RUnlock()
	if !ok {
		return ErrUnknownVoter
	}
	e, err := vs.election(electionID)
	if err != nil {
		return err
	}
	return e.CastVote(voterID, candidateID)
}

func (vs *VotingSystem) Results(electionID string) (TallyResult, error) {
	e, err := vs.election(electionID)
	if err != nil {
		return TallyResult{}, err
	}
	if e.Status() != StatusClosed {
		return TallyResult{}, ErrInvalidState
	}
	return vs.tally.Tally(e), nil
}

func (vs *VotingSystem) election(id string) (*Election, error) {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	e, ok := vs.elections[id]
	if !ok {
		return nil, ErrNotFound
	}
	return e, nil
}
