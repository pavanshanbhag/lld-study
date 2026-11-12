# Task Management System — Go (modernized)

Idiomatic Go reference implementation:

- `NewTaskManager()` constructor (no singleton)
- Mutating ops return `error` (`CreateTask`, `UpdateTask`, `DeleteTask`, …)
- Exported fields instead of redundant `Get*` getters
- `TaskHistory` (no `Get` prefix); search uses `strings.Contains`
- `log/slog` in demo
- Behavior tests + `-race`

## Run

```bash
# from repo root — uncomment import in main.go, then:
go run .

# tests
go test -race ./01-easy/task-management-system/golang/...
```

## API highlights

| Old | New |
|-----|-----|
| `GetId()` / `GetTitle()` | exported fields `Id` / `Title` |
| `GetTaskHistory()` | `TaskHistory()` |
| silent update/delete miss | `error` returns |
| prefix-only `contains` | `strings.Contains` |
