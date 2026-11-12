"""Voting system domain model."""

from __future__ import annotations

import logging
import threading
from dataclasses import dataclass, field
from enum import StrEnum

logger = logging.getLogger(__name__)


class ElectionStatus(StrEnum):
    CREATED = "Created"
    OPEN = "Open"
    CLOSED = "Closed"


class VotingError(Exception):
    """Base voting error."""


class NotFoundError(VotingError):
    pass


class AlreadyExistsError(VotingError):
    pass


class InvalidStateError(VotingError):
    pass


class AlreadyVotedError(VotingError):
    pass


class UnknownCandidateError(VotingError):
    pass


class UnknownVoterError(VotingError):
    pass


@dataclass(frozen=True, slots=True)
class Candidate:
    id: str
    name: str


@dataclass(frozen=True, slots=True)
class Voter:
    id: str
    name: str


@dataclass(frozen=True, slots=True)
class Ballot:
    voter_id: str
    candidate_id: str


@dataclass(slots=True)
class TallyResult:
    counts: dict[str, int]
    winner_id: str
    winner_name: str
    total_votes: int


class Election:
    def __init__(self, election_id: str, title: str) -> None:
        self.id = election_id
        self.title = title
        self._status = ElectionStatus.CREATED
        self._candidates: dict[str, Candidate] = {}
        self._ballots: list[Ballot] = []
        self._voted: set[str] = set()
        self._lock = threading.Lock()

    @property
    def status(self) -> ElectionStatus:
        with self._lock:
            return self._status

    def add_candidate(self, candidate: Candidate) -> None:
        with self._lock:
            if self._status is not ElectionStatus.CREATED:
                raise InvalidStateError("add candidates only before open")
            if candidate.id in self._candidates:
                raise AlreadyExistsError(candidate.id)
            self._candidates[candidate.id] = candidate

    def open(self) -> None:
        with self._lock:
            if self._status is not ElectionStatus.CREATED or not self._candidates:
                raise InvalidStateError("cannot open")
            self._status = ElectionStatus.OPEN

    def close(self) -> None:
        with self._lock:
            if self._status is not ElectionStatus.OPEN:
                raise InvalidStateError("cannot close")
            self._status = ElectionStatus.CLOSED

    def cast_vote(self, voter_id: str, candidate_id: str) -> None:
        with self._lock:
            if self._status is not ElectionStatus.OPEN:
                raise InvalidStateError("election not open")
            if candidate_id not in self._candidates:
                raise UnknownCandidateError(candidate_id)
            if voter_id in self._voted:
                raise AlreadyVotedError(voter_id)
            self._voted.add(voter_id)
            self._ballots.append(Ballot(voter_id, candidate_id))

    def ballots(self) -> list[Ballot]:
        with self._lock:
            return list(self._ballots)

    def candidates(self) -> list[Candidate]:
        with self._lock:
            return list(self._candidates.values())


class PluralityTally:
    def tally(self, election: Election) -> TallyResult:
        counts = {c.id: 0 for c in election.candidates()}
        for b in election.ballots():
            counts[b.candidate_id] += 1
        winner = max(election.candidates(), key=lambda c: counts[c.id])
        return TallyResult(
            counts=counts,
            winner_id=winner.id,
            winner_name=winner.name,
            total_votes=len(election.ballots()),
        )


class VotingSystem:
    def __init__(self) -> None:
        self._elections: dict[str, Election] = {}
        self._voters: dict[str, Voter] = {}
        self._tally = PluralityTally()
        self._lock = threading.RLock()

    def create_election(self, election_id: str, title: str) -> Election:
        with self._lock:
            if election_id in self._elections:
                raise AlreadyExistsError(election_id)
            e = Election(election_id, title)
            self._elections[election_id] = e
            return e

    def register_voter(self, voter_id: str, name: str) -> Voter:
        with self._lock:
            if voter_id in self._voters:
                raise AlreadyExistsError(voter_id)
            v = Voter(voter_id, name)
            self._voters[voter_id] = v
            return v

    def add_candidate(self, election_id: str, candidate_id: str, name: str) -> None:
        self._election(election_id).add_candidate(Candidate(candidate_id, name))

    def open_election(self, election_id: str) -> None:
        self._election(election_id).open()

    def close_election(self, election_id: str) -> None:
        self._election(election_id).close()

    def cast_vote(self, election_id: str, voter_id: str, candidate_id: str) -> None:
        with self._lock:
            if voter_id not in self._voters:
                raise UnknownVoterError(voter_id)
        self._election(election_id).cast_vote(voter_id, candidate_id)

    def results(self, election_id: str) -> TallyResult:
        e = self._election(election_id)
        if e.status is not ElectionStatus.CLOSED:
            raise InvalidStateError("results only after close")
        return self._tally.tally(e)

    def _election(self, election_id: str) -> Election:
        with self._lock:
            e = self._elections.get(election_id)
            if e is None:
                raise NotFoundError(election_id)
            return e
