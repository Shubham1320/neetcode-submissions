class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:

        def isValidRate(piles, rate):
            hrs = 0
            for pile in piles:
                hrs += math.ceil(pile/rate)
                if hrs > h:
                    return False
            return True

        def binarySearch(piles, i, j):
            
            print(i,j)
            if j < i:
                return -1

            mid = i + int((j-i)/2)

            if isValidRate(piles, mid):
                print(mid)
                ans = binarySearch(piles, i, mid-1)
                if ans == -1:
                    return mid
                return ans

            else:
                return binarySearch(piles, mid+1, j)

        return binarySearch(piles, 1, max(piles))

                    