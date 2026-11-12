# LLD Study (Go + Python)

Personal study repo for **Low Level Design (LLD)** interview prep — reorganized from [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design) by AlgoMaster (Go and Python only).

## Layout

```
00-foundations/     OOP concepts + design patterns (Go & Python)
01-easy/            7 interview problems
02-medium/          15 interview problems
03-hard/            11 interview problems
04-bonus/           Extra solutions (voting-system)
tools/              check.sh + how to run demos
docs/               Modernization notes
assets/images/      Branding assets from upstream
```

Each interview topic folder contains one study session:

```
01-easy/parking-lot/
├── problem.md          Requirements & design notes
├── class-diagram.png   Upstream UML
├── class-diagram.mmd   Mermaid (modernized Go types)
├── golang/             Go solution + README + tests
└── python/             Python solution + README + tests
```

## Study order

### 1. Foundations (`00-foundations/`)

| Section | Contents |
|---------|----------|
| `oop/` | Classes, inheritance, polymorphism, composition, etc. |
| `design-patterns/creational/` | Singleton, Factory, Builder, Prototype (+ Python Abstract Factory) |
| `design-patterns/structural/` | Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy |
| `design-patterns/behavioral/` | Strategy, Observer, State, Command, and more |

**Go pattern demos:** 19 topics with `golang/cmd/demo/` (run via `./tools/check.sh patterns`). Command / Visitor / Abstract Factory are Python-oriented or incomplete on the Go side — see [design-patterns README](00-foundations/design-patterns/README.md).

### 2. Easy (`01-easy/`)

Parking Lot · Stack Overflow · Vending Machine · Logging Framework · Traffic Signal · Coffee Vending Machine · Task Management

### 3. Medium (`02-medium/`)

ATM · LinkedIn · LRU Cache · Tic-Tac-Toe · Pub/Sub · Elevator · Car Rental · Online Auction · Hotel Management · Digital Wallet · Airline · Library · Social Network · Restaurant · Concert Tickets

### 4. Hard (`03-hard/`)

CricInfo · Splitwise · Chess · Snake & Ladder · Ride Sharing · Course Registration · Movie Tickets · Online Shopping · Stock Brokerage · Music Streaming · Food Delivery

### 5. Bonus (`04-bonus/`)

Voting System (solution only)

## How to run

See **[tools/README.md](tools/README.md)** and **[docs/MODERNIZATION.md](docs/MODERNIZATION.md)**.

```bash
uv sync
./tools/check.sh all        # Go tests + Python tests + patterns + build
./tools/check.sh test-go    # all LLD Go packages (incl. voting)
./tools/check.sh test-py    # all LLD Python tests
```

**Go demo (from repo root):** uncomment import + `Run()` in `main.go`, then `go run .`

**Python demo (from repo root):**

```bash
uv run python 01-easy/parking-lot/python/parking_lot_demo.py
```

## Attribution

Based on [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design) by [ashishps1](https://github.com/ashishps1) / [AlgoMaster.io](https://algomaster.io). Licensed under upstream terms — see [LICENSE](LICENSE).

This is a **personal fork**: trimmed to Go/Python, reorganized for study. Not affiliated with the original authors.
