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
            validReps = 1
            while validReps * nums[i] <= self.target:
                validReps += 1
            k = 1
            combos1 = []
            while k != validReps:
                newItem = [nums[i]]*k
                combos1.append(newItem)
                print(combos1)

                if nums[i] * k == self.target:
                    self.ans.append(newItem)
                    k += 1
                    continue

                for item in self.combos:
                    itemSum = 0
                    for x in item:
                        itemSum += x
                    if (nums[i] * k) + itemSum == self.target:
                        self.ans.append(newItem+item)
                    elif (nums[i] * k) + itemSum < self.target:
                        combos1.append(newItem+item)
                k += 1
            self.combos = self.combos + combos1
            print(self.combos)

    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        self.target = target
        self.ans = []
        self.combos = []
        self.combine(nums, 0)
        return self.ans
                
                