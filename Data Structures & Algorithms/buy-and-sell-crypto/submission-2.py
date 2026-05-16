

#min = 1 
#max = 7

#10,1,5,6,7,1
#           i j
    


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        i = 0
        j = 1
        ans = 0
        minimum = prices[i]
        maximum = 0

        while j != len(prices):
            if prices[j] <= prices[i]:
                if prices[j] <= minimum:
                    i = j
                    j = i + 1
                    if maximum - minimum > ans:
                        ans = maximum - minimum
                    minimum = prices[i]
                    maximum = 0
                else:
                    j = j + 1
            else: 
                if prices[j] > maximum:
                    maximum = prices[j]
                j = j + 1 

        if ans < 0:
            return 0

        if maximum - minimum > ans:
            return maximum - minimum

        return ans
