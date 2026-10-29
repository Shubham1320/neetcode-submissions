// Pseudo code - 
// 
// Input - nums -> integer array
//
// Hashmap -> store each element 
// 
// Time Complexity - O(n)
//
//
//

using namespace std;

class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {

        unordered_map<int, bool> mapA;

        int i;

        printf("%d",nums.size());

        // nums = [1, 2, 3, 3]

        //mapA = {
        //    1: 1
        //    2: 1
        //    3: 2
        //}

        for (i = 0; i < nums.size(); i++) {

            if (mapA.contains(nums[i])) {
                return true;
            }
            mapA[nums[i]] = true;
        }      
        return false;
    }
};