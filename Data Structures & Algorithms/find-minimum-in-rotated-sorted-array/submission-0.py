class Solution:
    def findMin(self, nums: List[int]) -> int:
        
        def find(nums, i, j):

            mid = i + int((j-i)/2)

            if i == j:
                return nums[mid]

            if nums[mid] < nums[mid-1]:
                return nums[mid]

            if nums[mid] < nums[j]:
                return find(nums, i, mid-1)
            else:
                return find(nums, mid+1, j)

        return find(nums, 0, len(nums)-1)