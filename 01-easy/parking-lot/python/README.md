# Parking Lot — Python (modernized)

Single-module implementation in `parking_lot.py`:

- `@dataclass` / `StrEnum` / `Protocol`
- `datetime` for ticket duration (no millisecond math)
- Constructor-based `ParkingLot` (no singleton)
- `logging` in domain code; demo uses stdout for UX
- Thread-safe spots and lot operations

## Run

```bash
# from repo root
uv run python 01-easy/parking-lot/python/parking_lot_demo.py
uv run pytest 01-easy/parking-lot/python/tests -q
```

## Design patterns

- **Strategy** — `ParkingStrategy`, `FeeStrategy`
- **Factory** — vehicle dataclasses (`Bike`, `Car`, `Truck`)
