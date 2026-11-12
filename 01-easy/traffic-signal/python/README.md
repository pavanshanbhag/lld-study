# Traffic Signal — Python (modernized)

Intersection controllers with timed light cycles:

- `StrEnum` for `Direction` / `LightColor`
- `logging` for monitor and cycle feedback
- Constructor-based `TrafficControlSystem` (no singleton)

## Run

```bash
# from repo root
uv run python 01-easy/traffic-signal/python/traffic_system_demo.py
uv run pytest 01-easy/traffic-signal/python/tests -q
```

## Design patterns

- **State** — north-south / east-west green cycles
- **Observer** — central monitor on light changes
- **Builder** — intersection construction
