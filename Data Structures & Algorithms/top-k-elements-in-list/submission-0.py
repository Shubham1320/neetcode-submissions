class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        
        hashmap = {}
        maxOccur = 0

        for num in nums:
            try:
                if hashmap[num]:
                    hashmap[num] = hashmap[num] + 1
            except KeyError as ke:
                hashmap[num] = 1

        new_hashmap = {}

        for key, value in hashmap.items():
            try:
                if new_hashmap[value]:
                    new_hashmap[value].append(key)
            except KeyError as ke:
                new_hashmap[value] = [key]
                if value > maxOccur:
                    maxOccur = value

        ans = []
        while k != 0:
            try:
                if new_hashmap[maxOccur]:
                    for i in new_hashmap[maxOccur]:
                        ans.append(i)
                        k = k - 1
                        if k == 0:
                            break
            except KeyError as ke:
                pass
            finally:
                maxOccur = maxOccur - 1

        if k != 0:
            return None

        return ans