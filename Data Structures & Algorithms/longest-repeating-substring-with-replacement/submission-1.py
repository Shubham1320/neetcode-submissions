class Solution:

    def getMaxOccurence(self, hashmap):
        maxVal = 0
        for key, value in hashmap.items():
            if value > maxVal:
                maxVal = value
        return maxVal

    def characterReplacement(self, s: str, k: int) -> int:
        
        max_len = 0
        i = 0
        j = 0

        hashmap = {}
        hashmap[s[i]] = 1
        relacements = 0

        while j != len(s):

            length = j - i + 1
            noOfReplacements = length - self.getMaxOccurence(hashmap)
            print(s[i])
            print(s[j])
            print(noOfReplacements)
            print(hashmap)
            if noOfReplacements <= k: 
                if length > max_len:
                    max_len = length

                j = j + 1
                if j == len(s):
                    break
                if s[j] in hashmap.keys():
                    hashmap[s[j]] = hashmap[s[j]] + 1
                else:
                    hashmap[s[j]] = 1
            else:
                hashmap[s[i]] = hashmap[s[i]] - 1
                i = i + 1
            
            print(max_len)

        return max_len
            