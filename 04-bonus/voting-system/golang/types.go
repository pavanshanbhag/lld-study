package votingsystem

type ElectionStatus int

const (
	StatusCreated ElectionStatus = iota
	StatusOpen
	StatusClosed
)

func (s ElectionStatus) String() string {
	switch s {
	case StatusCreated:
		return "Created"
	case StatusOpen:
		return "Open"
	case StatusClosed:
		return "Closed"
	default:
		return "Unknown"
	}
}

type Candidate struct {
	ID   string
	Name string
}

func NewCandidate(id, name string) *Candidate {
	return &Candidate{ID: id, Name: name}
}

type Voter struct {
	ID   string
	Name string
}

func NewVoter(id, name string) *Voter {
	return &Voter{ID: id, Name: name}
}

type Ballot struct {
	VoterID     string
	CandidateID string
}
