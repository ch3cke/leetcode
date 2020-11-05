#include <iostream>
#include <vector>
using namespace std;

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
int main(int argc, char** argv){
    Solution s;
    vector<string> info;
    info.push_back("");
    info.push_back("");
    info.push_back("hack");
    cout<<s.longestCommonPrefix(info)<<endl;
    return 0;
}