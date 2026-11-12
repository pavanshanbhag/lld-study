from __future__ import annotations

from dataclasses import dataclass

from symbol import Symbol


@dataclass(slots=True)
class Player:
    name: str
    symbol: Symbol

    def get_name(self) -> str:
        return self.name

    def get_symbol(self) -> Symbol:
        return self.symbol
