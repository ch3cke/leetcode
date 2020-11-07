#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        if(nums.size()>2){
            sort(nums.begin(), nums.end());
            int flag = 0;
            int begin = 0;
            while(nums[begin]<=0&&begin<nums.size()-1){
                if(begin>0&&nums[begin]==nums[begin-1]){
                    begin+=1;
                    continue;
                }
                int ends = nums.size()-1;
                int index = begin+1;
                while (index<ends){
                    if(nums[index]+nums[ends]+nums[begin]==0){
                        result.push_back({nums[begin],nums[index],nums[ends]});
                        index +=1;
                    } else if(nums[index]+nums[ends]+nums[begin]>0){
                        ends = ends -1;
                    } else if (nums[index]+nums[ends]+nums[begin]<0){
                        index = index +1;
                    }
                }
                begin+=1;
            }
        }
        result.erase(unique(result.begin(), result.end()), result.end());
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
    //[-1, 0, 1, 2, -1, -4]
//    info.push_back(1);
//    info.push_back(-1);
//    info.push_back(0);
    info.push_back(0);
    info.push_back(0);
    info.push_back(0);
    result = s.threeSum(info);
    showVectors(result);
    return 0;
}