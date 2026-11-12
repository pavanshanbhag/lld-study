from stock_brokerage_system import StockBrokerageSystem


def test_stock_brokerage_system() -> None:
    system = StockBrokerageSystem()
    user = system.register_user("Alice", 10000.0)
    stock = system.add_stock("AAPL", 150.0)

    assert user.get_account().get_balance() == 10000.0
    assert stock.get_symbol() == "AAPL"


def test_register_multiple_users() -> None:
    system = StockBrokerageSystem()
    alice = system.register_user("Alice", 1000.0)
    bob = system.register_user("Bob", 2000.0)
    assert alice.get_account().get_balance() == 1000.0
    assert bob.get_account().get_balance() == 2000.0
