# Vending Machine — Python (modernized)

State-machine implementation with inventory and coin handling:

- `@dataclass` `Item`, `IntEnum` `Coin`
- `logging` in domain/state feedback; demo uses stdout for UX
- Constructor-based `VendingMachine` (no singleton)

## Run

```bash
# from repo root
uv run python 01-easy/vending-machine/python/vending_machine_demo.py
uv run pytest 01-easy/vending-machine/python/tests -q
```

## Design patterns

- **State** — idle / selected / has-money / dispensing
- **Inventory** — stock map keyed by item code
