class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        
        hashmap = defaultdict()

        index = 1
        for num in numbers:
            req = target - num
            try:
                if hashmap[req]:
                    return [hashmap[req], index]
            except KeyError as ke:
                hashmap[num] = index
            index = index + 1
