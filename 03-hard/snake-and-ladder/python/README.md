# Snake And Ladder — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 03-hard/snake-and-ladder/python/snake_and_ladder_demo.py
PYTHONPATH=03-hard/snake-and-ladder/python uv run pytest 03-hard/snake-and-ladder/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
