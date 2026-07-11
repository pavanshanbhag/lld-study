from datetime import datetime

from concert import Concert
from concert_ticket_booking_system import ConcertTicketBookingSystem
from seat import Seat, SeatType
from user import User


def test_book_tickets() -> None:
    system = ConcertTicketBookingSystem()
    seats = [Seat("S1", "S1", SeatType.REGULAR, 50.0)]
    concert = Concert("C001", "Artist 1", "Venue 1", datetime.now(), seats)
    system.add_concert(concert)

    user = User("U001", "John Doe", "john@example.com")
    booking = system.book_tickets(user, concert, seats)

    assert booking.id in system.bookings
    assert seats[0].status.name == "BOOKED"
