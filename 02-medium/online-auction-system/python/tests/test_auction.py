from decimal import Decimal
from datetime import datetime, timedelta

from auction_service import AuctionService


def test_auction_service_constructor() -> None:
    service = AuctionService()
    user = service.create_user("Alice")
    assert service.get_user(user.get_id()) is user
    service.shutdown()


def test_place_bid_updates_highest() -> None:
    service = AuctionService()
    bidder = service.create_user("Bidder")
    end = datetime.now() + timedelta(hours=1)
    auction = service.create_auction("Watch", "Vintage", Decimal("100"), end)

    service.place_bid(auction.get_id(), bidder.get_id(), Decimal("150"))
    highest = auction.get_highest_bid()
    assert highest is not None
    assert highest.get_amount() == Decimal("150")
    assert highest.get_bidder() is bidder
    service.shutdown()
