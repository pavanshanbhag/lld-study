from auction_service import AuctionService


def test_auction_service_constructor() -> None:
    service = AuctionService()
    user = service.create_user("Alice")
    assert service.get_user(user.get_id()) is user
    service.shutdown()
