# Tic Tac Toe — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 02-medium/tic-tac-toe/python/tictactoe_demo.py
PYTHONPATH=02-medium/tic-tac-toe/python uv run pytest 02-medium/tic-tac-toe/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
