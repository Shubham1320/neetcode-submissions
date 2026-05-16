import copy

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        
        ans = []

        def genSubset(nums, i):


            if len(nums) - 1 == i:
                ans.append([nums[i]])
                return

            genSubset(nums, i + 1)

            ans1 = []
            ans1.append([nums[i]])
            for item in ans:
                new_item = copy.deepcopy(item)
                print(new_item)
                new_item.append(nums[i])
                ans1.append(new_item)

            for item in ans1:
                ans.append(item)

        genSubset(nums, 0)
        ans.append([])
        print(ans)
        return ans
         