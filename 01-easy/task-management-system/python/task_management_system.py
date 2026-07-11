from datetime import date

from task import Task
from task_list import TaskList
from task_observer import ActivityLogger
from task_priority import TaskPriority
from task_sort_strategy import TaskSortStrategy
from task_status import TaskStatus
from user import User


class TaskManagementSystem:
    def __init__(self) -> None:
        self._users: dict[str, User] = {}
        self._tasks: dict[str, Task] = {}
        self._task_lists: dict[str, TaskList] = {}

    def create_user(self, name: str, email: str) -> User:
        user = User(name, email)
        self._users[user.id] = user
        return user

    def create_task_list(self, list_name: str) -> TaskList:
        task_list = TaskList(list_name)
        self._task_lists[task_list.id] = task_list
        return task_list

    def create_task(
        self,
        title: str,
        description: str,
        due_date: date,
        priority: TaskPriority,
        created_by_user_id: str,
    ) -> Task:
        created_by = self._users.get(created_by_user_id)
        if created_by is None:
            raise ValueError("User not found.")

        task = (
            Task.TaskBuilder(title)
            .description(description)
            .due_date(due_date)
            .priority(priority)
            .created_by(created_by)
            .build()
        )

        task.add_observer(ActivityLogger())
        self._tasks[task.get_id()] = task
        return task

    def list_tasks_by_user(self, user_id: str) -> list[Task]:
        user = self._users.get(user_id)
        return [task for task in self._tasks.values() if task.get_assignee() == user]

    def list_tasks_by_status(self, status: TaskStatus) -> list[Task]:
        return [task for task in self._tasks.values() if task.get_status() == status]

    def delete_task(self, task_id: str) -> None:
        if task_id in self._tasks:
            del self._tasks[task_id]

    def search_tasks(self, keyword: str, sorting_strategy: TaskSortStrategy) -> list[Task]:
        matching_tasks = [
            task
            for task in self._tasks.values()
            if keyword in task.get_title() or keyword in task.get_description()
        ]
        sorting_strategy.sort(matching_tasks)
        return matching_tasks
