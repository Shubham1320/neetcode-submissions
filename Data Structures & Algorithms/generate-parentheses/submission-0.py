class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        
        ans = []

        def generate(string, opens, closes):
            
            if len(string) == 2 * n:
                ans.append(string)
                return
            
            if opens + 1 <= n:
                generate(string + '(', opens + 1, closes)
            
            if closes + 1 <= opens:
                generate(string + ')', opens, closes + 1)


        string = ''
        generate(string + '(', 1, 0)
    
        return ans