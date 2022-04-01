/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode *middleNode(struct ListNode *head) {
    struct ListNode *node1 = head, *node2 = head;

    while (node2 && node2->next) {
        node1 = node1->next;
        node2 = node2->next->next;
    }

    return node1;
}