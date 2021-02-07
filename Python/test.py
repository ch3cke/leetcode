# @project : Python
# @author  : ch3ck
# @file   : test.py
# @ide    : PyCharm
# @time   : 2020-11-08 16:10:00
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

if __name__ == '__main__':
    s = Solution()
    print s.simplifyPath('/../')