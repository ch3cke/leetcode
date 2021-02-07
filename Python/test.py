# @project : Python
# @author  : ch3ck
# @file   : test.py
# @ide    : PyCharm
# @time   : 2020-11-08 16:10:00
# @contact: ch3cke@gmail.com
class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        for i in range(x):
            if (i+1) * (i+1) > x:
                return i
if __name__ == '__main__':
    s = Solution()
    print s.mySqrt(4)