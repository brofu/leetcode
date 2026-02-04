import pytest
from tests.helpers import Case
from array_py.array_lc567 import Solution

CASES_checkInclusion = [
    Case(name="basic", args=("ab", "eidbaooo"), want=True),
    Case(name="basic2", args=("ab", "eidboaoo"), want=False),
    Case(name="basic3", args=("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine"), want=True),
    # Case(name="error", args=(), raises=ValueError),
]

@pytest.mark.parametrize("case", CASES_checkInclusion, ids=lambda c: c.name)
def test_checkInclusion(case: Case):
    s = Solution()

    if case.raises:
        with pytest.raises(case.raises):
            s.checkInclusion(*case.args, **case.kwargs)
        return

    got = s.checkInclusion(*case.args, **case.kwargs)
    assert got == case.want

