from user import User
from typing import Optional
from threading import Lock
from commentable_entity import Post


class UserRepository:
    _lock = Lock()

    def __init__(self):
        self.users = {}

    def save(self, user: User):
        self.users[user.get_id()] = user

    def find_by_id(self, user_id: str) -> Optional[User]:
        return self.users.get(user_id)


class PostRepository:
    _lock = Lock()

    def __init__(self):
        self.posts = {}

    def save(self, post: Post):
        self.posts[post.get_id()] = post

    def find_by_id(self, post_id: str) -> Optional[Post]:
        return self.posts.get(post_id)
