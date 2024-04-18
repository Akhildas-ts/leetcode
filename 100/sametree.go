package main 


// question 
// we got two trees with parameter , check the tree is equal or not ? 
// so we need to check node is equal and,left, right is equal , if its equal then true 


// Algoritham 
// we working with an recustion function 
//  both of the tree val, left, right ,must same then only true 

// p == nil && p == nil means its comes at end of the one part of the recurtion function 
// when these condtioon statisfied it means true 



type treeNode struct{
	Val int 
	left *treeNode
	right *treeNode
}


func isSameTree(p *treeNode,q *treeNode)bool{

	if p == nil && q == nil {
		return true

	}

	if p==  nil || q==nil{
		return false
	}

	if p.Val != q.Val{
		return false
	}


return isSameTree(p.left,q.left) && isSameTree(p.right,q.right)
	


}