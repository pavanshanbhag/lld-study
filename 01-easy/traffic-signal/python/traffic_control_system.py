
from intersection_controller import IntersectionController
from observer import CentralMonitor


class TrafficControlSystem:
    def __init__(self) -> None:
        self._intersections: list[IntersectionController] = []
        self._executor_service = None

    def add_intersection(self, intersection_id: int, green_duration: int, yellow_duration: int) -> None:
        intersection = (
            IntersectionController.Builder(intersection_id)
            .with_durations(green_duration, yellow_duration)
            .add_observer(CentralMonitor())
            .build()
        )
        self._intersections.append(intersection)

    def start_system(self) -> None:
        if not self._intersections:
            print("No intersections to manage. System not starting.")
            return

        print("--- Starting Traffic Control System ---")
        from concurrent.futures import ThreadPoolExecutor

        self._executor_service = ThreadPoolExecutor(max_workers=len(self._intersections))
        for intersection in self._intersections:
            self._executor_service.submit(intersection.run)

    def stop_system(self) -> None:
        print("\n--- Shutting Down Traffic Control System ---")

        for intersection in self._intersections:
            intersection.stop()

        if self._executor_service:
            self._executor_service.shutdown(wait=True)

        print("All intersections stopped. System shut down.")

    @property
    def intersections(self) -> list[IntersectionController]:
        return self._intersections
