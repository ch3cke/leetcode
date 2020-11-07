//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
// 👍 622 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
//leetcode submit region end(Prohibit modification and deletion)
