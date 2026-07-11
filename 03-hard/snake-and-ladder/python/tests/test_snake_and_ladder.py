from board import Board
from ladder import Ladder
from snake import Snake


def test_snake_and_ladder_board() -> None:
    board = Board(100, [Ladder(1, 38), Snake(16, 6)])
    assert board.get_final_position(1) == 38
    assert board.get_final_position(16) == 6
    assert board.get_final_position(5) == 5
