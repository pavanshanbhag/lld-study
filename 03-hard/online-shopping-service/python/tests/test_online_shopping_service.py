from enums import ProductCategory
from product import Product
from online_shopping_system import OnlineShoppingSystem


def test_online_shopping_system() -> None:
    system = OnlineShoppingSystem()
    product = Product.Builder("Laptop", 1500.0).with_category(ProductCategory.ELECTRONICS).build()
    system.add_product(product, 5)

    results = system.search_products("Laptop")
    assert len(results) == 1
