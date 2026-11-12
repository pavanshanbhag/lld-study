# Tools

## Toolchain

- **Go 1.25** — module at repo root (`go.mod`)
- **Python 3.14** — managed with `uv` (`pyproject.toml`, `.python-version`)

```bash
uv sync
./tools/check.sh all        # Go tests + ruff/mypy (parking) + Python tests + patterns + build
./tools/check.sh easy       # 01-easy Go + Python
./tools/check.sh medium     # 02-medium Go + Python
./tools/check.sh hard       # 03-hard Go + Python
./tools/check.sh test-go    # all LLD Go pkgs (easy/medium/hard + voting)
./tools/check.sh test-py    # all LLD Python tests (incl. voting)
./tools/check.sh patterns   # 19 design-pattern Go demos
./tools/check.sh parking    # parking-lot reference (race + ruff + mypy + pytest)
```

Other: `easy-py` / `medium-py` / `hard-py` / `lru` / `lint-py` / `test` / `sync`.

## Running Go LLD demos

From the **repo root**:

1. Open `main.go`
2. Uncomment the import and `Run()` call for the topic you want
3. Run: `go run .`

### Go tests

```bash
go test -race ./01-easy/parking-lot/golang/
./tools/check.sh test-go
```

## Running Python LLD demos

Always from **repo root** (so imports resolve under that topic):

```bash
uv run python 01-easy/parking-lot/python/parking_lot_demo.py
uv run python 02-medium/atm/python/atm_demo.py
uv run python 03-hard/splitwise/python/splitwise_demo.py
```

### Python tests

Topics share common module names (`user.py`, `enums.py`, …). **Isolate `PYTHONPATH` per topic** — do not run a bare multi-topic `pytest` from root.

```bash
PYTHONPATH=01-easy/parking-lot/python uv run pytest 01-easy/parking-lot/python/tests -q
./tools/check.sh test-py
```

### Strict lint (parking-lot reference only)

```bash
uv run ruff check 01-easy/parking-lot/python
uv run mypy 01-easy/parking-lot/python
```

## Running Go design patterns

Nineteen patterns ship `golang/cmd/demo/`:

```bash
cd 00-foundations/design-patterns/behavioral/strategy/golang
go run ./cmd/demo/
```

Or: `./tools/check.sh patterns`

Folders without a Go `cmd/demo/` (Abstract Factory, Command, Visitor) are Python-first / incomplete — not in `check.sh patterns`.
