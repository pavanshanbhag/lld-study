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

**Singletons removed on purpose.** Use `NewService()` / constructors. Learn singleton only in:

```
00-foundations/design-patterns/creational/singleton/
```

## Class diagrams

| File | Purpose |
|------|---------|
| `class-diagram.png` | Upstream UML (kept) |
| `class-diagram.mmd` | Mermaid matching modernized Go types |

All 33 LLD topics have both. Do not delete PNGs.

## Modernization status

| Tier | Topics | Status |
|------|--------|--------|
| **Reference** | parking-lot | Full idiomatic + ruff + mypy strict |
| **01-easy** | 7 | Parking-lot pattern: slog demos, errors, Get→idiomatic, dataclasses/StrEnum/logging, READMEs, tests |
| **02-medium** | 15 | Same pattern |
| **03-hard** | 11 | Same pattern |
| **Design patterns** | 19 | Go `cmd/demo/` + Go 1.25; catalog README accurate |
| **04-bonus** | voting-system | Constructor-based + Go/Python tests |

## Go checklist (applied across LLD)

- [x] `NewService()` constructor (no `Get*` singleton wrappers)
- [x] `error` from fallible operations (where applicable)
- [x] Idiomatic names (`LicensePlate()` not `GetLicensePlate()`)
- [x] `log/slog` in demos
- [x] `*_test.go` + `go test -race`
- [x] Short `golang/README.md` per topic

## Python checklist (applied across LLD)

- [x] Constructor, not `get_instance()`
- [x] Domain `logging` (demos may still print UX)
- [x] `@dataclass` / `StrEnum` where models are simple
- [x] Tests under `tests/`
- [x] Short `python/README.md` per topic
- [ ] mypy strict — parking-lot only (expand later)
- [ ] Single-module consolidation — parking-lot only

## Next steps (optional)

1. Expand mypy strict beyond parking-lot
2. Implement `04-bonus/voting-system`
3. Optional: consolidate multi-file Python topics like parking-lot
