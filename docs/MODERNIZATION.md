# Modernization Guide

## Toolchain

| Tool | Version | Config |
|------|---------|--------|
| Go | 1.25 | Root `go.mod` + each design-pattern `golang/go.mod` |
| Python | 3.14 | `.python-version`, `pyproject.toml`, [uv](https://docs.astral.sh/uv/) |

```bash
uv sync
./tools/check.sh all       # full suite: Go + Python + patterns + build
./tools/check.sh easy      # 01-easy Go + Python
./tools/check.sh medium    # 02-medium Go + Python
./tools/check.sh hard      # 03-hard Go + Python
./tools/check.sh patterns  # all 19 design-pattern Go demos
./tools/check.sh parking   # parking-lot reference (lint + strict mypy)
```

## Singleton: what changed and why

### LLD solutions (01-easy … 03-hard)

**We removed singletons from modernized LLD code on purpose.**

| Before (upstream) | After (modernized) |
|-------------------|-------------------|
| `GetParkingLotInstance()` | `NewParkingLot()` |
| `ParkingLot.get_instance()` | `ParkingLot()` constructor |
| Global hidden state | Explicit dependency injection |

**Nothing is broken** for interview study — the **design** (classes, relationships, patterns) is unchanged. Only the **lifecycle** is better: you create the object, pass it around, test it in isolation.

Singleton is still available as a **pattern to learn** in:

```
00-foundations/design-patterns/creational/singleton/
```

Learn singleton **there**. Use constructors **in LLD problems**.

## Class diagrams (PNG files)

### Keep the existing PNGs — add Mermaid sources alongside

Each topic has `class-diagram.png` copied from upstream (interview reference). **Do not delete these.**

When modernized code diverges from upstream naming, add **`class-diagram.mmd`** next to the PNG:

| File | Purpose |
|------|---------|
| `class-diagram.png` | Original upstream UML (kept) |
| `class-diagram.mmd` | Mermaid source matching modernized Go/Python code |

**All 33 LLD topics** (01-easy through 03-hard) now have `class-diagram.mmd` reflecting modernized Go types. Original `class-diagram.png` files are unchanged.

Preview in VS Code (Mermaid extension), GitHub markdown, or export:

```bash
# optional, after installing @mermaid-js/mermaid-cli
mmdc -i class-diagram.mmd -o class-diagram-modern.png
```

**Not yet:** `04-bonus/voting-system` — no implementation; add `.mmd` when code lands.

## Modernization status

| Tier | Topics | Status |
|------|--------|--------|
| **Reference** | parking-lot | Full idiomatic (Go + Python, tests, ruff, mypy strict) |
| **01-easy** | 6 others | Singleton → constructor, smoke tests pass |
| **02-medium** | 15 | Singleton → constructor, Go + Python smoke tests |
| **03-hard** | 11 | Singleton → constructor, Go + Python smoke tests |
| **Design patterns** | 19 | All Go demos via `cmd/demo/`; Python demos unchanged |
| **04-bonus** | voting-system | Stub only (README placeholders; no code yet) |

## Go modernization checklist (per topic)

- [x] `NewService()` constructor (no deprecated `Get*` wrappers)
- [ ] Return `error` from fallible operations (where applicable)
- [ ] Idiomatic names (`LicensePlate()` not `GetLicensePlate()`)
- [ ] `log/slog` in demos, not domain `println`
- [x] `*_test.go` with table-driven or smoke tests
- [x] `go test -race ./…`

## Python modernization checklist (per topic)

- [ ] Single module or proper package with `__init__.py`
- [ ] `@dataclass`, `StrEnum`, `Protocol` (parking-lot reference only so far)
- [ ] `datetime` not millisecond ints
- [ ] `logging` in domain; prints only in demo
- [x] Constructor, not `get_instance()`
- [x] Tests under `tests/` (mypy strict: parking-lot only)

## Next steps (optional depth)

1. Full idiomatic refactor per topic (like parking-lot) — dataclasses, consolidated modules
2. Expand mypy strict coverage beyond parking-lot
3. Implement `04-bonus/voting-system` (Go + Python + tests + `class-diagram.mmd`)
