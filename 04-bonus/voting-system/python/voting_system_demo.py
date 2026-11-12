import logging

from voting_system import AlreadyVotedError, VotingSystem

logging.basicConfig(level=logging.INFO, format="%(message)s")
log = logging.getLogger("demo")


def main() -> None:
    vs = VotingSystem()
    vs.create_election("E1", "City Mayor 2026")
    vs.add_candidate("E1", "C1", "Alice")
    vs.add_candidate("E1", "C2", "Bob")
    vs.register_voter("V1", "Nina")
    vs.register_voter("V2", "Omar")
    vs.register_voter("V3", "Priya")
    vs.open_election("E1")
    vs.cast_vote("E1", "V1", "C1")
    vs.cast_vote("E1", "V2", "C2")
    vs.cast_vote("E1", "V3", "C1")
    try:
        vs.cast_vote("E1", "V1", "C2")
    except AlreadyVotedError as e:
        log.info("duplicate vote rejected: %s", e)
    vs.close_election("E1")
    result = vs.results("E1")
    log.info("winner=%s total=%s counts=%s", result.winner_name, result.total_votes, result.counts)


if __name__ == "__main__":
    main()
