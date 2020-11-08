//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2728 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
//leetcode submit region end(Prohibit modification and deletion)
