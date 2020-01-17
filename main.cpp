#include <vector>
#include <math.h>
#include <iostream>
using namespace std;

class Solution {
public:
    string intToRoman(int num) {
        string result = "";
        while (num>0){
            if(num>=1000){
                result.push_back('M');
                num = num - 1000;
            } else if(num >= 900){
                result.push_back('C');
                result.push_back('M');
                num = num -900;
            } else if(num>=500){
                result.push_back('D');
                num = num-500;
            } else if(num>=400){
                result.push_back('C');
                result.push_back('D');
                num = num -400;
            } else if(num>=100){
                result.push_back('C');
                num = num-100;
            } else if(num>=90){
                result.push_back('X');
                result.push_back('C');
                num = num -90;
            } else if(num>= 50){
                result.push_back('L');
                num = num-50;
            } else if(num>= 40){
                result.push_back('X');
                result.push_back('L');
                num = num-40;
            } else if(num>= 10){
                result.push_back('X');
                num = num-10;
            } else if(num== 9){
                result.push_back('I');
                result.push_back('X');
                num = num-9;
            }else if(num>= 5){
                result.push_back('V');
                num = num-5;
            }else if(num== 4){
                result.push_back('I');
                result.push_back('V');
                num = num-4;
            } else if(num>0){
                result.push_back('I');
                num = num-1;
            }
        }
        return result;
    }
    int myAtoi(string str) {
        string result = "";
        if(str.size()==0){
            return 0;
        }
        int flag = 1;
        for(int i=0;i<str.size();i++) {
            if(str[i]==' ' && result.size()==0){
                continue;
            } else {
                if(result.size()==0 && str[i]=='0'){
                    result.push_back(str[i]);
                } else if((result[0]=='-'||result[0]=='+')&&str[i]=='0'&&result.size()==1){
                    result.push_back(str[i]);
                } else{
                    if ((str[i] - '0') < 10 && (str[i] - '0') >= 0) {
                        result.push_back(str[i]);
                    } else if (result.size() == 0 && (str[i] == '-' || str[i]=='+')) {
                        result.push_back(str[i]);
                    } else if (result.size() != 0) {
                        break;
                    } else if (result.size() == 0 && !(((str[i] - '0') < 10) && ((str[i] - '0') > 0))) {
                        return 0;
                    }
                }
            }
        }
        int s = 0;
        if((result[0]=='-' || result[0]=='+')&&result.size()==1){
            return 0;
        }
        if(result.size()==0){
            return 0;
        }
        try {
            s = stoi(result);
        }catch(exception e) {
            if(result[0]=='-'){
                s = INT_MIN;
            } else{
                s=INT_MAX;
            }
        }
        return s;
    }
};

int main() {
    string x = "-000000000000001";
    Solution test;
    cout<<test.myAtoi(x);
    return 0;
}