#include <vector>
#include <math.h>
#include <iostream>
#include <stack>
#include <stdio.h>

using namespace std;

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        vector<int>::iterator it = nums.begin();
        int num = 0;
        if(nums.size()==0){
            return 0;
        }
        if(nums.size()==1){
            return 1;
        }
        while(it<nums.end()){
            if(*it != *(it+1)){
                num = num+1;
            } else{
                it = nums.erase(it);
                continue;
            }
            it++;
        }
        return num;
    }
};

int main() {
    vector<int> test;
    test.push_back(1);
    test.push_back(2);
    Solution s;
    cout<<s.removeDuplicates(test);
    cout<<"sss"<<endl;
}