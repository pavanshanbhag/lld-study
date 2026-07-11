from typing import List

from commentable_entity import Comment, Post
from post_observer import PostObserver
from repository import PostRepository
from user import User


class PostService:
    def __init__(self, post_repository: PostRepository):
        self.post_repository = post_repository
        self.observers: List[PostObserver] = []

    def add_observer(self, observer: PostObserver):
        self.observers.append(observer)

    def create_post(self, author: User, content: str) -> Post:
        post = Post(author, content)
        self.post_repository.save(post)
        author.add_post(post)
        for observer in self.observers:
            observer.on_post_created(post)
        return post

    def like_post(self, user: User, post_id: str):
        post = self.post_repository.find_by_id(post_id)
        post.add_like(user)
        for observer in self.observers:
            observer.on_like(post, user)

    def add_comment(self, author: User, commentable_id: str, content: str):
        comment = Comment(author, content)
        post = self.post_repository.find_by_id(commentable_id)
        post.add_comment(comment)
        for observer in self.observers:
            observer.on_comment(post, comment)
