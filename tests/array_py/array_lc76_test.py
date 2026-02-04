import pytest
from tests.helpers import Case
from array_py.array_lc76 import Solution

CASES_minWindow = [
    Case(name="basic", args=("ADOBECODEBANC", "ABC"), want="BANC"),
    Case(name="base 1", args=("a", "a"), want="a"),
    Case(name="base 1", args=("a", "aa"), want=""),
    Case(name="base 3", args=("ab", "b"), want="b"),
]

@pytest.mark.parametrize("case", CASES_minWindow, ids=lambda c: c.name)
def test_minWindow(case: Case):
    s = Solution()

    if case.raises:
        with pytest.raises(case.raises):
            s.minWindow(*case.args, **case.kwargs)
        return

    got = s.minWindow(*case.args, **case.kwargs)
    assert got == case.want

CASES_minWindowV2 = [
    Case(name="basic", args=("ADOBECODEBANC", "ABC"), want="BANC"),
    Case(name="base 1", args=("a", "a"), want="a"),
    Case(name="base 1", args=("a", "aa"), want=""),
    Case(name="base 3", args=("ab", "b"), want="b"),
]

@pytest.mark.parametrize("case", CASES_minWindowV2, ids=lambda c: c.name)
def test_minWindowV2(case: Case):
    s = Solution()

    if case.raises:
        with pytest.raises(case.raises):
            s.minWindowV2(*case.args, **case.kwargs)
        return

    got = s.minWindowV2(*case.args, **case.kwargs)
    assert got == case.want

