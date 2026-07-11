from cricinfo_service import CricInfoService
from enums import PlayerRole


def test_cricinfo_service() -> None:
    service = CricInfoService()
    player = service.add_player("P1", "Virat", PlayerRole.BATSMAN)

    assert player.get_name() == "Virat"
    assert service.player_repository.find_by_id("P1") is not None
