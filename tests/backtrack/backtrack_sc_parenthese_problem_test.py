import pytest
from tests.helpers import Case
from backtrack.backtrack_sc_parenthese_problem import parentheseProblem

CASES_parentheseProblem = [
    Case(name="basic", args=(3,), want=['((()))', '(()())', '(())()', '()(())', '()()()']),
    # Case(name="error", args=(), raises=ValueError),
]

@pytest.mark.parametrize("case", CASES_parentheseProblem, ids=lambda c: c.name)
def test_parentheseProblem(case: Case):
    if case.raises:
        with pytest.raises(case.raises):
            parentheseProblem(*case.args, **case.kwargs)
        return

    got = parentheseProblem(*case.args, **case.kwargs)
    assert got == case.want

