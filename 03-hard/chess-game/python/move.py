from __future__ import annotations

from dataclasses import dataclass


@dataclass(slots=True)
class Move:
    piece: object
    dest_row: int
    dest_col: int
