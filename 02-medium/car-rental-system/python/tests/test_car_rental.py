from car import Car
from rental_system import RentalSystem


def test_rental_system_constructor() -> None:
    rental_system = RentalSystem()
    rental_system.add_car(Car("Toyota", "Camry", 2022, "ABC123", 50.0))
    assert "ABC123" in rental_system.cars
