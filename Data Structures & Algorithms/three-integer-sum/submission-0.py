class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        hashmap = defaultdict()
        i = 0
        while i != len(nums):
            j = i + 1
            while j != len(nums):
                twosum = nums[i] + nums[j]
                try:
                    if hashmap[twosum]:
                        hashmap[twosum].append([i,j]) 
                except KeyError as ke:
                    hashmap[twosum] = [[i,j]]
                j = j + 1
            i = i + 1
        print(hashmap)

        ans = []
        for index, num in enumerate(nums):
            req = -1 * num
            try:
                if hashmap[req]:
                    for item in hashmap[req]:
                        if index in item:
                            continue
                        out = [num, nums[item[0]], nums[item[1]]]
                        out = sorted(out)
                        ans.append(tuple(out))
            except KeyError as ke:
                pass

        print(ans)
        print(set(ans))

        return [[item[0], item[1], item[2]] for item in set(ans)]