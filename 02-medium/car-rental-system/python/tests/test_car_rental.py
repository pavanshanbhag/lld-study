from datetime import date, timedelta

from car import Car
from customer import Customer
from rental_system import RentalSystem


def test_rental_system_constructor() -> None:
    rental_system = RentalSystem()
    rental_system.add_car(Car("Toyota", "Camry", 2022, "ABC123", 50.0))
    assert "ABC123" in rental_system.cars


def test_make_reservation_and_payment() -> None:
    rental_system = RentalSystem()
    car = Car("Toyota", "Camry", 2022, "ABC123", 50.0)
    rental_system.add_car(car)
    customer = Customer("John", "john@example.com", "DL1")

    start = date.today()
    end = start + timedelta(days=2)
    available = rental_system.search_cars("Toyota", "Camry", start, end)
    assert len(available) == 1

    reservation = rental_system.make_reservation(customer, car, start, end)
    assert reservation is not None
    assert rental_system.process_payment(reservation) is True
    assert car.available is False
