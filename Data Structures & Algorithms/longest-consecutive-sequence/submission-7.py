
#{
#    20: 1
#    4: 4
#    10: 1 

#}

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        hashmap = {}
        for num in nums:
            try:
                if hashmap[num-1] is not None and num > hashmap[num-1]:
                    hashmap[num-1] = num
            except KeyError as ke:
                try:
                    if hashmap[num]:
                        print('already present')
                except KeyError as e:
                    hashmap[num] = num
            #print(hashmap)

        print(hashmap)

        for key, val in hashmap.items():
            while True :
                try:
                    if val == key:
                        break
                    if val and hashmap[val] is None:
                        break
                    elif val is None:
                        break
                    else:
                        hashmap[key] = hashmap[val]
                        val1 = val
                        val = hashmap[val]
                        hashmap[val1] = None
                except KeyError as ke:
                    break

        print(hashmap)

        for key, val in hashmap.items():
            while True :
                try:
                    if val and hashmap[val+1] is None:
                        break
                    elif val is None:
                        break
                    else:
                        hashmap[key] = hashmap[val+1]
                        val1 = val
                        val = hashmap[val+1]
                        hashmap[val1+1] = None
                    print(val)
                    print(hashmap)
                except KeyError as ke:
                    break


        print(hashmap)

        ans = 0
        for key, val in hashmap.items():
            if val is None:
                continue
            if val - key + 1 > ans:
                ans = val - key + 1

        return ans

