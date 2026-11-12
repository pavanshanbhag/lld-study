# Coffee Vending Machine — Go (modernized)

Idiomatic Go reference implementation:

- `NewCoffeeMachine()` constructor (no singleton)
- `SelectCoffee` / `DispenseCoffee` return errors
- Getters without `Get` prefix (`Name`, `Price`, `Recipe`)
- `log/slog` in demo
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/coffee-vending-machine/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| `GetName()` / `GetPrice()` | `Name()` / `Price()` |
| `SelectCoffee() *Coffee` | `SelectCoffee() (*Coffee, error)` |
| `DispenseCoffee() error` | `DispenseCoffee() (change float64, error)` |
| `DisplayMenu()` println | `Menu() []string` |
