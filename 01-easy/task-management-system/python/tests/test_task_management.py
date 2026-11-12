from datetime import date

from task_management_system import TaskManagementSystem
from task_priority import TaskPriority
from task_sort_strategy import SortByDueDate
from task_status import TaskStatus


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


def test_start_and_complete_task() -> None:
    system = TaskManagementSystem()
    user = system.create_user("Bob", "bob@example.com")
    task = system.create_task(
        "Ship feature",
        "desc",
        date.today(),
        TaskPriority.MEDIUM,
        user.id,
    )
    task.start_progress()
    assert task.get_status() == TaskStatus.IN_PROGRESS
    task.complete_task()
    assert task.get_status() == TaskStatus.DONE
