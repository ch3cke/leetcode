#include <vector>
#include <iostream>
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

int main() {
    vector<int> nums,result;
    nums.push_back(3);
    nums.push_back(2);
    nums.push_back(4);
    int target = 6;
    Solution *test = new Solution();
    result = test->twoSum(nums,target);
    return 0;
}