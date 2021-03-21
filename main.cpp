#include <iostream>
#include <string>
#include <stack>

using namespace std;


int calculate(string s){
    stack<int> ops;
    ops.push(1);
    int sign=1;

    int ret = 0;
    int n = s.length();
    int i=0;
    while (i<n){
        if(s[i]==' '){
            i++;
        } else if (s[i]=='+'){
            sign = ops.top();
            i++;
        } else if (s[i]=='-'){
            sign = -ops.top();
            i++;
        } else if (s[i]=='('){
            ops.push(sign);
            i++;
        } else if (s[i]==')'){
            ops.pop();
            i++;
        } else{
            long num = 0;
            while (i<n&&s[i]>='0'&&s[i]<='9'){
                num=num*10+s[i]-'0';
                i++;
            }
            ret += sign*num;
        }
    }
    return ret;
}



int main(){
    uint8_t a, b,c;
    c = 39;
    a = c*32;
    printf("%d\n",a);
    return 0;
}
