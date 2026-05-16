class Solution:

    def isInteger(self, num: int) -> bool:
        try:
            out = int(num)
            return True
        except Exception as ex:
            return False

    def isPalindrome(self, s: str) -> bool:

        index = 0
        last = len(s) - 1

        print(ord('0'))
        print(ord('9'))

        while index != last and index < last:
            first_lower = s[index].lower()
            last_lower = s[last].lower()

            if not self.isInteger(first_lower) and (ord(first_lower) < 97 or ord(first_lower) > 122):
                index = index + 1
                continue

            if not self.isInteger(last_lower) and (ord(last_lower) < 97 or ord(last_lower) > 122):
                last = last - 1
                continue

            if s[index].lower() != s[last].lower():
                return False
            index = index + 1
            last = last - 1

        return True

