# Logging Framework — Go (modernized)

Idiomatic Go reference implementation:

- `NewLogger(config)` constructor (no singleton; nil → defaults)
- Log methods return `error` from appenders
- Level filtering under `RWMutex`
- Demo uses `log/slog` for its own flow; console/file appenders remain the sink
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/logging-framework/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| singleton logger | `NewLogger(config)` |
| silent log calls | `Info/Debug/...() error` |
| demo `fmt.Printf` errors | `log/slog` |
