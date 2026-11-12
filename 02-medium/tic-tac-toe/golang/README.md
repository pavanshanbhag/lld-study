# Tic-Tac-Toe — Go (modernized)

Idiomatic Go implementation:

- `NewGame(p1, p2)` constructor (no singleton)
- `MakeMove` returns `error` for invalid moves
- `log/slog` for game flow logs
- Interactive stdin prompts remain on stdout
- Table-driven board tests + `-race`

## Run

```bash
go test -race ./02-medium/tic-tac-toe/golang/...
```
