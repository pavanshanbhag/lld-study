# Ride Sharing Service — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 03-hard/ride-sharing-service/python/ride_sharing_demo.py
PYTHONPATH=03-hard/ride-sharing-service/python uv run pytest 03-hard/ride-sharing-service/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
