#include <vector>
#include <math.h>
#include <iostream>
using namespace std;


class Solution {
public:
    int romanToInt(string s) {
        int result = 0;
        for(int i =0;i<s.size();i++){
            if(s[i]=='I'){
                if(s[i+1]=='V'){
                    result = result+4;
                    i++;
                } else if(s[i+1]=='X'){
                    result = result+9;
                    i++;
                } else{
                    result = result+1;
                }
                continue;
            }
            if(s[i]=='X'){
                if(s[i+1]=='L'){
                    result = result+40;
                    i++;
                } else if(s[i+1]=='C'){
                    result = result+90;
                    i++;
                } else{
                    result = result+10;
                }
                continue;
            }
            if(s[i]=='C'){
                if(s[i+1]=='D'){
                    result = result+400;
                    i++;
                } else if(s[i+1]=='M'){
                    result = result+900;
                    i++;
                } else{
                    result = result+100;
                }
                continue;
            }
            switch (s[i]){
                case 'I':
                    result = result+ 1;
                    break;
                case 'V': result = result+ 5;
                    break;
                case 'X': result = result+ 10;
                    break;
                case 'L': result = result+ 50;
                    break;
                case 'C': result = result+ 100;
                    break;
                case 'D': result = result+ 500;
                    break;
                case 'M': result = result+ 1000;
                    break;

            }
        }
        return result;
    }
};

int main() {
    string s = "MCMXCIV";
    Solution test;
    cout<<test.romanToInt(s);
    return 0;
}