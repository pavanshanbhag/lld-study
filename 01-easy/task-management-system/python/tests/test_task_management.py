from datetime import date

from task_management_system import TaskManagementSystem
from task_priority import TaskPriority
from task_sort_strategy import SortByDueDate


def test_create_task() -> None:
    system = TaskManagementSystem()
    user = system.create_user("Alice", "alice@example.com")
    task = system.create_task(
        "Write docs",
        "Modernization guide",
        date.today(),
        TaskPriority.HIGH,
        user.id,
    )
    results = system.search_tasks("docs", SortByDueDate())
    assert any(t.get_id() == task.get_id() for t in results)
