class Solution:

    def isSubset(self, hash_s, hash_t):
        for key, value in hash_t.items():
            if key not in hash_s.keys():
                return False
            else:
                if hash_s[key] < value:
                    return False
        return True

    def minWindow(self, s: str, t: str) -> str:
    
        if len(t) > len(s):
            return ""

        hash_t = {}

        for k in range(len(t)):
            if t[k] in hash_t.keys():
                hash_t[t[k]] = hash_t[t[k]] + 1
            else:
                hash_t[t[k]] = 1

        minLen = len(s) + 1
        ans = ""

        hash_s = {}
        i = 0
        j = len(t) - 1

        k = 0
        while k != j + 1:
            if s[k] in hash_s.keys():
                hash_s[s[k]] = hash_s[s[k]] + 1
            else:
                hash_s[s[k]] = 1
            k = k + 1

        print(hash_s)
        print(hash_t)
        if hash_s.items() < hash_t.items():
            print("True")        

        while j <= len(s):
            
            #print(hash_s)

            if self.isSubset(hash_s, hash_t):
                print(i)
                print(j)
                if j - i + 1 < minLen:
                    minLen = j - i + 1
                    ans = s[i:j+1]
                    print(minLen)
                i = i + 1
                j = i + len(t) - 1
                if j >= len(s):
                    break
                hash_s = {}
                k = i
                while k <= j:
                    if s[k] in hash_s.keys():
                        hash_s[s[k]] = hash_s[s[k]] + 1
                    else:
                        hash_s[s[k]] = 1
                    k = k + 1
            else:
                j = j + 1
                if j == len(s):
                    return ans
                if s[j] in hash_s.keys():
                    hash_s[s[j]] = hash_s[s[j]] + 1
                else:
                    hash_s[s[j]] = 1
                
                if s[i] not in t:
                    hash_s[s[i]] = hash_s[s[i]] - 1
                    if hash_s[s[i]] == 0:
                        hash_s.pop(s[i])
                    i = i + 1
        return ans

