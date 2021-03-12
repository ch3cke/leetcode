# @project : Python
# @author  : ch3ck
# @file   : 227.py
# @ide    : PyCharm
# @time   : 2021-03-12 17:26:34
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

