
from enums import Ingredient


class Inventory:
    def __init__(self) -> None:
        self._stock: dict[Ingredient, int] = {}

    def add_stock(self, ingredient: Ingredient, quantity: int) -> None:
        self._stock[ingredient] = self._stock.get(ingredient, 0) + quantity

    def has_ingredients(self, recipe: dict[Ingredient, int]) -> bool:
        return all(self._stock.get(ingredient, 0) >= quantity for ingredient, quantity in recipe.items())

    def deduct_ingredients(self, recipe: dict[Ingredient, int]) -> None:
        if not self.has_ingredients(recipe):
            print("Not enough ingredients to make coffee.")
            return

        for ingredient, quantity in recipe.items():
            self._stock[ingredient] -= quantity

    def print_inventory(self) -> None:
        print("--- Current Inventory ---")
        for ingredient, quantity in self._stock.items():
            print(f"{ingredient.value}: {quantity}")
        print("-------------------------")
