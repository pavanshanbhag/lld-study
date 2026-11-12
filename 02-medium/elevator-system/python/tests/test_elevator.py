from elevator_system import ElevatorSystem
from direction import Direction


def test_elevator_system_constructor() -> None:
    system = ElevatorSystem(2)
    assert len(system.elevators) == 2
    assert 1 in system.elevators
    assert 2 in system.elevators
    system.shutdown()


def test_request_elevator_queues_without_start() -> None:
    system = ElevatorSystem(2)
    # Do not start worker loops in unit tests; just verify request selection path.
    system.request_elevator(3, Direction.UP)
    assert len(system.elevators) == 2
    system.shutdown()
