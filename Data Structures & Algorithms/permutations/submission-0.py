class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        
        res = []

        def generate(nums, ans):
            print(ans)
            if len(ans) == len(nums):
                res.append(ans)

            i = 0
            while i != len(nums):
                if nums[i] not in ans:
                    generate(nums, ans + [nums[i]])
                i += 1
        
        generate(nums, [])
        return res
