# Voting System — Go

Idiomatic Go 1.25 bonus implementation:

- `NewVotingSystem()` constructor (no singleton)
- Typed errors (`ErrAlreadyVoted`, `ErrInvalidState`, …)
- `sync.Mutex` / `RWMutex` for concurrent ballots
- Plurality tally via `TallyStrategy`
- `log/slog` demo + `-race` tests

## Run

```bash
# from repo root — uncomment in main.go:
# votingsystem "lld-study/04-bonus/voting-system/golang"
go run .

go test -race ./04-bonus/voting-system/golang/
```
