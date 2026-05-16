class Solution:
    ans = []

    def generate(self, nums, i):
        if i == len(nums) - 1:
            self.ans.append([nums[i]])
            return

        self.generate(nums, i+1)

        ans1 = []
        ans1.append([nums[i]])
        for item in self.ans:
            ans1.append(item + [nums[i]])
            
        self.ans = self.ans + ans1

    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        

        self.ans = []
        self.generate(nums, 0)
        self.ans.append([])

        res = list(set(tuple(sorted(x)) for x in self.ans))
        res = [list(x) for x in res]

        return res