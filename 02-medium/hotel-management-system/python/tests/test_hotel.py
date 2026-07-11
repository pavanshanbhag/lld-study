from datetime import date, timedelta

from guest import Guest
from hotel_management_system import HotelManagementSystem
from room import Room, RoomType


def test_book_room() -> None:
    system = HotelManagementSystem()
    guest = Guest("G001", "John Doe", "john@example.com", "1234567890")
    room = Room("R001", RoomType.SINGLE, 100.0)
    system.add_guest(guest)
    system.add_room(room)

    check_in = date.today()
    check_out = check_in + timedelta(days=2)
    reservation = system.book_room(guest, room, check_in, check_out)

    assert reservation is not None
    assert reservation.id in system.reservations
