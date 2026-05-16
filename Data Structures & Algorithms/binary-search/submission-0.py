class Solution:
    def search(self, nums: List[int], target: int) -> int:
        
        def find(nums, i, j):

            print(i, j)
            if j < i:
                return -1

            mid = i + int((j - i)/2)
            print(mid)
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                return find(nums, mid+1, j)
            else:
                return find(nums, i, mid-1)
            
        return find(nums,0, len(nums)-1)
            