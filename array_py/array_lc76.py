class Solution:

        def hit(self, w: dict[str,int], c: dict[str, int]) -> bool:
            for k, v in c.items():
                if w.get(k, 0) < v:
                    return False
            return True

        def minWindow(self, s: str, t: str) -> str:
           
            resultCount = len(s)
            resultStr = ""

            c : dict[str, int] = {}
            w : dict[str, int] = {}
            for ch in t:
                c[ch] = c.get(ch, 0)+1

            slow, fast = 0, 0

            while fast < len(s):
                ch = s[fast]
                w[ch] = w.get(ch, 0)+1
                if not self.hit(w, c):
                    fast+= 1 
                    continue

                # record result
                if resultCount > fast - slow:
                    resultCount = fast - slow
                    resultStr = s[slow:fast+1]

                # shrink the window
                while slow <= fast - len(t)+1 and self.hit(w, c):
                    if resultCount > fast - slow:
                        resultCount = fast - slow
                        resultStr = s[slow:fast+1]
                    ch = s[slow]
                    w[ch] = w.get(ch, 0)-1
                    slow+=1
                   
                fast += 1

            return resultStr

        def minWindowV2(self, s: str, t: str) -> str:
            w: dict[str, int] = {}
            c: dict[str, int] = {}

            for ch in t:
                c[ch] = c.get(ch, 0)+1

            slow, fast, valid = 0, 0, 0
            result = ""
            best_len = float("inf")

            while fast < len(s):
                # extend the window
                n = s[fast]
                w[n] = w.get(n, 0)+1
                if w.get(n, 0) == c.get(n, 0) and not c.get(n, 0) == 0:
                    valid+=1 
                if valid < len(c):
                    fast += 1
                    continue

                # record the result
                if best_len > fast - slow + 1:
                    result = s[slow:fast+1]
                    best_len = fast - slow + 1

                # shrink the window
                while valid == len(c):
                    if best_len > fast - slow +1:
                        result = s[slow:fast+1]
                        best_len = fast - slow + 1
                    n = s[slow]
                    w[n] = w.get(n, 0)-1
                    if c.get(n, 0) != 0 and w.get(n, 0) < c.get(n, 0):
                        valid -= 1
                    slow+=1

                fast += 1

            return result
