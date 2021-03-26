#include <iostream>
#include <string>
#include <stack>
#include <vector>

using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int money = 0;
        int i=1;
        while(i<prices.size()){
            if(prices[i]-prices[i-1]>0){
                money+= (prices[i]-prices[i-1]);
            }
            i+=1;
        }
        return money;
    }
};
int main(){
    vector<int>nums;
    nums.push_back(2);
    nums.push_back(4);
    nums.push_back(1);

    Solution* s = new Solution();
    cout<<s->maxProfit(nums);
}
