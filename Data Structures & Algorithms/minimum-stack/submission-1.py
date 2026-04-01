class MinStack:

    def __init__(self):
        self.stack = []
        

    def push(self, val: int) -> None:
        self.stack.append(val)

        

    def pop(self) -> None:
        self.stack[-1] == None
        self.stack = self.stack[:-1]
        

    def top(self) -> int:
        return self.stack[-1]
        

    def getMin(self) -> int:
        min = self.stack[-1]
        for i,v in enumerate(self.stack):
            if v <= min :
                min = v
        return min