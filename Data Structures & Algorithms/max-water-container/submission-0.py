class Solution:
    def maxArea(self, heights: List[int]) -> int:

        index = 0
        area = 0
        while index != len(heights):
            pointer = 0
            while pointer != len(heights):
                if heights[index] < heights[pointer]:
                    height = heights[index]
                else:
                    height = heights[pointer]
                if index > pointer:
                    width = index - pointer
                else:
                    width = pointer - index
                a = height * width
                if a > area :
                    area = a
                pointer = pointer + 1
            index = index + 1
        
        return area



        