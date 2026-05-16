#.  z.  x.  y.   z.  x.   y.   z
#.                   i         j      




class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        hashset = set()
        maxLen = 0
        i = 0
        j = 1
        if len(s) <= 1:
            return len(s)
        
        hashset.add(ord(s[0]))
        while j != len(s):
            if ord(s[j]) in hashset:
                while s[i] != s[j]:
                    hashset.remove(ord(s[i]))
                    i = i + 1
                hashset.remove(ord(s[i]))
                i = i + 1
            else:
                if j - i + 1 > maxLen:
                    maxLen = j - i + 1
                hashset.add(ord(s[j]))
                j = j + 1
            print(i,j,maxLen)


        return maxLen
                