
class Solution:
        def checkInclusion(self, s1: str, s2: str) -> bool:
            print("flag", s1, s2)
            c: dict[str, int] = {}
            w: dict[str, int] = {}

            for ch in s1:
                c[ch] = c.get(ch, 0) + 1

            slow, fast = 0, 0
            valid = 0

            print("flag", s1, s2)
            while fast < len(s2):

                # extend the window
                ch = s2[fast]
                w[ch] = w.get(ch, 0)+1
                if not c.get(ch, 0) == 0 and c.get(ch, 0) == w.get(ch, 0):
                    valid += 1
                if not c.get(ch, 0) == 0 and c.get(ch, 0) + 1 == w.get(ch, 0):
                    valid -= 1
                if fast - slow + 1 < len(s1):
                    fast+=1
                    continue

                
                # check result
                if fast - slow + 1 == len(s1) and valid == len(c):
                    return True

                # shrink the window
                ch = s2[slow]
                w[ch] = w.get(ch, 0) - 1
                if not c.get(ch, 0) == 0 and w.get(ch, 0) < c.get(ch, 0):
                    valid -= 1
                slow += 1

                fast += 1

            return False



                    

