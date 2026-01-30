

def parentheseProblem(n: int) -> list[str]:
   
    result: list[str] = []

    def bt(left:int, right:int, path:str):

        # base case
        if left == 0 and right == 0 :
            result.append(path)
            return

        # choose, explore, cancel-choose
        if left > 0:
            bt(left-1, right, path+"(")

        if right > 0 and right > left:
            bt(left, right-1, path+")")

    bt(n, n, "")
    return result
