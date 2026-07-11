from abc import ABC, abstractmethod

from content import Question
from tag import Tag
from user import User


class SearchStrategy(ABC):
    @abstractmethod
    def filter(self, questions: list[Question]) -> list[Question]:
        pass

class KeywordSearchStrategy(SearchStrategy):
    def __init__(self, keyword: str):
        self.keyword = keyword.lower()

    def filter(self, questions: list[Question]) -> list[Question]:
        return [q for q in questions 
                if self.keyword in q.get_title().lower() or self.keyword in q.get_body().lower()]

class TagSearchStrategy(SearchStrategy):
    def __init__(self, tag: Tag):
        self.tag = tag

    def filter(self, questions: list[Question]) -> list[Question]:
        return [q for q in questions 
                if any(t.get_name().lower() == self.tag.get_name().lower() for t in q.get_tags())]

class UserSearchStrategy(SearchStrategy):
    def __init__(self, user: User):
        self.user = user

    def filter(self, questions: list[Question]) -> list[Question]:
        return [q for q in questions if q.get_author().get_id() == self.user.get_id()]