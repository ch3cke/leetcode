# @project : Python
# @author  : ch3ck
# @file   : 71.py
# @ide    : PyCharm
# @time   : 2021-02-08 00:46:14
# @contact: ch3cke@gmail.com
class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        tmp = path.split('/')
        for i in range(len(tmp)):
            tmp[i] += '/'
        result_tmp = []
        result_tmp.append(tmp[0])
        for i in range(1,len(tmp)):
            if tmp[i]=='../':
                if len(result_tmp)>1:
                    result_tmp.pop()
            elif tmp[i] == '/':
                continue
            elif tmp[i] == './':
                continue
            else:
                result_tmp.append(tmp[i])
        result = ''
        for i in result_tmp:
            result += i
        if result[-1]=='/' and len(result)>1:
            return result[:-1]
        else:
            return result