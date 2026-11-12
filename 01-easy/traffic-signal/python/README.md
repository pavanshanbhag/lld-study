# Traffic Signal — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 01-easy/traffic-signal/python/traffic_system_demo.py
PYTHONPATH=01-easy/traffic-signal/python uv run pytest 01-easy/traffic-signal/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
