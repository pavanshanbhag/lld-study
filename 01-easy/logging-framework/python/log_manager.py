from typing import TYPE_CHECKING

from async_log_processor import AsyncLogProcessor
from logger import Logger

if TYPE_CHECKING:
    pass


class LogManager:
    def __init__(self) -> None:
        self.loggers: dict[str, Logger] = {}
        self.processor = AsyncLogProcessor()
        self.root_logger = Logger("root", None, self)
        self.loggers["root"] = self.root_logger

    def get_logger(self, name: str) -> Logger:
        if name not in self.loggers:
            self.loggers[name] = self._create_logger(name)
        return self.loggers[name]

    def _create_logger(self, name: str) -> Logger:
        if name == "root":
            return self.root_logger

        last_dot = name.rfind(".")
        parent_name = "root" if last_dot == -1 else name[:last_dot]
        parent = self.get_logger(parent_name)
        return Logger(name, parent, self)

    def get_root_logger(self) -> Logger:
        return self.root_logger

    def get_processor(self) -> AsyncLogProcessor:
        return self.processor

    def shutdown(self) -> None:
        self.processor.stop()

        all_appenders = set()
        for logger in self.loggers.values():
            for appender in logger.get_appenders():
                all_appenders.add(appender)

        for appender in all_appenders:
            appender.close()

        print("Logging framework shut down gracefully.")
