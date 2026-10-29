
//
// nums = [3,4,5,6], target = 7
//
//. mapA -> key = num, value = index
//
// for (i = 0; i < nums.size(); i++) {
//      diff = target - nums[i] // diff = 4
//.     
//      if (mapA.contains(diff)) {
//            return {i, mapA[diff]}
//        }
//      
//        mapA[nums[i]] = i
//
//.     check if diff exists in the array, reuturn current index, that index
//
// }
//
//
//
//
//

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {

        unordered_map<int, int> mapA;

        int i = 0;
        int diff = 0;

        for (i = 0; i < mapA.size(); i++) {

            diff = target - nums[i];

            if (mapA.contains(diff)) {
                
                if (i < mapA[diff]) {
                    return {i, mapA[diff]};
                }
                return {mapA[diff], i};
            }

            mapA[nums[i]] = i;
        }
    }
};
