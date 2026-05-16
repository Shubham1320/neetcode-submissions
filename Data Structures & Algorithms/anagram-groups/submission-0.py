class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hashmap = {}
        for index, string in enumerate(strs):
            #print(string)
            #print(sorted(string))
            sorted_str = "".join(sorted(string))
            try:
                if hashmap[sorted_str]:
                    hashmap[sorted_str].append(string)
            except KeyError as ke:
                hashmap[sorted_str] = [string]
            print(hashmap)

        ans = []
        index = 0
        
        for key, value in hashmap.items():
            ans.insert(index, value)
            print(ans)
            index = index + 1
        return ans