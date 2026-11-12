# Coffee Vending Machine — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 01-easy/coffee-vending-machine/python/coffee_vending_machine_demo.py
PYTHONPATH=01-easy/coffee-vending-machine/python uv run pytest 01-easy/coffee-vending-machine/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
