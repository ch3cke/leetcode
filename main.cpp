#include <iostream>
#include <stdio.h>//printf
#include <vector>
#include <algorithm>
#include <map>
using namespace std;

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