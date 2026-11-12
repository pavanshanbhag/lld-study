from enum import Enum, StrEnum

class PaymentStatus(StrEnum):
    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"
    PENDING = "PENDING"

class SeatStatus(StrEnum):
    AVAILABLE = "AVAILABLE"
    LOCKED = "LOCKED"  # Temporarily held during booking process
    BOOKED = "BOOKED"

class SeatType(Enum):
    REGULAR = 50.0
    PREMIUM = 80.0
    RECLINER = 120.0

    def get_price(self) -> float:
        return self.value
