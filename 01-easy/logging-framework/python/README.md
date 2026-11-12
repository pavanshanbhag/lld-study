# Logging Framework — Python (modernized)

Hierarchical async logging with pluggable appenders:

- `IntEnum` `LogLevel`, `@dataclass` `LogMessage`
- `logging` for framework diagnostics; console appender remains the sink
- Constructor-based `LogManager` (no singleton)

## Run

```bash
# from repo root
uv run python 01-easy/logging-framework/python/logging_framework_demo.py
uv run pytest 01-easy/logging-framework/python/tests -q
```

## Design patterns

- **Composite** — logger hierarchy with parent fallback
- **Strategy** — formatters / appenders
