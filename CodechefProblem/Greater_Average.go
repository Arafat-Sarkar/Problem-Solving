package codechefproblem
import "fmt"

func AvarageCalulate(a int , b int) float64 {
     result := float64 (a+b)/2
	 return result
}

func main(){
    
    var t int
    fmt.Scan(&t)
    
    for  i :=0; i<t ;i++ {
        var a, b, c int
	fmt.Scan(&a,&b,&c)

	avrge:= AvarageCalulate(a,b)
	if avrge > float64(c) {
		fmt.Println("YES")
	} else{
		fmt.Println("NO")
	}
	
    }
	

}
