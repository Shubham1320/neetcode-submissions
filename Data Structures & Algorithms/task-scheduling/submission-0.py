class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        
        freqHashMap = {}
        countHashMap = {}

        for task in tasks:
            if task in freqHashMap.keys():
                freqHashMap[task] += 1
            else:
                freqHashMap[task] = 1
                countHashMap[task] = 0

        print(freqHashMap)
        print(countHashMap)
        print('\n')

        cycles = 0

        while len(freqHashMap.values()) != 0:

            heap = [(-value, key) for key, value in freqHashMap.items()]
            heapq.heapify(heap)


            top = heapq.heappop(heap)

            while top[1] in countHashMap.keys() and countHashMap[top[1]] > 0:
                if len(heap) == 0:
                    top = None
                    break
                
                top = heapq.heappop(heap)

            for key, val in countHashMap.items():
                countHashMap[key] -= 1
                val -= 1

            print(countHashMap)

            if top:
                freqHashMap[top[1]] -= 1
                if freqHashMap[top[1]] == 0:
                    del freqHashMap[top[1]]
                countHashMap[top[1]] = n

            print(freqHashMap)
            print(countHashMap)
            print('\n')

            cycles += 1

        return cycles



