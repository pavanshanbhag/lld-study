from address import Address
from food_delivery_service import FoodDeliveryService


def test_food_delivery_service() -> None:
    service = FoodDeliveryService()
    address = Address("123 Main St", "Springfield", "12345", 40.7, -74.0)
    customer = service.register_customer("Alice", "555-0100", address)
    restaurant = service.register_restaurant("Pizza Place", address)

    assert customer.get_name() == "Alice"
    assert restaurant.get_name() == "Pizza Place"
