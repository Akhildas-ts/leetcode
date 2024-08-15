func isThree(n int) bool {

	count := 0
	 for i:=1; i<=n;i++ {
 
 
		 for j:=1;j<=n;j++{
 
			 if j * i == n {
 
				 count ++
 
			 }
		 }
	 }
 
	 if 3 <= count {
		 return true 
	 }
 
	 return false
 
 
	 
 }