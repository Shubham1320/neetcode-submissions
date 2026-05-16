class MinStack:

    def __init__(self):
        self.stack = []
        self.minimum = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        if len(self.minimum) == 0:
            self.minimum.append(val)
        elif val < self.minimum[-1]:
            self.minimum.append(val)
        else:
            self.minimum.append(self.minimum[-1])  

    def pop(self) -> None:
        val = self.stack.pop()
        self.minimum.pop()

    def top(self) -> int:
        if len(self.stack) != 0:
            return self.stack[-1]

    def getMin(self) -> int:
        if len(self.minimum) != 0:
            return self.minimum[-1]

