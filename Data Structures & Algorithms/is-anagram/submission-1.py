class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        hashmap = {}
        for i in s:
            try:
                if hashmap[i]:
                    hashmap[i] = hashmap[i] + 1
            except Exception as ex:
                hashmap[i] = 1

        print(hashmap)

        for i in t:
            try:
                if hashmap[i] > 0:
                    hashmap[i] = hashmap[i] - 1
                    if hashmap[i] == 0:
                        hashmap.pop(i)
                else:
                    return False 
            except Exception:
                return False

        print(hashmap.keys())

        if len(hashmap.keys()) == 0:
            return True
        return False