# @project : Python
# @author  : ch3ck
# @file   : test.py
# @ide    : PyCharm
# @time   : 2020-11-08 16:10:00
# @contact: ch3cke@gmail.com
class Solution(object):
    def isInterleave(self, s1, s2, s3):
        """
        :type s1: str
        :type s2: str
        :type s3: str
        :rtype: bool
        """
        tmp = ''
        for i in s1:
            for j in s3:
                if i == j:
                    break
        print s3

if __name__ == '__main__':
    s = Solution()
    print s.isInterleave("aabcc","dbbca","aadbbcbcac")