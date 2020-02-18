package codinginterviews

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
//
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return root
}
