#include <vector>
#include <math.h>
#include <iostream>
using namespace std;


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

int main() {
    string s = "aab";
    Solution test;
    cout<<test.lengthOfLongestSubstring(s);
    return 0;
}