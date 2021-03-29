#include <iostream>
#include <string.h>
#include <stack>
#include <stdio.h>
#include <stdlib.h>
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


void encode(int* v, int* k){
    int num = 0;
    int v6 = v[1];
    int v3 = v[0];
    for(int i=0;i<32;i++){
        num += 0x9e3779b9;
        v3 += (k[1] + (v6 >> 5)) ^ (num + v6) ^ (k[0] + 16 * v6);
        v6 += (k[3] + (v3 >> 5)) ^ (num + v3) ^ (k[2] + 16 * v3);
    }
    v[1] = v6;
    v[0] = v3;
}

void decode(int *v, int* k){
    int num = 0xc6ef3720;
    int y = v[0];
    int z = v[1];
    for(int i=0;i<32;i++){
        z -= ((y*16)+k[2])^(y+num)^((y>>5)+k[3]);
        y -= ((z*16)+k[0])^(z+num)^((z>>5)+k[1]);
        num -= 0x9e3779b9;
    }
    v[0] = y;
    v[1] = z;
}

int main(){
    int key[4] = {0x67616C66, 0x6B61667B, 0x6C665F65, 0x7D216761};
    int v[2] = {0x2A819E47, static_cast<int>(0x942CAFE7)};
//    encode(v, key);
    //printf("%x, %x\n",v[0],v[1]);
    decode(v, key);
    printf("%x, %x\n",v[0],v[1]);
    char buffer[4];
    char buffer2[4];
    int a[2];
    unsigned char buf[4];
    FILE * m_File = fopen("tea.png1.out", "rb");
    FILE * p_file = fopen("1.png", "wb");
    while (fread(buffer, sizeof(char ), 4, m_File)!=0&&fread(buffer2,sizeof(char ), 4, m_File)!=0){
        a[0] = buffer[3];
        a[0] <<=8;
        a[0] += buffer[2]+1;
        a[0] <<=8;
        a[0] += buffer[1];
        a[0] <<=8;
        a[0] += buffer[0];
        a[1] = buffer2[3];
        a[1] <<=8;
        a[1] += buffer2[2];
        a[1]<<=8;
        a[1] += buffer2[1]+1;
        a[1] <<=8;
        a[1] += buffer2[0];
        printf("%x, %x\n",a[0],a[1]);
        decode(a, key);
        fwrite(a, sizeof(int),2 ,p_file);
    }
    return 0;
}
