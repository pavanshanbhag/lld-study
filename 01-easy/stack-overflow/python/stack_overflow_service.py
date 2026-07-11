
from content import Answer, Post, Question
from enums import VoteType
from reputation_manager import ReputationManager
from search_strategy import SearchStrategy
from tag import Tag
from user import User


class StackOverflowService:
    def __init__(self):
        self.users: dict[str, User] = {}
        self.questions: dict[str, Question] = {}
        self.answers: dict[str, Answer] = {}
        self.reputation_manager = ReputationManager()

    def create_user(self, name: str) -> User:
        user = User(name)
        self.users[user.get_id()] = user
        return user

    def post_question(self, user_id: str, title: str, body: str, tags: set[Tag]) -> Question:
        author = self.users[user_id]
        question = Question(title, body, author, tags)
        question.add_observer(self.reputation_manager)
        self.questions[question.get_id()] = question
        return question

    def post_answer(self, user_id: str, question_id: str, body: str) -> Answer:
        author = self.users[user_id]
        question = self.questions[question_id]
        answer = Answer(body, author)
        answer.add_observer(self.reputation_manager)
        question.add_answer(answer)
        self.answers[answer.get_id()] = answer
        return answer

    def vote_on_post(self, user_id: str, post_id: str, vote_type: VoteType):
        user = self.users[user_id]
        post = self.find_post_by_id(post_id)
        post.vote(user, vote_type)

    def accept_answer(self, question_id: str, answer_id: str):
        question = self.questions[question_id]
        answer = self.answers[answer_id]
        question.accept_answer(answer)

    def search_questions(self, strategies: list[SearchStrategy]) -> list[Question]:
        results = list(self.questions.values())

        for strategy in strategies:
            results = strategy.filter(results)

        return results

    def get_user(self, user_id: str) -> User:
        return self.users[user_id]

    def find_post_by_id(self, post_id: str) -> Post:
        if post_id in self.questions:
            return self.questions[post_id]
        elif post_id in self.answers:
            return self.answers[post_id]
        
        raise KeyError("Post not found")