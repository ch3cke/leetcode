//
// Created by ch3ck on 2020/1/17.
//
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    bool isValid(string s) {
        stack<char> tool;
        for(int i = 0;i<s.size();i++){
            if(tool.empty()){
                tool.push(s[i]);
            } else if(check(tool.top(),s[i])){
                tool.pop();
            } else{
                tool.push(s[i]);
            }
        }
        if(tool.empty()){
            return true;
        } else{
            return false;
        }
    }
    bool check(char s1,char s2){
        if((s1=='{' && s2=='}')||(s2=='{' && s1=='}')){
            return true;
        } else if((s1=='(' && s2==')')||(s2=='(' && s1==')')){
            return true;
        } else if((s1=='[' && s2==']')||(s2=='['&& s1==']')){
            return true;
        } else {
            return false;
        }
    }
};
