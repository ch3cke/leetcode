#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());
        int first, second, third;
        int iv;
        int tmp=100000;
        int result = 0;
        int sum;
        for(first = 0; first < nums.size();first++){
            if(first>0&& nums[first]==nums[first-1]) continue;
            third = nums.size()-1;
            second = first+1;
            while (third!=second&& second<nums.size()){
                if(sum>target&&nums[third]==nums[third-1]){
                    third=third-1;
                    continue;
                }
                sum = nums[first]+nums[second]+nums[third];
                iv  = (sum-target)*(sum-target);
                if(iv<tmp){
                    tmp = iv;
                    result = sum;
                }
                if(sum>target){
                    third = third-1;
                } else{
                    second = second +1;
                }
            }
        }
        return result;
    }
};

void showVector(vector<int> num){
    for(int i  = 0; i < num.size();i++){
        cout<<num[i]<<"\t";
    }
}

void showVectors(vector<vector<int>> nums){
    for(int i  = 0; i < nums.size();i++){
        showVector(nums[i]);
    }
}

int main(int argc, char** argv){
    Solution s;
    vector<int> info;
    vector<vector<int>> result;
    // [-1,2,1,-4]
//    info.push_back(1);
//    info.push_back(-1);
    info.push_back(-1);
    info.push_back(2);
    info.push_back(1);
    info.push_back(-4);
    cout<<s.threeSumClosest(info,1);
    return 0;
}