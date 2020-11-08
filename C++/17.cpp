//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法
// 👍 985 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    vector<string> letterCombinations(string digits) {
        vector<string> resultStr;
        if(digits.size()==0){
            return resultStr;
        }
        map<char,string> table{
                {'0', " "}, {'1',"*"}, {'2', "abc"},
                {'3',"def"}, {'4',"ghi"}, {'5',"jkl"},
                {'6',"mno"}, {'7',"pqrs"},{'8',"tuv"},
                {'9',"wxyz"}};
        setStr(resultStr, digits, 0, "", table);
        return resultStr;
    }

    void setStr(vector<string>& result, string digit, int m, string tmpstr, map<char, string> table){
        if (tmpstr.size() == digit.size()){
            result.push_back(tmpstr);
            return;
        }
        string tmptable = table[digit[m]];
        for(char w : tmptable){
            tmpstr += w;
            setStr(result,digit,m+1,tmpstr,table);
            tmpstr.pop_back();
        }
    }
};
//leetcode submit region end(Prohibit modification and deletion)
