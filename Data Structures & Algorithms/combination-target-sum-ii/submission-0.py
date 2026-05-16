class Solution:

    ans = []
    combos = []

    def combine(self, nums, i):

        if len(nums) == i:
            return

        self.combine(nums, i + 1)

        if nums[i] == self.target:
            self.ans.append([nums[i]])
        elif nums[i] < self.target:
            combos1 = []
            combos1.append([nums[i]])
            for item in self.combos:
                itemSum = 0
                for x in item:
                    itemSum += x
                if nums[i] + itemSum == self.target:
                    self.ans.append([nums[i]]+item)
                elif nums[i] + itemSum < self.target:
                    combos1.append([nums[i]]+item)
            self.combos = self.combos + combos1
            print(self.combos)

    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        self.target = target
        self.ans = []
        self.combos = []
        self.combine(candidates, 0)
        res = list(set(tuple(sorted(sub)) for sub in self.ans))
        res = [list(x) for x in res]
        return res
