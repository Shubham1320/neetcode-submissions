

#min = 1 
#max = 7

#10,1,5,6,7,1
#           i j
    

#10,1,5,6,7,1

#min = 1

#0   0.  4.  5.  6.  6 


class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        profit = [0] * len(prices)
        minimum = prices[0]
        i = 1

        while i != len(prices):

            if prices[i] > minimum:
                profit[i] = max(prices[i]-minimum, profit[i-1])
            else:
                minimum = prices[i]
                profit[i] = profit[i-1]

            i = i + 1

        return profit[len(prices) - 1]