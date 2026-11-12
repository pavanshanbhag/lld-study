from enums import VoteType
from stack_overflow_service import StackOverflowService
from tag import Tag


def test_create_user_and_question() -> None:
    service = StackOverflowService()
    user = service.create_user("alice")
    question = service.post_question(user.get_id(), "How?", "Help", {Tag("go")})
    assert question.get_title() == "How?"


def test_vote_updates_count() -> None:
    service = StackOverflowService()
    alice = service.create_user("alice")
    bob = service.create_user("bob")
    question = service.post_question(alice.get_id(), "Q?", "body", {Tag("go")})
    service.vote_on_post(bob.get_id(), question.get_id(), VoteType.UPVOTE)
    assert question.get_vote_count() == 1
