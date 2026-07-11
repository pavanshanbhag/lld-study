import sys
from concurrent.futures import ThreadPoolExecutor

from direction import Direction
from elevator import Elevator
from elevator_observer import Display
from elevator_selection_strategy import NearestElevatorStrategy
from request import Request
from request_source import RequestSource


class ElevatorSystem:
    def __init__(self, num_elevators: int):
        self.selection_strategy = NearestElevatorStrategy()
        self.executor_service = ThreadPoolExecutor(max_workers=num_elevators)

        elevator_list = []
        display = Display()

        for i in range(1, num_elevators + 1):
            elevator = Elevator(i)
            elevator.add_observer(display)
            elevator_list.append(elevator)

        self.elevators = {elevator.get_id(): elevator for elevator in elevator_list}

    def start(self):
        for elevator in self.elevators.values():
            self.executor_service.submit(elevator.run)

    def request_elevator(self, floor: int, direction: Direction):
        print(f"\n>> EXTERNAL Request: User at floor {floor} wants to go {direction.value}")
        request = Request(floor, direction, RequestSource.EXTERNAL)

        selected_elevator = self.selection_strategy.select_elevator(list(self.elevators.values()), request)

        if selected_elevator:
            selected_elevator.add_request(request)
        else:
            print("System busy, please wait.")

    def select_floor(self, elevator_id: int, destination_floor: int):
        print(f"\n>> INTERNAL Request: User in Elevator {elevator_id} selected floor {destination_floor}")
        request = Request(destination_floor, Direction.IDLE, RequestSource.INTERNAL)

        elevator = self.elevators.get(elevator_id)
        if elevator:
            elevator.add_request(request)
        else:
            print("Invalid elevator ID.", file=sys.stderr)

    def shutdown(self):
        print("Shutting down elevator system...")
        for elevator in self.elevators.values():
            elevator.stop_elevator()
        self.executor_service.shutdown()
