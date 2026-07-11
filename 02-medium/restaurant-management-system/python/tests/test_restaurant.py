from restaurant_management_system_facade import RestaurantManagementSystemFacade


def test_restaurant_facade_constructor() -> None:
    facade = RestaurantManagementSystemFacade()
    table = facade.add_table(1, 4)
    chef = facade.add_chef("CHEF01", "Gordon")
    waiter = facade.add_waiter("W01", "Alice")
    pizza = facade.add_menu_item("PIZZA01", "Margherita Pizza", 12.50)

    order = facade.take_order(table.id, waiter.id, [pizza.get_id()])
    assert order.order_id == 1
