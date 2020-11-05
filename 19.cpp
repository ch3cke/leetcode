//
// Created by ch3ck on 2020/3/15.
//
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(head->next==NULL){
            return NULL;
        }else{
            ListNode* Remove = head->next;
            ListNode* First = head;
            ListNode* End = head;
            while(n>0&&End->next!=NULL){
                n = n-1;
                End = End->next;
            }
            if(n!=0){
                head = head->next;
            } else{
                while(End->next!=NULL){
                    if(Remove->next==NULL){
                        break;
                    }else{
                        First=First->next;
                        Remove = Remove->next;
                        End = End->next;
                    }
                }
                First->next=Remove->next;
            }
            return head;
        }
    }
};
//leetcode submit region end(Prohibit modification and deletion)

