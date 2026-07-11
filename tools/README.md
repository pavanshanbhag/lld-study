# Tools

## Toolchain

- **Go 1.25** — module at repo root (`go.mod`)
- **Python 3.14** — managed with `uv` (`pyproject.toml`, `.python-version`)

```bash
uv sync
./tools/check.sh all       # full suite
./tools/check.sh easy      # 01-easy Go + Python
./tools/check.sh medium    # 02-medium Go + Python
./tools/check.sh hard      # 03-hard Go + Python
./tools/check.sh patterns  # 19 design-pattern Go demos
./tools/check.sh parking   # parking-lot reference (ruff + mypy strict)
```

## Running Go LLD demos

From the **repo root**:

1. Open `main.go`
2. Uncomment the import and `Run()` call for the topic you want
3. Run: `go run .`

### Go tests

```bash
go test -race ./01-easy/parking-lot/golang/
go test -race ./02-medium/atm/golang/
go test -race ./03-hard/splitwise/golang/
```

Or run all: `./tools/check.sh test-go`

## Running Python LLD demos

```bash
uv run python 01-easy/parking-lot/python/parking_lot_demo.py
uv run python 02-medium/atm/python/atm_demo.py
uv run python 03-hard/splitwise/python/splitwise_demo.py
```

### Python tests (isolated import path per topic)

```bash
PYTHONPATH=01-easy/parking-lot/python uv run pytest 01-easy/parking-lot/python/tests -q
```

Or run all: `./tools/check.sh test-py`

### Strict lint (parking-lot reference only)

```bash
uv run ruff check 01-easy/parking-lot/python
uv run mypy 01-easy/parking-lot/python
```

## Running Go design patterns

All patterns use `cmd/demo/`:

```bash
cd 00-foundations/design-patterns/behavioral/strategy/golang
go run ./cmd/demo/
```

Or run all: `./tools/check.sh patterns`
