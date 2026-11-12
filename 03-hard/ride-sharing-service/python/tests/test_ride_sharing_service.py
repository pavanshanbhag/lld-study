from enums import RideType
from location import Location
from ride_sharing_service import RideSharingService
from vehicle import Vehicle


def test_ride_sharing_service() -> None:
    service = RideSharingService()
    rider = service.register_rider("Alice", "555-0100")

    assert rider.get_name() == "Alice"


def test_register_driver() -> None:
    service = RideSharingService()
    vehicle = Vehicle("KA-01", "Sedan", RideType.SEDAN)
    driver = service.register_driver("Bob", "555-0200", vehicle, Location(12.0, 77.0))
    assert driver.get_name() == "Bob"
