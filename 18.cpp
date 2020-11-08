//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
// Related Topics 数组 哈希表 双指针
// 👍 667 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
//leetcode submit region end(Prohibit modification and deletion)
