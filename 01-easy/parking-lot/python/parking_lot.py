"""Parking lot domain model."""

from __future__ import annotations

import logging
import threading
import uuid
from abc import ABC, abstractmethod
from dataclasses import dataclass, field
from datetime import UTC, datetime, timedelta
from enum import StrEnum
from typing import Protocol

logger = logging.getLogger(__name__)


class VehicleSize(StrEnum):
    SMALL = "SMALL"
    MEDIUM = "MEDIUM"
    LARGE = "LARGE"


_SIZE_RANK: dict[VehicleSize, int] = {
    VehicleSize.SMALL: 0,
    VehicleSize.MEDIUM: 1,
    VehicleSize.LARGE: 2,
}


class Vehicle(ABC):
    @property
    @abstractmethod
    def license_number(self) -> str: ...

    @property
    @abstractmethod
    def size(self) -> VehicleSize: ...


@dataclass(frozen=True, slots=True)
class Bike(Vehicle):
    license_number: str

    @property
    def size(self) -> VehicleSize:
        return VehicleSize.SMALL


@dataclass(frozen=True, slots=True)
class Car(Vehicle):
    license_number: str

    @property
    def size(self) -> VehicleSize:
        return VehicleSize.MEDIUM


@dataclass(frozen=True, slots=True)
class Truck(Vehicle):
    license_number: str

    @property
    def size(self) -> VehicleSize:
        return VehicleSize.LARGE


class ParkingError(Exception):
    """Base error for parking operations."""


class LotFullError(ParkingError):
    """Raised when no spot is available."""


class TicketNotFoundError(ParkingError):
    """Raised when no active ticket exists for a vehicle."""


@dataclass(slots=True)
class ParkingTicket:
    vehicle: Vehicle
    spot_id: str
    ticket_id: str = field(default_factory=lambda: str(uuid.uuid4()))
    entry_time: datetime = field(default_factory=lambda: datetime.now(UTC))
    exit_time: datetime | None = None

    def close(self) -> None:
        self.exit_time = datetime.now(UTC)

    @property
    def duration(self) -> timedelta:
        end = self.exit_time or datetime.now(UTC)
        return end - self.entry_time


class FeeStrategy(Protocol):
    def calculate_fee(self, ticket: ParkingTicket) -> float: ...


class FlatRateFeeStrategy:
    rate_per_hour: float = 10.0

    def calculate_fee(self, ticket: ParkingTicket) -> float:
        hours = max(1, int(ticket.duration.total_seconds() // 3600) + 1)
        return hours * self.rate_per_hour


class VehicleBasedFeeStrategy:
    hourly_rates: dict[VehicleSize, float] = {
        VehicleSize.SMALL: 10.0,
        VehicleSize.MEDIUM: 20.0,
        VehicleSize.LARGE: 30.0,
    }

    def calculate_fee(self, ticket: ParkingTicket) -> float:
        hours = max(1, int(ticket.duration.total_seconds() // 3600) + 1)
        return hours * self.hourly_rates[ticket.vehicle.size]


class ParkingStrategy(Protocol):
    def find_spot(self, floors: list[ParkingFloor], vehicle: Vehicle) -> ParkingSpot | None: ...


class NearestFirstStrategy:
    def find_spot(self, floors: list[ParkingFloor], vehicle: Vehicle) -> ParkingSpot | None:
        for floor in floors:
            spot = floor.find_available_spot(vehicle)
            if spot is not None:
                return spot
        return None


@dataclass
class ParkingSpot:
    spot_id: str
    spot_size: VehicleSize
    _occupied: bool = field(default=False, repr=False)
    _vehicle: Vehicle | None = field(default=None, repr=False)
    _lock: threading.Lock = field(default_factory=threading.Lock, repr=False)

    @property
    def is_available(self) -> bool:
        with self._lock:
            return not self._occupied

    def can_fit(self, vehicle: Vehicle) -> bool:
        if self._occupied:
            return False
        if vehicle.size == VehicleSize.SMALL:
            return self.spot_size == VehicleSize.SMALL
        if vehicle.size == VehicleSize.MEDIUM:
            return self.spot_size in {VehicleSize.MEDIUM, VehicleSize.LARGE}
        return self.spot_size == VehicleSize.LARGE

    def park(self, vehicle: Vehicle) -> None:
        with self._lock:
            if not self.can_fit(vehicle):
                raise ParkingError(
                    f"spot {self.spot_id} cannot fit vehicle {vehicle.license_number}"
                )
            self._vehicle = vehicle
            self._occupied = True

    def unpark(self) -> None:
        with self._lock:
            self._vehicle = None
            self._occupied = False


@dataclass
class ParkingFloor:
    floor_number: int
    spots: dict[str, ParkingSpot] = field(default_factory=dict)
    _lock: threading.Lock = field(default_factory=threading.Lock, repr=False)

    def add_spot(self, spot: ParkingSpot) -> None:
        self.spots[spot.spot_id] = spot

    def find_available_spot(self, vehicle: Vehicle) -> ParkingSpot | None:
        with self._lock:
            candidates = [s for s in self.spots.values() if s.is_available and s.can_fit(vehicle)]
            if not candidates:
                return None
            return min(candidates, key=lambda s: _SIZE_RANK[s.spot_size])

    def availability(self) -> dict[VehicleSize, int]:
        counts = dict.fromkeys(VehicleSize, 0)
        for spot in self.spots.values():
            if spot.is_available:
                counts[spot.spot_size] += 1
        return counts


@dataclass
class ParkingLot:
    floors: list[ParkingFloor] = field(default_factory=list)
    active_tickets: dict[str, ParkingTicket] = field(default_factory=dict)
    fee_strategy: FeeStrategy = field(default_factory=FlatRateFeeStrategy)
    parking_strategy: ParkingStrategy = field(default_factory=NearestFirstStrategy)
    _lock: threading.Lock = field(default_factory=threading.Lock, repr=False)

    def add_floor(self, floor: ParkingFloor) -> None:
        self.floors.append(floor)

    def park_vehicle(self, vehicle: Vehicle) -> ParkingTicket:
        with self._lock:
            spot = self.parking_strategy.find_spot(self.floors, vehicle)
            if spot is None:
                raise LotFullError(f"no available spot for {vehicle.license_number}")

            spot.park(vehicle)
            ticket = ParkingTicket(vehicle=vehicle, spot_id=spot.spot_id)
            self.active_tickets[vehicle.license_number] = ticket
            logger.info("parked %s at %s", vehicle.license_number, spot.spot_id)
            return ticket

    def unpark_vehicle(self, license_number: str) -> float:
        with self._lock:
            ticket = self.active_tickets.pop(license_number, None)
            if ticket is None:
                raise TicketNotFoundError(license_number)

            for floor in self.floors:
                spot = floor.spots.get(ticket.spot_id)
                if spot is not None:
                    spot.unpark()
                    break

            ticket.close()
            fee = self.fee_strategy.calculate_fee(ticket)
            logger.info("unparked %s from %s fee=%.2f", license_number, ticket.spot_id, fee)
            return fee
