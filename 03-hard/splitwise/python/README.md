# Splitwise — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 03-hard/splitwise/python/splitwise_demo.py
PYTHONPATH=03-hard/splitwise/python uv run pytest 03-hard/splitwise/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
