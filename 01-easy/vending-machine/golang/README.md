# Vending Machine — Go (modernized)

Idiomatic Go reference implementation:

- `NewVendingMachine()` constructor (no singleton)
- State methods return `error` for invalid transitions
- Inventory `Dispense` decrements stock and fails when empty
- `log/slog` in demo (no domain println)
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/vending-machine/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| silent `SelectProduct` fail | `SelectProduct(...) error` |
| state `fmt.Println` feedback | `error` returns |
| demo `fmt.Println` | `log/slog` |
