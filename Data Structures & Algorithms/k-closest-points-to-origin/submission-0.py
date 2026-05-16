class Solution:

    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:

        hashmap = {}
        heap = []

        for point in points:
            xsqr = point[0] * point[0]
            ysqr = point[1] * point[1]
            distance = pow(xsqr+ysqr, 0.5)
            print(distance)

            if distance not in hashmap.keys():
                hashmap[distance] = [point]
            else:
                hashmap[distance].append(point)
            heapq.heappush(heap, distance)
            heapq.heapify(heap)

        i = 0
        ans = []
        while i < k:
            distance = heapq.heappop(heap)

            if len(hashmap[distance]) == 1:
                ans.append(hashmap[distance][0])
                del hashmap[distance]
            else:
                ans.append(hashmap[distance][0])
                hashmap[distance].pop(0)
            i+=1

        return ans
            