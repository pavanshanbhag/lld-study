"""Parking lot demo entrypoint."""

from __future__ import annotations

import logging

from parking_lot import (
    Bike,
    Car,
    ParkingFloor,
    ParkingLot,
    ParkingSpot,
    Truck,
    VehicleBasedFeeStrategy,
    VehicleSize,
)

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")


def main() -> None:
    lot = ParkingLot(fee_strategy=VehicleBasedFeeStrategy())

    floor1 = ParkingFloor(floor_number=1)
    floor1.add_spot(ParkingSpot("F1-S1", VehicleSize.SMALL))
    floor1.add_spot(ParkingSpot("F1-M1", VehicleSize.MEDIUM))
    floor1.add_spot(ParkingSpot("F1-L1", VehicleSize.LARGE))

    floor2 = ParkingFloor(floor_number=2)
    floor2.add_spot(ParkingSpot("F2-M1", VehicleSize.MEDIUM))
    floor2.add_spot(ParkingSpot("F2-M2", VehicleSize.MEDIUM))

    lot.add_floor(floor1)
    lot.add_floor(floor2)

    print("--- Vehicle Entries ---")
    _print_availability(floor1, floor2)

    bike = Bike("B-123")
    car = Car("C-456")
    truck = Truck("T-789")

    lot.park_vehicle(bike)
    lot.park_vehicle(car)
    lot.park_vehicle(truck)

    print("--- Availability after parking ---")
    _print_availability(floor1, floor2)

    lot.park_vehicle(Car("C-999"))

    try:
        lot.park_vehicle(Bike("B-000"))
    except Exception as exc:  # noqa: BLE001 - demo
        print(f"Expected failure for B-000: {exc}")

    print("--- Vehicle Exits ---")
    fee = lot.unpark_vehicle(car.license_number)
    print(f"Car C-456 unparked. Fee: ${fee:.2f}")

    print("--- Availability after one car leaves ---")
    _print_availability(floor1, floor2)


def _print_availability(floor1: ParkingFloor, floor2: ParkingFloor) -> None:
    for floor in (floor1, floor2):
        print(f"--- Floor {floor.floor_number} Availability ---")
        for size, count in floor.availability().items():
            print(f"  {size.value} spots: {count}")


if __name__ == "__main__":
    main()
