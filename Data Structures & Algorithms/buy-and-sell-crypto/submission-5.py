

#min = 1 
#max = 7

#10,1,5,6,7,1
#           i j
    

#10,1,5,6,7,1

#min = 1

#0   0.  4.  5.  6.  6 


class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        max_profit = 0
        i = 0
        j = 1

        while j != len(prices):
            if prices[j] > prices[i]:
                if prices[j] - prices[i] > max_profit:
                    max_profit = prices[j] - prices[i]
                j = j + 1
            else:
                i = j
                j = i + 1

        return max_profit