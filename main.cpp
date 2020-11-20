#include <iostream>
#include <vector>
#include <algorithm>
#include <map>
using namespace std;

class Solution {
public:
    vector<vector<int>> four(vector<int>& nums, int target) {
        vector<vector<int>> result;
        if(nums.size()>3){
            sort(nums.begin(), nums.end());
            int first, second, third, fourth;
            first = 0;
            while (first<nums.size()-3){
                while (first>0&&nums[first]==nums[first-1]){
                    first = first +1;
                }
                second = first+1;
                while (second< nums.size()-2){
                    while (second>first+1&&nums[second]==nums[second-1]){
                        second = second+1;
                    }
                    third = second+1;
                    fourth = nums.size()-1;
                    while (third < fourth){
                        int tmpnum = nums[first]+nums[second]+nums[third]+nums[fourth];
                        if (tmpnum==target){
                            result.push_back({nums[first],nums[second],nums[third], nums[fourth]});
                            third +=1;
                        }
                        else if(tmpnum> target){
                            fourth = fourth-1;
                            while (fourth>third&&nums[fourth]==nums[fourth+1]){
                                fourth = fourth-1;
                            }
                        } else if(tmpnum<target){
                            third = third+1;
                            while (third<fourth&&nums[third]==nums[third-1]){
                                third = third+1;
                            }
                        }
                    }
                    second += 1;
                }
                first += 1;
            }

        }
        result.erase(unique(result.begin(), result.end()), result.end());
        return result;
    }
};

void showVector(vector<int > num){
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
    srand(0x1548u);
    for(int i = 0; i <=9; i++){
        cout<<rand()%127<<endl;
    }
}