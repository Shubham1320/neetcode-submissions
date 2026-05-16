class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        
        if len(s2) < len(s1):
            return False

        hash_s1 = {}
        for i in range(len(s1)):
            if s1[i] in hash_s1.keys():
                hash_s1[s1[i]] = hash_s1[s1[i]] + 1
            else:
                hash_s1[s1[i]] = 1
        
        hash_s2 = {}

        i = 0
        j = len(s1) - 1

        k = 0
        while k != j + 1:
            if s2[k] in hash_s2.keys():
                hash_s2[s2[k]] = hash_s2[s2[k]] + 1
            else:
                hash_s2[s2[k]] = 1
            k = k + 1

        print(hash_s1)
        print(hash_s2)

        while j != len(s2):
            print(i)
            print(j)
            if hash_s1 == hash_s2:
                return True
            else:
                hash_s2[s2[i]] = hash_s2[s2[i]] - 1
                if hash_s2[s2[i]] == 0:
                    hash_s2.pop(s2[i])

                i = i + 1
                j = j + 1
                if j == len(s2):
                    break
                if s2[j] in hash_s2.keys():
                    hash_s2[s2[j]] = hash_s2[s2[j]] + 1
                else:
                    hash_s2[s2[j]] = 1

        return False
