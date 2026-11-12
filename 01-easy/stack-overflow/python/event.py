from __future__ import annotations

from dataclasses import dataclass
from typing import TYPE_CHECKING

from enums import EventType
from user import User

if TYPE_CHECKING:
    from content import Post


@dataclass(slots=True)
class Event:
    type: EventType
    actor: User
    target_post: Post

    def get_type(self) -> EventType:
        return self.type

    def get_actor(self) -> User:
        return self.actor

    def get_target_post(self) -> Post:
        return self.target_post
