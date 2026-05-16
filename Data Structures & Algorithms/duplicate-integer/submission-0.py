class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash_map = {}
        for num in nums:
            try:
                if hash_map[num]:
                    return True
            except Exception as ex:
                hash_map[num] = True
        return False
