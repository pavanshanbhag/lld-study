# Coffee Vending Machine — Python (modernized)

State machine + recipe inventory for coffee drinks:

- `StrEnum` for coffee/ingredient/topping types
- `logging` for domain feedback; demo uses stdout for UX
- Constructor-based machine with injectable inventory

## Run

```bash
# from repo root
uv run python 01-easy/coffee-vending-machine/python/coffee_vending_machine_demo.py
uv run pytest 01-easy/coffee-vending-machine/python/tests -q
```

## Design patterns

- **State** — ready / selecting / paid / out-of-ingredients
- **Template Method** — coffee prepare steps
- **Decorator** — toppings
