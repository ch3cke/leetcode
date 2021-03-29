# @project : Python
# @author  : ch3ck
# @file   : 137.py
# @ide    : PyCharm
# @time   : 2021-03-28 14:49:15
# @contact: ch3cke@gmail.com
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return  (3 * sum(set(nums))-sum(nums))/2