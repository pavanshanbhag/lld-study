# Modernization Guide

## Toolchain

| Tool | Version | Config |
|------|---------|--------|
| Go | 1.25 | Root `go.mod` + each design-pattern `golang/go.mod` |
| Python | 3.14 | `.python-version`, `pyproject.toml`, [uv](https://docs.astral.sh/uv/) |

```bash
uv sync
./tools/check.sh all        # Go + parking lint + Python + patterns + build
./tools/check.sh test-go    # all LLD Go (incl. 04-bonus/voting-system)
./tools/check.sh test-py    # all LLD Python (incl. voting)
./tools/check.sh easy       # 01-easy Go + Python
./tools/check.sh medium     # 02-medium Go + Python
./tools/check.sh hard       # 03-hard Go + Python
./tools/check.sh patterns   # 19 design-pattern Go demos
./tools/check.sh parking    # parking-lot reference (lint + strict mypy)
```

Full command list: [tools/README.md](../tools/README.md).

## Singleton: what changed and why

### LLD solutions (`01-easy` … `03-hard` + `04-bonus`)

**Singletons removed on purpose.** Use `NewService()` / constructors. Learn singleton only in:

```
00-foundations/design-patterns/creational/singleton/
```

## Class diagrams

| File | Purpose |
|------|---------|
| `class-diagram.png` | Upstream UML (kept) |
| `class-diagram.mmd` | Mermaid matching modernized Go types |

All **34** topic folders (7 easy + 15 medium + 11 hard + voting) have both. Do not delete PNGs.

## Modernization status

| Tier | Topics | Status |
|------|--------|--------|
| **Reference** | parking-lot | Full idiomatic + ruff + mypy strict |
| **01-easy** | 7 | Go+Python: constructors, slog/tests, short READMEs |
| **02-medium** | 15 | Same baseline |
| **03-hard** | 11 | Same baseline |
| **Design patterns** | 19 Go demos | `cmd/demo/` + Go 1.25 (`check.sh patterns`) |
| **04-bonus** | voting-system | Constructor-based + Go/Python tests |

## Go checklist (applied across LLD)

- [x] `NewService()` constructor (no singleton `Get*` / `sync.Once`)
- [x] `error` from fallible operations (where applicable)
- [x] Drop redundant `GetStatus()`; prefer `Status()` / idiomatic names
- [x] `log/slog` in demos
- [x] `*_test.go` + `go test -race` (incl. voting bonus)
- [x] Short `golang/README.md` per topic
- [ ] Mass `GetFoo()`→`Foo()` rename on every remaining Java-style getter (optional)

## Python checklist (applied across LLD)

- [x] Constructor, not `get_instance()`
- [x] Tests under `tests/` (run via `./tools/check.sh test-py`)
- [x] Short `python/README.md` per topic (run/demo notes)
- [ ] Domain `logging` / `@dataclass` / `StrEnum` — parking-lot (+ a few easy); expand later
- [ ] mypy strict — parking-lot only
- [ ] Single-module consolidation — parking-lot only

## Next steps (optional)

1. Mass `GetFoo()`→`Foo()` rename on remaining Java-style getters
2. Expand Python dataclasses/StrEnum/logging beyond parking-lot
3. Expand mypy strict beyond parking-lot
4. Optional: consolidate multi-file Python topics like parking-lot
