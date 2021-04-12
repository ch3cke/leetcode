#include <iostream>

#pragma warning(disable:4996)
using namespace std;

int main()
{
    __int64 a[6] = { 3746099070, 550153460, 3774025685, 1548802262, 2652626477, 2230518816 };
    unsigned int a2[4] = { 2,2,3,4 };
    unsigned int v3, v4;
    int v5;
    for (int j = 0; j <= 4; j += 2) {
        v3 = a[j];
        v4 = a[j + 1];
        v5 = 1166789954*0x40;
        for (int i = 0; i <= 0x3F; ++i) {
            v4 -= (v3 + v5 + 20) ^ ((v3 << 6) + a2[2]) ^ ((v3 >> 9) + a2[3]) ^ 0x10;
            v3 -= (v4 + v5 + 11) ^ ((v4 << 6) + *a2) ^ ((v4 >> 9) + a2[1]) ^ 0x20;
            v5 -= 1166789954;
        }
        a[j] = v3;
        a[j + 1] = v4;
    }

    /*将整型数组作为字符输出，注意计算机小端排序*/
    for (int i = 0; i < 6; ++i) {
        cout << *((char*)&a[i] + 2) << *((char*)&a[i] + 1) <<  * ((char*)&a[i]);
    }

    system("PAUSE");
    return 0;
}