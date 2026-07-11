from traffic_control_system import TrafficControlSystem


def test_add_intersection() -> None:
    system = TrafficControlSystem()
    system.add_intersection(1, 500, 200)
    assert len(system.intersections) == 1
