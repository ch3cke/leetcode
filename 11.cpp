//
// Created by ch3ck on 2020/11/5.
//

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
