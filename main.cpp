#include <vector>
#include <math.h>
#include <iostream>
#include <stack>
#include <stdio.h>

using namespace std;

struct ListNode {
       int val;
       ListNode *next;
       ListNode(int x) : val(x), next(NULL) {}
     };

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* Remove = head->next;
        ListNode* First = head;
        ListNode* End = head;
        while(n>0){
            n = n-1;
            End = End->next;
        }
        while(End->next!=NULL){
            First=First->next;
            Remove = Remove->next;
            End = End->next;
        }
        First->next=Remove->next;
        return Remove;
    }
};

int main() {
    int num = 10;
    ListNode* s;
    ListNode q;
    s->val=10;
    ListNode* success = s;
    while(num>0){
        num--;
        q->val=num;
        s->next = q;
        s = s->next;
    }

    Solution sou ;
    cout<<sou.removeNthFromEnd(success,2)->val;
}