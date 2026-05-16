class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        arr1 = [1] * len(nums)
        index = 1
        while index != len(nums):
            arr1[index] = nums[index-1]*arr1[index-1]
            index = index+1

        arr2 = [1] * len(nums)
        index = len(nums) - 2
        while index != -1:
            arr2[index] = nums[index+1]*arr2[index+1]
            index = index - 1

        print(arr1)
        print(arr2)

        index = 0
        ans = [1] * len(nums)
        while index != len(nums):
            ans[index] = arr1[index]*arr2[index]
            index = index + 1
            print(ans)

        return ans










