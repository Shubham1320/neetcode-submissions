class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        ans = []
        dq = deque()
        i = 0
        j = 0

        while j < len(nums):
            while dq and nums[dq[-1]] < nums[j]:
                dq.pop()
            dq.append(j)

            if i > dq[0]:
                dq.popleft()

            if j + 1 >= k:
                ans.append(nums[dq[0]])
                i = i + 1
            j = j + 1 

        return ans

