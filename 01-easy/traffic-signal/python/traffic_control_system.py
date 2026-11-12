from intersection_controller import IntersectionController
from observer import CentralMonitor
from typing import List
import threading
from concurrent.futures import ThreadPoolExecutor

class TrafficControlSystem:
    _lock = threading.Lock()

    def __init__(self):
        self._intersections: List[IntersectionController] = []
        self._executor_service = None

    @property
    def intersections(self) -> List[IntersectionController]:
        return self._intersections

    def add_intersection(self, intersection_id: int, green_duration: int, yellow_duration: int):
        intersection = IntersectionController.Builder(intersection_id) \
            .with_durations(green_duration, yellow_duration) \
            .add_observer(CentralMonitor()) \
            .build()
        self._intersections.append(intersection)

    def start_system(self):
        if not self._intersections:
            print("No intersections to manage. System not starting.")
            return

        print("--- Starting Traffic Control System ---")
        self._executor_service = ThreadPoolExecutor(max_workers=len(self._intersections))
        
        for intersection in self._intersections:
            self._executor_service.submit(intersection.run)

    def stop_system(self):
        print("\n--- Shutting Down Traffic Control System ---")
        
        for intersection in self._intersections:
            intersection.stop()
        
        if self._executor_service:
            self._executor_service.shutdown(wait=True)
        
        print("All intersections stopped. System shut down.")