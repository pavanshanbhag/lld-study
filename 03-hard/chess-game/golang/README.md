# Chess Game — Go

Constructor-based LLD solution (`New…`, no singleton). Demo uses `log/slog`.

## Run

```bash
# from repo root — uncomment import + Run() in main.go, then:
go run .

# tests
go test -race ./03-hard/chess-game/golang/...
```

## Notes

- Prefer `NewX()` constructors over `GetInstance`
- Fallible ops return `error` where applicable
- See `docs/MODERNIZATION.md` for toolchain
