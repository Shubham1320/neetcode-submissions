




class Solution:
    def maxArea(self, heights: List[int]) -> int:

        i = 0
        j = len(heights) - 1

        ans = 0
        while i < j:
            area = min(heights[i], heights[j]) * (j-i)
            ans = max(area, ans)           
            if heights[i] <= heights[j]:
                i = i + 1
            else:
                j = j - 1

        return ans