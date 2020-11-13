package tree

//题目描述
//输入一个链表，反转链表后，输出新链表的表头。
//示例1
//输入
//{1,2,3}
//返回值
//{3,2,1}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	if pHead.Next == nil {
		return pHead
	}
	a := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return a
}
