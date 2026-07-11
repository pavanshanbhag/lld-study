# Parking Lot — Go (modernized)

Idiomatic Go 1.25 reference implementation:

- `NewParkingLot()` constructor (no singleton)
- `(ParkResult, error)` / `errors.Is` for failures
- `sync.RWMutex` for concurrent access
- `log/slog` in demo (no domain println)
- Table-driven tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/parking-lot/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| `GetParkingLotInstance()` | `NewParkingLot()` |
| `GetLicensePlate()` | `LicensePlate()` |
| `ParkVehicle() bool` | `Park(v) (ParkResult, error)` |
| `DisplayAvailability()` | `Availability() []SpotStatus` |
