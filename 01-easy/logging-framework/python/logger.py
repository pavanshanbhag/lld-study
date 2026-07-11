from typing import TYPE_CHECKING

from log_appender import LogAppender
from log_level import LogLevel
from log_message import LogMessage

if TYPE_CHECKING:
    from log_manager import LogManager


class Logger:
    def __init__(self, name: str, parent: Logger | None, manager: LogManager) -> None:
        self.name = name
        self.level: LogLevel | None = None
        self.parent = parent
        self.appenders: list[LogAppender] = []
        self.additivity = True
        self._manager = manager

    def add_appender(self, appender: LogAppender) -> None:
        self.appenders.append(appender)

    def get_appenders(self) -> list[LogAppender]:
        return self.appenders

    def set_level(self, min_level: LogLevel) -> None:
        self.level = min_level

    def set_additivity(self, additivity: bool) -> None:
        self.additivity = additivity

    def get_effective_level(self) -> LogLevel:
        logger: Logger | None = self
        while logger is not None:
            if logger.level is not None:
                return logger.level
            logger = logger.parent
        return LogLevel.DEBUG

    def log(self, message_level: LogLevel, message: str) -> None:
        if message_level.is_greater_or_equal(self.get_effective_level()):
            log_message = LogMessage(message_level, self.name, message)
            self._call_appenders(log_message)

    def _call_appenders(self, log_message: LogMessage) -> None:
        if self.appenders:
            self._manager.get_processor().process(log_message, self.appenders)

        if self.additivity and self.parent is not None:
            self.parent._call_appenders(log_message)

    def debug(self, message: str) -> None:
        self.log(LogLevel.DEBUG, message)

    def info(self, message: str) -> None:
        self.log(LogLevel.INFO, message)

    def warn(self, message: str) -> None:
        self.log(LogLevel.WARN, message)

    def error(self, message: str) -> None:
        self.log(LogLevel.ERROR, message)

    def fatal(self, message: str) -> None:
        self.log(LogLevel.FATAL, message)
