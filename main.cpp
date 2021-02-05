#include<stdio.h>

struct Test{
    char Test[13];
    int a = 64;
};

int main(){
    Test s;
    s.Test[0] = '0';
    s.Test[2] = '1';
    s.Test[12] = 'c';
    printf("%c", s.a);
    return 0;
}
