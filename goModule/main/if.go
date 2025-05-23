package main

import "fmt"

func main(){
	if 7%2 ==0 {
		fmt.Printf("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if(8%4 ==0){
		fmt.Println("8 is divisible by 4")
	}

	if(8%2==0 || 7%2==0){
		fmt.Println("either 8 or 7 are even")
	}

	if num:=9; num%3==0{
		fmt.Println("9%3=0")
	}
	// num相当于局部变量
	// fmt.Println(num)

}