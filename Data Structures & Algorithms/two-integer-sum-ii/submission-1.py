class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left = 0
        length = len(numbers)
        right = length - 1

        while left < right:
            total = numbers[left] + numbers[right]
            if total == target:
                return [left+1,right+1]
            elif total > target:
                right -= 1
            else:
                left += 1


