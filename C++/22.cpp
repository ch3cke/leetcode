//
// Created by ch3ck on 2020/3/6.
//

//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
// 例如，给出 n = 3，生成结果为：
//
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
//
// Related Topics 字符串 回溯算法


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> results;
        def("",results,n,n);
        return results;
    }

    void def(string result, vector<string> &results, int right, int left){
        if(left == 0 && right == 0){
            results.push_back(result);
            return;
        }
        if(left>0){
            def(result+'(',results, right, left-1);
        }

        if(right>left){
            def(result+')',results, right-1, left);
        }
    }
};
//leetcode submit region end(Prohibit modification and deletion)
