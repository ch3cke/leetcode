class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        mins = - 2 ** 31
        maxs = 2 ** 31 -1
        if(x<mins):
            return 0
        if(x>maxs):
            return 0
        if(x<0):
            q = -x
        else:
            q = x
        s = str(q)
        s = s[::-1]
        e = int(s)
        if(e<mins):
            return 0
        if(e>maxs):
            return 0
        if(x<0):
            return -e
        else:
            return e