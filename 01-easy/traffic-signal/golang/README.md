# Traffic Signal — Go (modernized)

Idiomatic Go reference implementation:

- `NewTrafficController()` constructor (no singleton)
- `HandleEmergency` returns `error` when road missing
- Stoppable control loop via channel (demo-friendly)
- `log/slog` in demo (no domain println)
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/traffic-signal/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| `HandleEmergency` silent miss | `HandleEmergency(...) error` |
| infinite goroutine loop | `StartTrafficControl(stop <-chan struct{})` |
| light `fmt.Printf` notify | no domain println; read `Signal()` |
