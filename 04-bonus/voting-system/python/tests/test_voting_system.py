import pytest

from voting_system import (
    AlreadyVotedError,
    InvalidStateError,
    VotingSystem,
)


def test_happy_path() -> None:
    vs = VotingSystem()
    vs.create_election("E1", "Mayor")
    vs.add_candidate("E1", "C1", "Alice")
    vs.add_candidate("E1", "C2", "Bob")
    vs.register_voter("V1", "A")
    vs.register_voter("V2", "B")
    vs.register_voter("V3", "C")
    vs.open_election("E1")
    vs.cast_vote("E1", "V1", "C1")
    vs.cast_vote("E1", "V2", "C1")
    vs.cast_vote("E1", "V3", "C2")
    vs.close_election("E1")
    r = vs.results("E1")
    assert r.winner_id == "C1"
    assert r.total_votes == 3


def test_duplicate_vote() -> None:
    vs = VotingSystem()
    vs.create_election("E1", "Mayor")
    vs.add_candidate("E1", "C1", "Alice")
    vs.register_voter("V1", "A")
    vs.open_election("E1")
    vs.cast_vote("E1", "V1", "C1")
    with pytest.raises(AlreadyVotedError):
        vs.cast_vote("E1", "V1", "C1")


def test_results_before_close() -> None:
    vs = VotingSystem()
    vs.create_election("E1", "Mayor")
    vs.add_candidate("E1", "C1", "Alice")
    vs.open_election("E1")
    with pytest.raises(InvalidStateError):
        vs.results("E1")
