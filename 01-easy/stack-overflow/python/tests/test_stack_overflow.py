from stack_overflow_service import StackOverflowService
from tag import Tag


def test_create_user_and_question() -> None:
    service = StackOverflowService()
    user = service.create_user("alice")
    question = service.post_question(user.get_id(), "How?", "Help", {Tag("go")})
    assert question.get_title() == "How?"
