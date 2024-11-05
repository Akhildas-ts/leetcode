package main 




// question: 141 
// check its a cycle loop  if  its cycle loop return true its not cycle linked list return false 




type ListNode struct {

	Val int 
	Next *ListNode
}



func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
    return false 
 }


 slow,fast := head ,head.Next 


 for fast != nil && fast.Next != nil{


    if slow == fast {
        return true 
    }

    slow = slow.Next 
    fast = fast.Next.Next 
 }

 return false 


}