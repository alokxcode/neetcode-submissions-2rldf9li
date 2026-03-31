class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        index = 0
        returnList = []
        hash_table = {}
        my_set = set()
        while index < len(nums):
            startPointer = 0
            while startPointer != len(nums):
                if index == startPointer:
                    startPointer = startPointer + 1
                    continue
                diff = 0 - (nums[index] + nums[startPointer])
                if diff in hash_table:
                    sorted_list = sorted([nums[index],nums[startPointer],diff])
                    t = tuple(sorted_list)
                    my_set.add(t)
                    

                hash_table[nums[startPointer]] = startPointer
                startPointer = startPointer + 1
            index = index + 1
            hash_table = {}
        returnList = list(my_set)
        return returnList

        
        
        