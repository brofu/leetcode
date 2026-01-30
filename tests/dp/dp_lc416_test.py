import pytest
from tests.helpers import Case
from dp.dp_lc416 import Solution

CASES_canPartition = [
    Case(name="case 1", args=([1,5,11,5],), want=True),
    Case(name="case 2", args=([1,2,3,5],), want=False),
    # Case(name="error", args=(), raises=ValueError),
]

@pytest.mark.parametrize("case", CASES_canPartition, ids=lambda c: c.name)
def test_canPartition(case: Case):
    s = Solution()

    if case.raises:
        with pytest.raises(case.raises):
            s.canPartition(*case.args, **case.kwargs)
        return

    got = s.canPartition(*case.args, **case.kwargs)
    assert got == case.want

