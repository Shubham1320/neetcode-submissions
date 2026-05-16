class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        
        stack = []
        maxArea = 0

        stack.append([0, heights[0]])
        i = 1
        while i != len(heights):
            if heights[i] >= stack[-1][1]:
                stack.append([i, heights[i]])
            else:
                while stack and heights[i] < stack[-1][1]:
                    item = stack.pop()
                    maxArea = max(maxArea, (i-item[0]) * item[1])
                stack.append([item[0], heights[i]])
            i = i + 1
            print(stack)
            print(maxArea)

        while len(stack) != 0:
            item = stack.pop()
            maxArea = max(maxArea, (len(heights)-item[0]) * item[1])

        return maxArea

