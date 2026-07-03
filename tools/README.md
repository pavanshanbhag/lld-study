# Tools

## Running Go LLD demos

From the **repo root**:

1. Open `main.go`
2. Uncomment the import and `Run()` call for the topic you want, e.g.:

```go
import (
    parkinglot "lld-study/01-easy/parking-lot/golang"
)

func main() {
    parkinglot.Run()
}
```

3. Run: `go run .`

## Running Python LLD demos

From a topic folder:

```bash
cd 01-easy/parking-lot/python
python parking_lot_demo.py
```

Each project has a `*_demo.py` file — check the local README if the name differs.

## Running Go design patterns

Most patterns have their own `go.mod` and `main.go`:

```bash
cd 00-foundations/design-patterns/behavioral/strategy/golang
go run .
```
