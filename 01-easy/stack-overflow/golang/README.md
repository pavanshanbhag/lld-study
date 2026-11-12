# Stack Overflow — Go (modernized)

Idiomatic Go reference implementation:

- `NewStackOverflow()` constructor (no singleton)
- Fallible ops return `error` (`AskQuestion`, `Vote`, `AcceptAnswer`, …)
- Getters without `Get` prefix (`VoteCount`, `Reputation`, `Tags`)
- `log/slog` in demo
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/stack-overflow/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| `GetVoteCount()` | `VoteCount()` |
| `GetReputation()` | `Reputation()` |
| `GetQuestionsByUser()` | `QuestionsByUser()` |
| demo `fmt.Println` | `log/slog` |
