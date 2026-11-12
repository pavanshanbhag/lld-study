from enum import IntEnum


class LogLevel(IntEnum):
    DEBUG = 1
    INFO = 2
    WARN = 3
    ERROR = 4
    FATAL = 5

    def is_greater_or_equal(self, other: LogLevel) -> bool:
        return self.value >= other.value
