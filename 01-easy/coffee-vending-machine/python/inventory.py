from typing import Dict
from enums import Ingredient
import threading

class Inventory:
    _lock = threading.Lock()
    
    
    def __init__(self):
        self._stock: Dict[Ingredient, int] = {}
        self._lock = threading.Lock()
    
    
    def stock(self):
        return self._stock

    def add_stock(self, ingredient: Ingredient, quantity: int):
        self._stock[ingredient] = self._stock.get(ingredient, 0) + quantity
    
    def has_ingredients(self, recipe: Dict[Ingredient, int]) -> bool:
        return all(self._stock.get(ingredient, 0) >= quantity 
                  for ingredient, quantity in recipe.items())
    
    def deduct_ingredients(self, recipe: Dict[Ingredient, int]):
        with self._lock:
            if not self.has_ingredients(recipe):
                print("Not enough ingredients to make coffee.")
                return
            
            for ingredient, quantity in recipe.items():
                self._stock[ingredient] -= quantity
    
    def print_inventory(self):
        print("--- Current Inventory ---")
        for ingredient, quantity in self._stock.items():
            print(f"{ingredient.value}: {quantity}")
        print("-------------------------")