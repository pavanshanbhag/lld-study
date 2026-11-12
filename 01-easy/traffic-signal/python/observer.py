from __future__ import annotations

import logging
from abc import ABC, abstractmethod

from direction import Direction
from light_color import LightColor

logger = logging.getLogger(__name__)


class TrafficObserver(ABC):
    @abstractmethod
    def update(self, intersection_id: int, direction: Direction, color: LightColor):
        pass


class CentralMonitor(TrafficObserver):
    def update(self, intersection_id: int, direction: Direction, color: LightColor):
        logger.info(
            "Intersection %s: Light for %s direction changed to %s.",
            intersection_id,
            direction.value,
            color.value,
        )
