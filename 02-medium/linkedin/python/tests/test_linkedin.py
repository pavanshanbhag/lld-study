from linkedin_system import LinkedInSystem
from member import Member


def test_linkedin_constructor() -> None:
    system = LinkedInSystem()
    member = Member.Builder("Alice", "alice@example.com").build()
    system.register_member(member)
    assert system.get_member("Alice") is member
