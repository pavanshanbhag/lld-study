# Design Patterns (Go + Python)

Implementations organized by category. Each pattern has `golang/` and/or `python/` subfolders.

```
design-patterns/
├── creational/   singleton, factory, builder, prototype, abstractfactory
├── structural/   adapter, bridge, composite, decorator, facade, flyweight, proxy
└── behavioral/   chainofresponsibility, command, iterator, mediator, memento,
                  observer, state, strategy, templatemethod, visitor
```

**Counts:** 19 Go patterns (`cmd/demo/`) · Python for all listed patterns (visitor is a stub).

## Running Python demos

Most Python packages use relative imports. Run from the **pattern directory** (parent of `python/`):

```bash
cd 00-foundations/design-patterns/<category>/<pattern>
python3 -m python.<demo_module>
```

Examples:

```bash
# Strategy
cd 00-foundations/design-patterns/behavioral/strategy
python3 -m python.ecommerce_app_demo

# Adapter
cd 00-foundations/design-patterns/structural/adapter
python3 -m python.ecommerce_app

# Factory
cd 00-foundations/design-patterns/creational/factory
python3 -m python.factory_method_demo

# Singleton
cd 00-foundations/design-patterns/creational/singleton
python3 -m python.singleton_demo

# Chain of Responsibility
cd 00-foundations/design-patterns/behavioral/chainofresponsibility
python3 -m python.main
```

## Running Go demos

Requires **Go 1.25** (same as repo root `go.mod`). Each pattern has its own `go.mod` under `golang/`.

From the pattern's **`golang/`** directory:

```bash
cd 00-foundations/design-patterns/behavioral/strategy/golang
go run ./cmd/demo/
```

All **19** Go patterns use the same entrypoint: `go run ./cmd/demo/`.

Run every demo: `./tools/check.sh patterns`

---

## Creational Patterns

| Pattern | Python | Go | Python demo | Go demo |
|---------|--------|-----|-------------|---------|
| [Singleton](creational/singleton/) | ✓ | ✓ | `python3 -m python.singleton_demo` | `go run ./cmd/demo/` |
| [Factory Method](creational/factory/) | ✓ | ✓ | `python3 -m python.factory_method_demo` | `go run ./cmd/demo/` |
| [Builder](creational/builder/) | ✓ | ✓ | `python3 -m python.http_app_builder` | `go run ./cmd/demo/` |
| [Prototype](creational/prototype/) | ✓ | ✓ | `python3 -m python.game` | `go run ./cmd/demo/` |
| [Abstract Factory](creational/abstractfactory/) | ✓ | — | `python3 -m python.shoe_manufacture` | — |

## Structural Patterns

| Pattern | Python | Go | Python demo | Go demo |
|---------|--------|-----|-------------|---------|
| [Adapter](structural/adapter/) | ✓ | ✓ | `python3 -m python.ecommerce_app` | `go run ./cmd/demo/` |
| [Bridge](structural/bridge/) | ✓ | ✓ | `python3 -m python.bridge_demo` | `go run ./cmd/demo/` |
| [Composite](structural/composite/) | ✓ | ✓ | `python3 -m python.organization.composite_demo` | `go run ./cmd/demo/` |
| [Decorator](structural/decorator/) | ✓ | ✓ | `python3 -m python.decorator_demo` | `go run ./cmd/demo/` |
| [Facade](structural/facade/) | ✓ | ✓ | `python3 -m python.deployment_app_facade` | `go run ./cmd/demo/` |
| [Flyweight](structural/flyweight/) | ✓ | ✓ | `python3 -m python.flyweight_demo` | `go run ./cmd/demo/` |
| [Proxy](structural/proxy/) | ✓ | ✓ | `python3 -m python.image_gallery_app_v2` | `go run ./cmd/demo/` |

## Behavioral Patterns

| Pattern | Python | Go | Python demo | Go demo |
|---------|--------|-----|-------------|---------|
| [Chain of Responsibility](behavioral/chainofresponsibility/) | ✓ | ✓ | `python3 -m python.main` | `go run ./cmd/demo/` |
| [Command](behavioral/command/) | ✓ | — | `python3 -m python.command_pattern_demo` | — |
| [Iterator](behavioral/iterator/) | ✓ | ✓ | `python3 -m python.iterator_demo` | `go run ./cmd/demo/` |
| [Mediator](behavioral/mediator/) | ✓ | ✓ | `python3 -m python.mediator_app` | `go run ./cmd/demo/` |
| [Memento](behavioral/memento/) | ✓ | ✓ | `python3 -m python.text_editor_undo_v2` | `go run ./cmd/demo/` |
| [Observer](behavioral/observer/) | ✓ | ✓ | `python3 -m python.fitness_app_observer_demo` | `go run ./cmd/demo/` |
| [State](behavioral/state/) | ✓ | ✓ | `python3 -m python.vending_machine_app` | `go run ./cmd/demo/` |
| [Strategy](behavioral/strategy/) | ✓ | ✓ | `python3 -m python.ecommerce_app_demo` | `go run ./cmd/demo/` |
| [Template Method](behavioral/templatemethod/) | ✓ | ✓ | `python3 -m python.report_app_template_method` | `go run ./cmd/demo/` |
| [Visitor](behavioral/visitor/) | ✓ | — | stub (`python/__init__.py` only) | — |

> **Python:** run from the pattern directory (e.g. `cd behavioral/strategy`).
> **Go:** run from that pattern's `golang/` directory (e.g. `cd behavioral/strategy/golang`).

## Pattern layout

```
<category>/<pattern>/
├── golang/          # Go 1.25 module + cmd/demo/main.go (19 patterns)
└── python/          # Python package (__init__.py + demo modules)
    ├── __init__.py
    ├── ...
    └── *_demo.py
```

## All-patterns runner

`all_patterns_demo.py` in this directory was written for the old flat layout and may not work without path adjustments. Prefer individual demos above, or `./tools/check.sh patterns` for all Go demos.

## Design principles

- **Open/Closed** — open for extension, closed for modification
- **Dependency Inversion** — depend on abstractions, not concretions
- **Single Responsibility** — one reason to change per class
- **Composition over Inheritance** — favor object composition
- **Encapsulation** — hide implementation behind interfaces
