# @project : Python
# @author  : ch3ck
# @file   : 331.py
# @ide    : PyCharm
# @time   : 2021-03-12 19:58:52
# @contact: ch3cke@gmail.com
# 序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
#
#       _9_
#     /   \
#    3     2
#   / \   / \
#  4   1  #  6
# / \ / \   / \
# # # # #   # #
#
#
#  例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
#
#  给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
#
#  每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
#
#  你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
#
#  示例 1:
#
#  输入: ",,#"
# 输出: true
#
#  示例 2:
#
#  输入: "1,#"
# 输出: false
#
#
#  示例 3:
#
#  输入: "9,#,#,1"
# 输出: false
#  Related Topics 栈
#  👍 260 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
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
# leetcode submit region end(Prohibit modification and deletion)
