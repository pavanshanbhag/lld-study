from movie import Movie
from movie_booking_service import MovieBookingService


def test_movie_booking_service() -> None:
    service = MovieBookingService()
    movie = Movie("M1", "The Matrix", 120)
    service.add_movie(movie)

    city = service.add_city("city1", "New York")
    assert city.get_name() == "New York"
    assert service.movies["M1"].get_title() == "The Matrix"


def test_create_user() -> None:
    service = MovieBookingService()
    user = service.create_user("Alice", "alice@example.com")
    assert user.get_name() == "Alice"
