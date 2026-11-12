from dataclasses import dataclass


@dataclass(frozen=True, slots=True)
class Tag:
    name: str

    def get_name(self) -> str:
        return self.name
