from game import Game
from board import Board
from color import Color
from move import Move
from player import Player


def test_chess_game_constructor() -> None:
    game = Game()
    assert game.board is not None
    assert len(game.players) == 2


def test_chess_valid_pawn_move() -> None:
    board = Board()
    player = Player(Color.WHITE)
    piece = board.get_piece(1, 0)
    move = Move(piece, 2, 0)
    player.make_move(board, move)
    assert board.get_piece(2, 0) is not None
