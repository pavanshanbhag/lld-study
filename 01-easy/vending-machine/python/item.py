from dataclasses import dataclass


@dataclass(slots=True)
class Item:
    code: str
    name: str
    price: int
