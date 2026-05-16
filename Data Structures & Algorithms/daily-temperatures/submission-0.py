class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        
        stack = []

        ans = [0] * len(temperatures)

        for index, item in enumerate(temperatures):
            if len(stack) == 0:
                stack.append((index, item))
            elif item < stack[-1][1]:
                stack.append((index, item))
            else:
                while len(stack) != 0 and stack[-1][1] < item:
                    top_idx, top_val = stack[-1]
                    ans[top_idx] = index-top_idx 
                    stack.pop()
                stack.append((index, item))

        return ans
  
