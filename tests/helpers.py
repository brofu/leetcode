
from __future__ import annotations

from dataclasses import dataclass, field
from typing import Any, Callable, Generic, Iterable, TypeVar

T = TypeVar("T")

@dataclass(frozen=True)
class Case(Generic[T]):
    name: str
    args: tuple[Any, ...] = ()
    kwargs: dict[str, Any] = field(default_factory=dict)
    want: T | None = None
    raises: type[BaseException] | None = None

    def __post_init__(self):
        if self.kwargs is None:
            object.__setattr__(self, "kwargs", {})

