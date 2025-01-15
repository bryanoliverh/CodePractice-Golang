class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        if len(s) % 2 == 1:
            return False

        open_bracket = []
        wildcard = []
        n = len(s)

        for i in range(n):
            if locked[i] == '0':
                wildcard.append(i)
            elif s[i] == '(':
                open_bracket.append(i)
            elif not open_bracket and not wildcard:
                return False
            else:
                if open_bracket:
                    open_bracket.pop()
                else:
                    wildcard.pop()
        while open_bracket and wildcard and open_bracket[-1] < wildcard[-1]:
            open_bracket.pop()
            wildcard.pop()

        if open_bracket:
            return False
        return True