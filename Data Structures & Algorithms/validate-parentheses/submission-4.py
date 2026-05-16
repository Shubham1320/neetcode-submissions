class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        hash_map = {
            ')' : '(',
            '}' : '{',
            ']' : '['
        }

        if len(s) <= 1:
            return False

        for i in range(len(s)):
            if s[i] in hash_map.values():
                stack.append(s[i])
            elif s[i] in hash_map.keys():
                if stack and stack[-1] == hash_map[s[i]]:
                    stack.pop()
                else:
                    return False
        
        if len(stack) != 0:
            return False
        return True
