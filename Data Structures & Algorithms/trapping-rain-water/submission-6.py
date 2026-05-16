
#0.  2.   0.   3.  1.  0.  1.  3.  2.  1. 







class Solution:
    def trap(self, height: List[int]) -> int:
        
        maxLeft = [0] * len(height)
        maxRight = [0] * len(height)
        leftMax = 0
        rightMax = 0

        i = 0
        while i != len(height):
            maxLeft[i] = leftMax
            if height[i] > leftMax:
                leftMax = height[i]
            i = i + 1
            
        j = len(height) - 1
        while j != -1:
            maxRight[j] = rightMax
            if height[j] > rightMax:
                rightMax = height[j]
            j = j - 1

        i = 0
        ans = 0
        while i != len(height):
            if min(maxRight[i], maxLeft[i]) - height[i] > 0:
                ans += min(maxRight[i], maxLeft[i]) - height[i]
            i = i + 1
        return ans