//
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// Created by ch3ck on 2019/12/31.
//
#include <vector>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int>::iterator it,is;
        vector<int> result;
        int i=0;
        int j=1;
        for(it=nums.begin();it!=nums.end();it++){
            for(is=it+1;is!=nums.end();is++){
                if(*is+*it==target){
                    result.push_back(i);
                    result.push_back(j);
                }
                j++;
            }
            i++;
            j=i+1;
        }
        return result;
    }
};