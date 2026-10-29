

//
// mapA<char, int>
//
// s = "racecar", t = "carrace"
//
// 
//
//
//

class Solution {
public:
    bool isAnagram(string s, string t) {

        unordered_map<char, int> mapA;

        int i;
        for (i = 0; i < s.size(); i++) {
            mapA[s[i]]++;
        }

        // mapA = {
        //
        //    r: 2
        //    a: 2
        //    c: 2
        //    e: 1
        //}

        // t = "carrace"

        // t = "shubham"


        // HashMap -> ops??
        // 1) storing data (insert)
        // 2) contains()
        // 3) reverse()
        // 4) iterate
        // 5) equals
        // 6) access
        // 7) delete

        int j;
        for (j = 0; j< t.size(); j++) {
            
            if (!mapA.contains(t[j])) {
                return false;
            }
            
            mapA[t[j]]--;

            if (mapA[t[j]] == 0) {
                mapA.erase(t[j]);
            }
        }

        if (mapA.size() == 0) {
            return true;
        }

        return false;
    }
};
