from social_network_facade import SocialNetworkFacade


def test_social_network_facade_constructor() -> None:
    network = SocialNetworkFacade()
    user = network.create_user("Alice", "alice@example.com")
    post = network.create_post(user.get_id(), "Hello from Alice!")
    assert post.get_author() is user


def test_friend_and_feed() -> None:
    network = SocialNetworkFacade()
    alice = network.create_user("Alice", "alice@example.com")
    bob = network.create_user("Bob", "bob@example.com")
    network.add_friend(alice.get_id(), bob.get_id())
    post = network.create_post(bob.get_id(), "Bob's post")
    feed = network.get_news_feed(alice.get_id())
    assert post in feed
