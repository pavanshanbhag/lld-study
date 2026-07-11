from game import Game
from player import Player
from symbol import Symbol
from tictactoesystem import TicTacToeSystem


def test_game_alice_wins_top_row() -> None:
    alice = Player("Alice", Symbol.X)
    bob = Player("Bob", Symbol.O)
    game = Game(alice, bob)

    game.make_move(alice, 0, 0)
    game.make_move(bob, 1, 0)
    game.make_move(alice, 0, 1)
    game.make_move(bob, 1, 1)
    game.make_move(alice, 0, 2)

    assert game.get_winner() == alice


def test_tic_tac_toe_system_constructor() -> None:
    system = TicTacToeSystem()
    alice = Player("Alice", Symbol.X)
    bob = Player("Bob", Symbol.O)
    system.create_game(alice, bob)
    assert system.game is not None
