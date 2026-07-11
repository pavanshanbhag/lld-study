from typing import Dict, List, Optional
from city import City
from cinema import Cinema
from movie import Movie
from screen import Screen
from show import Show
from user import User
from seat_lock_manager import SeatLockManager
from booking_manager import BookingManager
from payment_strategy import PaymentStrategy
from datetime import datetime
from booking import Booking
from pricing_strategy import PricingStrategy
from seat import Seat

class MovieBookingService:
    def __init__(self):
        self.cities: Dict[str, City] = {}
        self.cinemas: Dict[str, Cinema] = {}
        self.movies: Dict[str, Movie] = {}
        self.users: Dict[str, User] = {}
        self.shows: Dict[str, Show] = {}

        self.seat_lock_manager = SeatLockManager()
        self.booking_manager = BookingManager(self.seat_lock_manager)

    def get_booking_manager(self) -> BookingManager:
        return self.booking_manager

    def add_city(self, city_id: str, name: str) -> City:
        city = City(city_id, name)
        self.cities[city.get_id()] = city
        return city

    def add_cinema(self, cinema_id: str, name: str, city_id: str, screens: List[Screen]) -> Cinema:
        city = self.cities[city_id]
        cinema = Cinema(cinema_id, name, city, screens)
        self.cinemas[cinema.get_id()] = cinema
        return cinema

    def add_movie(self, movie: Movie) -> None:
        self.movies[movie.get_id()] = movie

    def add_show(self, show_id: str, movie: Movie, screen: Screen, start_time: datetime, pricing_strategy: PricingStrategy) -> Show:
        show = Show(show_id, movie, screen, start_time, pricing_strategy)
        self.shows[show.get_id()] = show
        return show

    def create_user(self, name: str, email: str) -> User:
        user = User(name, email)
        self.users[user.get_id()] = user
        return user

    def book_tickets(self, user_id: str, show_id: str, desired_seats: List[Seat], payment_strategy: PaymentStrategy) -> Optional[Booking]:
        return self.booking_manager.create_booking(
            self.users[user_id],
            self.shows[show_id],
            desired_seats,
            payment_strategy
        )

    def find_shows(self, movie_title: str, city_name: str) -> List[Show]:
        result = []
        for show in self.shows.values():
            if show.get_movie().get_title().lower() == movie_title.lower():
                cinema = self._find_cinema_for_show(show)
                if cinema and cinema.get_city().get_name().lower() == city_name.lower():
                    result.append(show)
        return result

    def _find_cinema_for_show(self, show: Show) -> Optional[Cinema]:
        for cinema in self.cinemas.values():
            if show.get_screen() in cinema.get_screens():
                return cinema
        return None

    def shutdown(self) -> None:
        self.seat_lock_manager.shutdown()
        print("MovieTicketBookingSystem has been shut down.")
