# LLD Study (Go + Python)

Personal study repo for **Low Level Design (LLD)** interview prep — reorganized from [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design) by AlgoMaster (Go and Python only).

## Layout

```
00-foundations/     OOP concepts + design patterns (Go & Python)
01-easy/            7 interview problems
02-medium/          15 interview problems
03-hard/            11 interview problems
04-bonus/           Extra solutions (no problem write-up)
tools/              How to run demos
assets/images/      Branding assets from upstream
```

Each topic folder contains everything for one study session:

```
01-easy/parking-lot/
├── problem.md          Requirements & design notes
├── class-diagram.png   UML diagram
├── golang/             Go solution
└── python/             Python solution
```

## Study order

### 1. Foundations (`00-foundations/`)

| Section | Contents |
|---------|----------|
| `oop/` | Classes, inheritance, polymorphism, composition, etc. |
| `design-patterns/creational/` | Singleton, Factory, Builder, Prototype, Abstract Factory |
| `design-patterns/structural/` | Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy |
| `design-patterns/behavioral/` | Strategy, Observer, State, Command, and more |

### 2. Easy (`01-easy/`)

Parking Lot · Stack Overflow · Vending Machine · Logging Framework · Traffic Signal · Coffee Vending Machine · Task Management

### 3. Medium (`02-medium/`)

ATM · LinkedIn · LRU Cache · Tic-Tac-Toe · Pub/Sub · Elevator · Car Rental · Online Auction · Hotel Management · Digital Wallet · Airline · Library · Social Network · Restaurant · Concert Tickets

### 4. Hard (`03-hard/`)

CricInfo · Splitwise · Chess · Snake & Ladder · Ride Sharing · Course Registration · Movie Tickets · Online Shopping · Stock Brokerage · Music Streaming · Food Delivery

### 5. Bonus (`04-bonus/`)

Voting System (solution only)

## Toolchain

| Tool | Version |
|------|---------|
| Go | **1.25** (`go.mod`) |
| Python | **3.14** (`.python-version`, managed with [uv](https://docs.astral.sh/uv/)) |

```bash
# One-time setup
uv sync

# Run full test suite (33 LLD topics + 19 pattern demos)
chmod +x tools/check.sh
./tools/check.sh all

# Tier-specific
./tools/check.sh easy      # 01-easy (7 topics)
./tools/check.sh medium    # 02-medium (15 topics)
./tools/check.sh hard      # 03-hard (11 topics)
./tools/check.sh patterns  # 19 Go design-pattern demos
```

**Modernization status:** All 33 LLD topics have constructor-based lifecycle, Go smoke tests (`go test -race`), and Python smoke tests. **parking-lot** remains the full idiomatic reference (ruff + mypy strict). See [docs/MODERNIZATION.md](docs/MODERNIZATION.md).

## How to run

See [tools/README.md](tools/README.md).

**Go (from repo root):** uncomment an import in `main.go`, then `go run .`

**Python (parking-lot pilot):** `uv run python 01-easy/parking-lot/python/parking_lot_demo.py`

**Other Python topics:** `cd <topic>/python && python3 <name>_demo.py`

## Attribution

Based on [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design) by [ashishps1](https://github.com/ashishps1) / [AlgoMaster.io](https://algomaster.io). Licensed under upstream terms — see [LICENSE](LICENSE).

This is a **personal fork**: trimmed to Go/Python, reorganized for study. Not affiliated with the original authors.
