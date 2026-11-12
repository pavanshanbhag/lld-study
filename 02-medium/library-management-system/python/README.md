# Library Management System — Python

Constructor-based solution (no `get_instance` singleton).

## Run

```bash
# from repo root
uv run python 02-medium/library-management-system/python/library_management_system_demo.py
PYTHONPATH=02-medium/library-management-system/python uv run pytest 02-medium/library-management-system/python/tests -q
```

## Notes

- Prefer `ClassName(...)` constructors over singletons
- Topic tests use isolated `PYTHONPATH` via `./tools/check.sh`
