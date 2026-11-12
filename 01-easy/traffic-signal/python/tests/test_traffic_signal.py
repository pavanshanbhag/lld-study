from traffic_control_system import TrafficControlSystem
from light_color import LightColor
from direction import Direction


def test_add_intersection() -> None:
    system = TrafficControlSystem()
    system.add_intersection(1, 500, 200)
    assert len(system.intersections) == 1


def test_intersection_has_four_lights() -> None:
    system = TrafficControlSystem()
    system.add_intersection(1, 100, 50)
    intersection = system.intersections[0]
    for direction in Direction:
        light = intersection.get_light(direction)
        assert light.get_current_color() in {
            LightColor.RED,
            LightColor.YELLOW,
            LightColor.GREEN,
            None,
        }
