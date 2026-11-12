# Stack Overflow — Python (modernized)

Forum domain with questions, answers, votes, and reputation:

- `@dataclass` / `StrEnum` for models and event types
- Thread-safe reputation updates
- Constructor-based service (no singleton)

## Run

```bash
# from repo root
uv run python 01-easy/stack-overflow/python/stack_overflow_demo.py
uv run pytest 01-easy/stack-overflow/python/tests -q
```

## Design patterns

- **Observer** — post events → reputation manager
- **Strategy** — search strategies
