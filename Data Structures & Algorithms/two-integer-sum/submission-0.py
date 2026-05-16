class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for index, num in enumerate(nums):
            needed_val = target - num
            try:
                if hashmap[needed_val] >= 0:
                    print('here')
                    return [hashmap[needed_val], index]
            except Exception as ex:
                pass
            
            try:
                if hashmap[num]:
                    print('already exists')
            except Exception as ex:
                hashmap[num] = index

        return []