package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitLinkList(ints []int) *ListNode {
    if ints == nil {
        return nil
    }
    head := ListNode{ints[0], nil}
    pre := &head
    for i := 1; i < len(ints); i ++ {
        node := ListNode{ints[i], nil}
        pre.Next = &node
        pre = pre.Next
    }
    return &head
}

func LinkListToSlice(head *ListNode) []int {
    ans := []int{}
    for head.Next != nil {
        ans = append(ans, head.Val)
        head = head.Next
    }
    ans = append(ans, head.Val)
    return ans
}