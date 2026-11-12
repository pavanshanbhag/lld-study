package cricinfo

type CricinfoSystem struct {
	matchService     *MatchService
	scorecardService *ScorecardService
}

func NewCricinfoSystem() *CricinfoSystem {
	return &CricinfoSystem{
		matchService:     NewMatchService(),
		scorecardService: NewScorecardService(),
	}
}

func (cs *CricinfoSystem) AddMatch(match *Match) {
	cs.matchService.AddMatch(match)
}

func (cs *CricinfoSystem) Match(matchID string) *Match {
	return cs.matchService.Match(matchID)
}

func (cs *CricinfoSystem) AllMatches() []*Match {
	return cs.matchService.AllMatches()
}

func (cs *CricinfoSystem) UpdateMatchStatus(matchID string, status MatchStatus) {
	cs.matchService.UpdateMatchStatus(matchID, status)
}

func (cs *CricinfoSystem) CreateScorecard(match *Match) string {
	return cs.scorecardService.CreateScorecard(match)
}

func (cs *CricinfoSystem) Scorecard(scorecardID string) *Scorecard {
	return cs.scorecardService.Scorecard(scorecardID)
}

func (cs *CricinfoSystem) UpdateScore(scorecardID string, teamID string, score int) {
	cs.scorecardService.UpdateScore(scorecardID, teamID, score)
}

func (cs *CricinfoSystem) AddInnings(scorecardID string, innings *Innings) {
	cs.scorecardService.AddInnings(scorecardID, innings)
}
