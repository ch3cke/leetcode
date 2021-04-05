# @project : Python
# @author  : ch3ck
# @file   : test.py
# @ide    : PyCharm
# @time   : 2020-11-08 16:10:00
# @contact: ch3cke@gmail.com
class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        i = 0
        nums = []
        ops = []
        while i<len(s):
            if '0' <= s[i] <= '9':
                tmp = 0
                while '0' <= s[i] <= '9':
                    tmp = tmp*10+int(s[i])
                    i+=1
                    if i>=len(s):
                        break
                nums.append(tmp)
            elif s[i]==' ':
                i+=1
                continue
            else:
                if len(ops)==0:
                    ops.append(s[i])
                    i+=1
                else:
                    if self.check(ops[-1], s[i]) or (len(ops)==0):
                        ops.append(s[i])
                        i=i+1
                    else:
                        while not self.check(ops[-1], s[i]):
                            sig = ops.pop()
                            num1 = nums.pop()
                            num2 = nums.pop()
                            if sig=='+':
                                nums.append(num1+num2)
                            elif sig =='-':
                                nums.append(num2-num1)
                            elif sig == '*':
                                nums.append(num2*num1)
                            elif sig == '/':
                                nums.append(num2/num1)
                            if len(ops)==0:
                                ops.append(s[i])
                                i+=1
                                break
        while len(ops)>0:
            sig = ops.pop()
            num1 = nums.pop()
            num2 = nums.pop()
            if sig=='+':
                nums.append(num1+num2)
            elif sig =='-':
                nums.append(num2-num1)
            elif sig == '*':
                nums.append(num2*num1)
            elif sig == '/':
                nums.append(num2/num1)
        return nums[0]

    def check(self,s1, s2):
        if (s1=='+' or s1=='-') and  (s2=='*' or s2=='/'):
            return True
        else:
            return False

    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        tmp=[]
        preorder = preorder.split(',')
        i = len(preorder)-1
        if preorder[i]!='#':
            return False
        while i>=0:
            if preorder[i]=='#':
                tmp.append(preorder[i])
            else:
                tmp.pop()
                if len(tmp)==0:
                    return False
            i-=1
        if len(tmp)==1 and tmp[0]=='#':
            return True
        else:
            return False

    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        s = ''.join(x for x in s if x.isalpha())
        if s=='':
            return True
        if len(s)==1:
            return False
        q = len(s)-1
        p = 0
        while q > p:
            if s[p].lower()!=s[q].lower():
                return False
            else:
                p +=1
                q -=1
        return True

    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        s = s.strip()

        tmp = s.split(' ')
        result =""
        i = len(tmp)-1
        while i>=0:
            if len(tmp[i])==0:
                i-=1
                continue
            else:
                result += tmp[i]
                result += " "
                i-=1
        return result.strip()



if __name__ == '__main__':
    s = Solution()
    print s.reverseWords("  Bob    Loves  Alice   ")