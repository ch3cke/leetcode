//
// Created by ch3ck on 2020/1/1.
//

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int temp = 0;
        ListNode *tr = new ListNode((l1->val+l2->val)%10);
        temp = (l1->val+l2->val)/10;
        ListNode *q = l1->next;
        ListNode *p = l2->next;
        ListNode *result=tr;
        while(q!=NULL || p!=NULL){
            int i1,i2;
            if(q!=NULL){
                i1=q->val;
                q=q->next;
            } else{
                i1=0;
            }
            if(p!=NULL){
                i2 = p->val;
                p=p->next;
            } else{
                i2=0;
            }
            tr->next=new ListNode((i1+i2+temp)%10);
            tr = tr->next;
            temp = (i1+i2+temp)/10;
        }
        if(temp!=0){
            tr->next = new ListNode(temp);
        }
        return result;
    }
};