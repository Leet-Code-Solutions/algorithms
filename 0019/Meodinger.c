/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode *removeNthFromEnd(struct ListNode *head, int n) {
    if (!head->next) return NULL;

    struct ListNode *node1 = head, *node2 = head;
    for (int i = 0; i <= n; i++) {
        node2 = node2->next;
        if (!node2 && i == n - 1) return head->next;
    }

    while (node2) {
        node1 = node1->next;
        node2 = node2->next;
    }
    node1->next = node1->next->next;

    return head;
}