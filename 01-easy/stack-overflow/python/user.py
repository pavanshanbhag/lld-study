from __future__ import annotations

import threading
import uuid
from dataclasses import dataclass, field


@dataclass
class User:
    name: str
    id: str = field(default_factory=lambda: str(uuid.uuid4()))
    reputation: int = 0
    _lock: threading.Lock = field(default_factory=threading.Lock, repr=False)

    def update_reputation(self, change: int) -> None:
        with self._lock:
            self.reputation += change

    def get_id(self) -> str:
        return self.id

    def get_name(self) -> str:
        return self.name

    def get_reputation(self) -> int:
        with self._lock:
            return self.reputation
