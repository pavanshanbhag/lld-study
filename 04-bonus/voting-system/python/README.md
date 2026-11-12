# Voting System — Python

Single-module implementation in `voting_system.py`:

- `@dataclass` / `StrEnum`
- Constructor-based `VotingSystem` (no singleton)
- Thread-safe election ballots
- Plurality tally strategy
- Domain exceptions; demo uses logging

## Run

```bash
uv run python 04-bonus/voting-system/python/voting_system_demo.py
PYTHONPATH=04-bonus/voting-system/python uv run pytest 04-bonus/voting-system/python/tests -q
```
