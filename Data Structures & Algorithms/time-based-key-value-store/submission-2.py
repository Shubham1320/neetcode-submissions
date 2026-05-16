class TimeMap:

    def __init__(self):
        self.data = {} 

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.data.keys():
            self.data[key] = [[timestamp, value]]
        else:
            self.data[key].append([timestamp, value])

        print(self.data)


    def get(self, key: str, timestamp: int) -> str:
        if key not in self.data.keys():
            return ""
        
        values = self.data[key]

        def find(values, i, j):


            mid = i + (j-i)//2

            print(i,j, mid)

            if j < i:
                return ""

            if timestamp >= values[mid][0]:
                if i == j:
                    return values[mid][1]
                elif timestamp < values[mid+1][0]:
                    return values[mid][1]

            if timestamp > values[mid][0]:
                return find(values, mid+1, j)
            elif timestamp < values[mid][0]:
                return find(values, i, mid-1)
        
        ans = find(values, 0, len(values)-1)
        print('\n')
        return ans

