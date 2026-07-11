from typing import Dict, List

from bill import BaseBill, Bill
from command import PrepareOrderCommand, ServeOrderCommand
from decorators import ServiceChargeDecorator, TaxDecorator
from menu_item import MenuItem
from order import Order
from order_item import OrderItem
from restaurant import Restaurant
from staff import Chef, Waiter
from table import Table


class RestaurantManagementSystemFacade:
    def __init__(self, restaurant: Restaurant | None = None):
        self._restaurant = restaurant or Restaurant()
        self._order_id_counter = 1
        self._orders: Dict[int, Order] = {}

    def add_table(self, table_id: int, capacity: int) -> Table:
        table = Table(table_id, capacity)
        self._restaurant.add_table(table)
        return table

    def add_waiter(self, waiter_id: str, name: str) -> Waiter:
        waiter = Waiter(waiter_id, name)
        self._restaurant.add_waiter(waiter)
        return waiter

    def add_chef(self, chef_id: str, name: str) -> Chef:
        chef = Chef(chef_id, name)
        self._restaurant.add_chef(chef)
        return chef

    def add_menu_item(self, item_id: str, name: str, price: float) -> MenuItem:
        item = MenuItem(item_id, name, price)
        self._restaurant.menu.add_item(item)
        return item

    def take_order(self, table_id: int, waiter_id: str, menu_item_ids: List[str]) -> Order:
        waiter = self._restaurant.get_waiter(waiter_id)
        if waiter is None:
            raise ValueError("Invalid waiter ID.")

        chefs = self._restaurant.get_chefs()
        if not chefs:
            raise RuntimeError("No chefs available.")
        chef = chefs[0]

        order = Order(self._order_id_counter, table_id)
        self._order_id_counter += 1

        for item_id in menu_item_ids:
            menu_item = self._restaurant.menu.get_item(item_id)
            order_item = OrderItem(menu_item, order)
            order_item.add_observer(waiter)
            order.add_item(order_item)

        prepare_order_command = PrepareOrderCommand(order, chef)
        prepare_order_command.execute()

        self._orders[order.order_id] = order
        return order

    def mark_items_as_ready(self, order_id: int):
        order = self._orders[order_id]
        print(f"\nChef has finished preparing order {order.order_id}")

        for item in order.order_items:
            item.next_state()
            item.next_state()

    def serve_order(self, waiter_id: str, order_id: int):
        order = self._orders[order_id]
        waiter = self._restaurant.get_waiter(waiter_id)

        serve_order_command = ServeOrderCommand(order, waiter)
        serve_order_command.execute()

    def generate_bill(self, order_id: int) -> Bill:
        order = self._orders[order_id]
        bill_component = BaseBill(order)
        bill_component = TaxDecorator(bill_component, 0.08)
        bill_component = ServiceChargeDecorator(bill_component, 5.00)

        return Bill(bill_component)
