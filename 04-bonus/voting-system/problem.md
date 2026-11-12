# Designing a Voting System

## Requirements
1. Support multiple elections; each election has a title, candidates, and a lifecycle (Created → Open → Closed).
2. Register voters; a voter may cast at most one ballot per election.
3. While an election is open, eligible voters cast a vote for one candidate.
4. Closing an election freezes ballots and enables tallying.
5. Report results: vote counts per candidate and the winner (plurality).
6. Reject invalid actions: voting twice, voting when closed, unknown voter/candidate.
7. Support concurrent casting of ballots safely.

## UML Class Diagram

See `class-diagram.mmd` (modernized). Upstream PNG may be absent for this bonus topic.

## Implementations
- [Go](./golang/)
- [Python](./python/)

## Classes
1. **VotingSystem** — facade: create elections, register voters, open/close, cast vote, results.
2. **Election** — candidates, status, ballots.
3. **Candidate** / **Voter** — identities.
4. **Ballot** — voter + candidate for one election.
5. **ElectionStatus** — Created | Open | Closed.
6. **TallyStrategy** (plurality) — compute counts and winner.

## Design Patterns
1. **Facade** — `VotingSystem` orchestrates domain objects.
2. **Strategy** — tallying (plurality; extensible to ranked choice later).
3. **State** (via status enum) — allowed operations depend on election lifecycle.
