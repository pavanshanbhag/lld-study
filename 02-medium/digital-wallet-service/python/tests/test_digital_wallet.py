from account import Account
from currency import Currency
from digital_wallet import DigitalWallet
from user import User


def test_digital_wallet_constructor() -> None:
    wallet = DigitalWallet()
    user = User("U001", "John Doe", "john@example.com", "password123")
    wallet.create_user(user)
    assert wallet.get_user("U001") is user

    account = Account("A001", user, "1234567890", Currency.USD)
    wallet.create_account(account)
    assert wallet.get_account("A001") is account
