from splitwise_service import SplitwiseService


def test_splitwise_service() -> None:
    service = SplitwiseService()
    alice = service.add_user("Alice", "alice@example.com")
    bob = service.add_user("Bob", "bob@example.com")

    group = service.add_group("Trip", [alice, bob])
    assert service.get_user(alice.get_id()) is not None
    assert service.get_group(group.get_id()) is not None
