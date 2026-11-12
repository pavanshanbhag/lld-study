from linkedin_system import LinkedInSystem
from member import Member


def test_linkedin_constructor() -> None:
    system = LinkedInSystem()
    member = Member.Builder("Alice", "alice@example.com").build()
    system.register_member(member)
    assert system.get_member("Alice") is member


def test_connection_request_flow() -> None:
    system = LinkedInSystem()
    alice = Member.Builder("Alice", "alice@example.com").build()
    bob = Member.Builder("Bob", "bob@example.com").build()
    system.register_member(alice)
    system.register_member(bob)

    request_id = system.send_connection_request(alice, bob)
    system.accept_connection_request(request_id)
    assert bob in alice.get_connections()
    assert alice in bob.get_connections()
