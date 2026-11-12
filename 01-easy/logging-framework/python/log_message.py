from __future__ import annotations

import threading
from dataclasses import dataclass, field
from datetime import datetime

from log_level import LogLevel


@dataclass(slots=True)
class LogMessage:
    level: LogLevel
    logger_name: str
    message: str
    timestamp: datetime = field(default_factory=datetime.now)
    thread_name: str = field(default_factory=lambda: threading.current_thread().name)

    def get_timestamp(self) -> datetime:
        return self.timestamp

    def get_level(self) -> LogLevel:
        return self.level

    def get_logger_name(self) -> str:
        return self.logger_name

    def get_thread_name(self) -> str:
        return self.thread_name

    def get_message(self) -> str:
        return self.message
