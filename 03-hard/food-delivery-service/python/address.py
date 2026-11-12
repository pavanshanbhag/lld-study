from __future__ import annotations

import math
from dataclasses import dataclass


@dataclass(slots=True)
class Address:
    street: str
    city: str
    zip_code: str
    latitude: float
    longitude: float

    def get_city(self) -> str:
        return self.city

    def distance_to(self, other: Address) -> float:
        lat_diff = self.latitude - other.latitude
        lon_diff = self.longitude - other.longitude
        return math.sqrt(lat_diff * lat_diff + lon_diff * lon_diff)

    def __str__(self) -> str:
        return f"{self.street}, {self.city}, {self.zip_code} @({self.latitude}, {self.longitude})"
