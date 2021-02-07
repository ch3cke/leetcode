# @project : Python
# @author  : ch3ck
# @file   : 70.py
# @ide    : PyCharm
# @time   : 2021-02-08 00:03:48
# @contact: ch3cke@gmail.com
class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        a = 0
        b = 0
        c = 1
        for i in range(n):
            a = b
            b = c
            c = a+b
        return c