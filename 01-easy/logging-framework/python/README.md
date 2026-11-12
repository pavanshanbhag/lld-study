# Logging Framework — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 01-easy/logging-framework/python/logging_framework_demo.py
PYTHONPATH=01-easy/logging-framework/python uv run pytest 01-easy/logging-framework/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
