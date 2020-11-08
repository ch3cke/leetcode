//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if(strs.size()==0){
            return "";
        } else if (strs.size()==1){
            return strs[0];
        } else{
            int i = 0;
            string tmpStr = "";
            tmpStr = getString(strs[0],strs[1]);
            for(i = 2;i<strs.size();i++){
                tmpStr = getString(tmpStr,strs[i]);
            }
            return tmpStr;
        }
    }
    string getString(string s1, string s2){
        string resultStr = "";
        int i = 0;
        if(s1.size()==0||s2.size()==0){
            return "";
        }
        while (s1[i] == s2[i]){
            resultStr.push_back(s1[i]);
            i++;
            if(i>(s1.length()-1)||i>(s2.length()-1)) break;
        }
        return resultStr;
    }
};
//leetcode submit region end(Prohibit modification and deletion)
