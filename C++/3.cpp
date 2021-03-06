//
// Created by ch3ck on 2020/1/2.
//
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int max =0;
        int k = 0;
        for(int i=0;i<s.size();i++){
            for(int j = k;j<i;j++){
                if(s[i]==s[j]){
                    k=j+1;
                    break;
                }
            }
            if(i-k+1>max){
                max = i-k+1;
            }
        }
        return max;
    }
};
