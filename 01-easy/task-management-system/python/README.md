# Task Management System — Python (modernized)

Tasks, lists, observers, and sort strategies:

- `StrEnum` for status/priority; `@dataclass` for User/Comment/ActivityLog
- `logging` for state/observer feedback
- Constructor-based system (no singleton)

## Run

```bash
# from repo root
uv run python 01-easy/task-management-system/python/task_management_system_demo.py
uv run pytest 01-easy/task-management-system/python/tests -q
```

## Design patterns

- **State** — TODO / IN_PROGRESS / DONE
- **Observer** — activity logger
- **Strategy** — task sort
- **Builder** — task construction
