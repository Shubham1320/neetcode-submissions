class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:

        stones = [-x for x in stones]

        heapq.heapify(stones)

        while True:
            if len(stones) < 2:
                break

            x = -heapq.heappop(stones)
            y = -heapq.heappop(stones)

            if x < y:
                y = y - x
                heapq.heappush(stones, -y)
            elif y < x:
                x = x - y
                heapq.heappush(stones, -x)
            
        stones = [-x for x in stones]

        if len(stones) == 0:
            return 0
        else:
            return stones[0]
                
