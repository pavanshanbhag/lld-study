# Design Patterns (Go + Python)

Implementations organized by category. Each pattern has `golang/` and/or `python/` subfolders.

```
design-patterns/
├── creational/     singleton, factory, builder, prototype, abstractfactory
├── structural/     adapter, bridge, composite, decorator, facade, flyweight, proxy
└── behavioral/     strategy, observer, state, command, and more
```

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

From the pattern's `golang/` folder (most have their own `go.mod` and `main.go`):

```bash
cd 00-foundations/design-patterns/behavioral/strategy/golang
go run .
```

Patterns without a local `go.mod` can still be read as reference code.

---

## Creational Patterns

| Pattern | Python | Go | Demo |
|---------|--------|-----|------|
| [Singleton](creational/singleton/) | ✓ | ✓ | `python3 -m python.singleton_demo` |
| [Factory Method](creational/factory/) | ✓ | ✓ | `python3 -m python.factory_method_demo` |
| [Builder](creational/builder/) | ✓ | ✓ | `python3 -m python.http_app_builder` |
| [Prototype](creational/prototype/) | ✓ | ✓ | `python3 -m python.game` |
| [Abstract Factory](creational/abstractfactory/) | ✓ | — | `python3 -m python.shoe_manufacture` |

## Structural Patterns

| Pattern | Python | Go | Demo |
|---------|--------|-----|------|
| [Adapter](structural/adapter/) | ✓ | ✓ | `python3 -m python.ecommerce_app` |
| [Bridge](structural/bridge/) | ✓ | ✓ | `python3 -m python.bridge_demo` |
| [Composite](structural/composite/) | ✓ | ✓ | `python3 -m python.organization.composite_demo` |
| [Decorator](structural/decorator/) | ✓ | ✓ | `python3 -m python.decorator_demo` |
| [Facade](structural/facade/) | ✓ | ✓ | `python3 -m python.deployment_app_facade` |
| [Flyweight](structural/flyweight/) | ✓ | ✓ | `python3 -m python.flyweight_demo` |
| [Proxy](structural/proxy/) | ✓ | ✓ | `python3 -m python.image_gallery_app_v2` |

## Behavioral Patterns

| Pattern | Python | Go | Demo |
|---------|--------|-----|------|
| [Chain of Responsibility](behavioral/chainofresponsibility/) | ✓ | ✓ | `python3 -m python.main` |
| [Command](behavioral/command/) | ✓ | — | `python3 -m python.command_pattern_demo` |
| [Iterator](behavioral/iterator/) | ✓ | ✓ | `python3 -m python.iterator_demo` |
| [Mediator](behavioral/mediator/) | ✓ | ✓ | (see `python/mediator_app.py`) |
| [Memento](behavioral/memento/) | ✓ | ✓ | `python3 -m python.text_editor_undo_v2` |
| [Observer](behavioral/observer/) | ✓ | ✓ | `python3 -m python.fitness_app_observer_demo` |
| [State](behavioral/state/) | ✓ | ✓ | `python3 -m python.vending_machine_app` |
| [Strategy](behavioral/strategy/) | ✓ | ✓ | `python3 -m python.ecommerce_app_demo` |
| [Template Method](behavioral/templatemethod/) | ✓ | ✓ | `python3 -m python.report_app_template_method` |
| [Visitor](behavioral/visitor/) | ✓ | — | (stub — no demo yet) |

> Run all commands from the pattern directory shown in the first column link, e.g. `cd behavioral/strategy` before `python3 -m python.ecommerce_app_demo`.

## Pattern layout

```
<category>/<pattern>/
├── golang/          # Go implementation (+ go.mod, main.go where applicable)
└── python/          # Python implementation (__init__.py + demo modules)
    ├── __init__.py
    ├── ...
    └── *_demo.py
```

## All-patterns runner

`all_patterns_demo.py` in this directory was written for the old flat layout and may not work without path adjustments. Prefer running individual pattern demos above.

## Design principles

- **Open/Closed** — open for extension, closed for modification
- **Dependency Inversion** — depend on abstractions, not concretions
- **Single Responsibility** — one reason to change per class
- **Composition over Inheritance** — favor object composition
- **Encapsulation** — hide implementation behind interfaces
