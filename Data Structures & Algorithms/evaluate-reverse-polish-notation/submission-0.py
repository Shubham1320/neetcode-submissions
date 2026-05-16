class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        operators = ['*', '/', '+', '-']
        for token in tokens:
            if token not in operators:
                stack.append(int(token))
            else:
                second_num = stack.pop()
                first_num = stack.pop()
                if token == '+':
                    ans = first_num + second_num
                elif token == '-':
                    ans = first_num - second_num
                elif token == '*':
                    ans = first_num * second_num
                elif token == '/':
                    ans = int(first_num / second_num)
                stack.append(ans)

        return stack[0]