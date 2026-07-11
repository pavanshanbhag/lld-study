from typing import Dict, List, Optional
from order import Order
from stock import Stock
from enums import OrderType, OrderStatus
from order_states import FilledState
from collections import defaultdict

class StockExchange:
    def __init__(self):
        self.buy_orders: Dict[str, List[Order]] = defaultdict(list)
        self.sell_orders: Dict[str, List[Order]] = defaultdict(list)

    def place_buy_order(self, order: Order) -> None:
        self.buy_orders[order.get_stock().get_symbol()].append(order)
        self._match_orders(order.get_stock())

    def place_sell_order(self, order: Order) -> None:
        self.sell_orders[order.get_stock().get_symbol()].append(order)
        self._match_orders(order.get_stock())

    def _match_orders(self, stock: Stock) -> None:
        buys = self.buy_orders.get(stock.get_symbol(), [])
        sells = self.sell_orders.get(stock.get_symbol(), [])

        if not buys or not sells:
            return

        match_found = True
        while match_found:
            match_found = False
            best_buy = self._find_best_buy(buys)
            best_sell = self._find_best_sell(sells)

            if best_buy and best_sell:
                buy_price = stock.get_price() if best_buy.get_type() == OrderType.MARKET else best_buy.get_price()
                sell_price = stock.get_price() if best_sell.get_type() == OrderType.MARKET else best_sell.get_price()

                if buy_price >= sell_price:
                    self._execute_trade(best_buy, best_sell, sell_price)
                    match_found = True

    def _execute_trade(self, buy_order: Order, sell_order: Order, trade_price: float) -> None:
        print(f"--- Executing Trade for {buy_order.get_stock().get_symbol()} at ${trade_price:.2f} ---")

        buyer = buy_order.get_user()
        seller = sell_order.get_user()

        trade_quantity = min(buy_order.get_quantity(), sell_order.get_quantity())
        total_cost = trade_quantity * trade_price

        buyer.get_account().debit(total_cost)
        buyer.get_account().add_stock(buy_order.get_stock().get_symbol(), trade_quantity)

        seller.get_account().credit(total_cost)
        seller.get_account().remove_stock(sell_order.get_stock().get_symbol(), trade_quantity)

        self._update_order_status(buy_order, trade_quantity)
        self._update_order_status(sell_order, trade_quantity)

        buy_order.get_stock().set_price(trade_price)

        print("--- Trade Complete ---")

    def _update_order_status(self, order: Order, quantity_traded: int) -> None:
        order.set_status(OrderStatus.FILLED)
        order.set_state(FilledState())
        stock_symbol = order.get_stock().get_symbol()

        if order in self.buy_orders[stock_symbol]:
            self.buy_orders[stock_symbol].remove(order)
        if order in self.sell_orders[stock_symbol]:
            self.sell_orders[stock_symbol].remove(order)

    def _find_best_buy(self, buys: List[Order]) -> Optional[Order]:
        open_orders = [o for o in buys if o.get_status() == OrderStatus.OPEN]
        if not open_orders:
            return None
        return max(open_orders, key=lambda o: o.get_price())

    def _find_best_sell(self, sells: List[Order]) -> Optional[Order]:
        open_orders = [o for o in sells if o.get_status() == OrderStatus.OPEN]
        if not open_orders:
            return None
        return min(open_orders, key=lambda o: o.get_price())
