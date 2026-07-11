from elevator_system import ElevatorSystem


def test_elevator_system_constructor() -> None:
    system = ElevatorSystem(2)
    assert len(system.elevators) == 2
    assert 1 in system.elevators
    assert 2 in system.elevators
    system.shutdown()
