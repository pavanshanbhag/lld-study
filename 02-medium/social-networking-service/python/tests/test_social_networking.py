from social_network_facade import SocialNetworkFacade


def test_social_network_facade_constructor() -> None:
    network = SocialNetworkFacade()
    user = network.create_user("Alice", "alice@example.com")
    post = network.create_post(user.get_id(), "Hello from Alice!")
    assert post.get_author() is user
