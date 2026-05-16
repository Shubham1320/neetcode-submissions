class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        
        self.heap = []

        for item in nums:
            self.heap.append(-item)

        self.k = k
        heapq.heapify(self.heap)


    def add(self, val: int) -> int:

        heapq.heappush(self.heap, -val)
        heapq.heapify(self.heap)

        print(self.heap)

        i = 0 
        arr = []
        ans = 0

        while i < self.k:
            ans = -heapq.heappop(self.heap)
            print(ans)
            arr.append(ans)
            i += 1

        for item in arr:
            heapq.heappush(self.heap, -item)
            heapq.heapify(self.heap)


        return ans

