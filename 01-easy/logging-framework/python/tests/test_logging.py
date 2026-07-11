import time

from log_appender import ConsoleAppender
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
