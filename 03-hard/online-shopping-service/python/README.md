# Online Shopping Service — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 03-hard/online-shopping-service/python/online_shopping_demo.py
PYTHONPATH=03-hard/online-shopping-service/python uv run pytest 03-hard/online-shopping-service/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
