from address import Address
from food_delivery_service import FoodDeliveryService
from menu_item import MenuItem


def test_food_delivery_service() -> None:
    service = FoodDeliveryService()
    address = Address("123 Main St", "Springfield", "12345", 40.7, -74.0)
    customer = service.register_customer("Alice", "555-0100", address)
    restaurant = service.register_restaurant("Pizza Place", address)

    assert customer.get_name() == "Alice"
    assert restaurant.get_name() == "Pizza Place"


def test_place_order() -> None:
    service = FoodDeliveryService()
    address = Address("123 Main St", "Springfield", "12345", 40.7, -74.0)
    customer = service.register_customer("Alice", "555-0100", address)
    restaurant = service.register_restaurant("Pizza Place", address)
    item = MenuItem("m1", "Margherita", 12.0)
    restaurant.add_to_menu(item)
    from order_item import OrderItem

    order = service.place_order(customer.get_id(), restaurant.get_id(), [OrderItem(item, 1)])
    assert order is not None
