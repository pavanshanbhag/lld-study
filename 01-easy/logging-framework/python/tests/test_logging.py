import time
from pathlib import Path

from log_appender import ConsoleAppender, FileAppender
from log_level import LogLevel
from log_manager import LogManager


def test_log_manager_hierarchy() -> None:
    manager = LogManager()
    root = manager.get_root_logger()
    root.set_level(LogLevel.INFO)
    root.add_appender(ConsoleAppender())

    logger = manager.get_logger("com.example.Main")
    logger.info("hello")
    manager.shutdown()
    time.sleep(0.1)


def test_file_appender_writes(tmp_path: Path) -> None:
    path = tmp_path / "app.log"
    manager = LogManager()
    root = manager.get_root_logger()
    root.set_level(LogLevel.DEBUG)
    root.add_appender(FileAppender(str(path)))
    manager.get_logger("app").info("persisted")
    time.sleep(0.2)
    manager.shutdown()
    assert "persisted" in path.read_text()
