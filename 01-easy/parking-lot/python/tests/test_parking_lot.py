from __future__ import annotations

from datetime import UTC, datetime, timedelta

import pytest
from parking_lot import (
    Bike,
    Car,
    LotFullError,
    ParkingFloor,
    ParkingLot,
    ParkingSpot,
    TicketNotFoundError,
    Truck,
    VehicleSize,
)


def test_park_and_unpark_car() -> None:
    lot = ParkingLot()
    floor = ParkingFloor(floor_number=1)
    floor.add_spot(ParkingSpot("F1-M1", VehicleSize.MEDIUM))
    lot.add_floor(floor)

    ticket = lot.park_vehicle(Car("C-123"))
    assert ticket.vehicle.license_number == "C-123"

    fee = lot.unpark_vehicle("C-123")
    assert fee >= 0


def test_lot_full_raises() -> None:
    lot = ParkingLot()
    floor = ParkingFloor(floor_number=1)
    floor.add_spot(ParkingSpot("F1-S1", VehicleSize.SMALL))
    lot.add_floor(floor)

    lot.park_vehicle(Bike("B-1"))
    with pytest.raises(LotFullError):
        lot.park_vehicle(Bike("B-2"))


def test_unpark_missing_ticket() -> None:
    lot = ParkingLot()
    with pytest.raises(TicketNotFoundError):
        lot.unpark_vehicle("missing")


def test_ticket_duration_uses_datetime() -> None:
    lot = ParkingLot()
    floor = ParkingFloor(floor_number=1)
    floor.add_spot(ParkingSpot("F1-L1", VehicleSize.LARGE))
    lot.add_floor(floor)

    ticket = lot.park_vehicle(Truck("T-1"))
    ticket.entry_time = datetime.now(UTC) - timedelta(hours=2)
    ticket.close()
    fee = lot.fee_strategy.calculate_fee(ticket)
    assert fee >= 20.0
