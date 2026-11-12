from __future__ import annotations

from dataclasses import dataclass, field
from typing import Generic, TypeVar

K = TypeVar("K")
V = TypeVar("V")


@dataclass(slots=True)
class Node(Generic[K, V]):
    key: K
    value: V
    prev: Node[K, V] | None = field(default=None, repr=False)
    next: Node[K, V] | None = field(default=None, repr=False)
