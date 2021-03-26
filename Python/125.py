# @project : Python
# @author  : ch3ck
# @file   : 125.py
# @ide    : PyCharm
# @time   : 2021-03-23 18:55:55
# @contact: ch3cke@gmail.com
class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        s = ''.join(x for x in s if x.isalnum())
        if s=='':
            return True
        q = len(s)-1
        p = 0
        while q > p:
            if s[p].lower()!=s[q].lower():
                return False
            else:
                p +=1
                q -=1
        return True