from ride_sharing_service import RideSharingService


def test_ride_sharing_service() -> None:
    service = RideSharingService()
    rider = service.register_rider("Alice", "555-0100")

    assert rider.get_name() == "Alice"
