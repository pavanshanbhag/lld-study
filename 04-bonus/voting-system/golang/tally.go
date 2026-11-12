package votingsystem

type TallyResult struct {
	Counts     map[string]int
	WinnerID   string
	WinnerName string
	TotalVotes int
}

type TallyStrategy interface {
	Tally(election *Election) TallyResult
}

type PluralityTally struct{}

func NewPluralityTally() *PluralityTally {
	return &PluralityTally{}
}

func (p *PluralityTally) Tally(election *Election) TallyResult {
	counts := make(map[string]int)
	for _, c := range election.Candidates() {
		counts[c.ID] = 0
	}
	for _, b := range election.Ballots() {
		counts[b.CandidateID]++
	}

	result := TallyResult{Counts: counts, TotalVotes: len(election.Ballots())}
	var best int = -1
	for _, c := range election.Candidates() {
		n := counts[c.ID]
		if n > best {
			best = n
			result.WinnerID = c.ID
			result.WinnerName = c.Name
		}
	}
	return result
}
