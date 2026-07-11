from datetime import datetime, timedelta

from airline_management_system import AirlineManagementSystem
from flight import Flight


def test_airline_constructor_and_search() -> None:
    system = AirlineManagementSystem()
    departure = datetime.now() + timedelta(days=1)
    arrival = departure + timedelta(hours=2)
    flight = Flight("F001", "New York", "London", departure, arrival)
    system.add_flight(flight)

    results = system.search_flights("New York", "London", departure.date())
    assert len(results) == 1
    assert results[0].flight_number == "F001"
