import datetime
from booking import Booking
from threading import Lock

class BookingManager:
    def __init__(self):
        self.bookings = {}
        self.booking_counter = 0
        self._lock = Lock()

    def create_booking(self, flight, passenger, seat, price):
        booking_number = self._generate_booking_number()
        booking = Booking(booking_number, flight, passenger, seat, price)
        with self._lock:
            self.bookings[booking_number] = booking
        return booking

    def cancel_booking(self, booking_number):
        with self._lock:
            booking = self.bookings.get(booking_number)
            if booking:
                booking.cancel()

    def _generate_booking_number(self):
        self.booking_counter += 1
        timestamp = datetime.datetime.now().strftime("%Y%m%d%H%M%S")
        return f"BKG{timestamp}{self.booking_counter:06d}"