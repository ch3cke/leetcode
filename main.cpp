#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int Max = 0;
        int tmpArea = 0;
        int j = height.size()-1;
        int i = 0;
        while(j != i){
            tmpArea = getArea(i, j, height);
            if (tmpArea > Max){
                cout<<i<<'\t'<<j<<tmpArea<<endl;
                Max = tmpArea;
            }
            if(height[i]>height[j]){
                j--;
            } else{
                i++;
            }
        }

        return Max;
    }
    int getArea(int i, int j, vector<int>& height){
        if (height[j] > height[i]){
            return (j-i) * height[i];
        } else{
            return (j-i) * height[j];
        }
    }
};
int main(int argc, char** argv){
    Solution s;
    vector<int> height;
    height.push_back(2);
    height.push_back(3);
    height.push_back(4);
    height.push_back(5);
    height.push_back(18);
    height.push_back(17);
    height.push_back(6);
//    height.push_back(3);
//    height.push_back(7);
    printf("%d",s.maxArea(height));
    return 0;
}