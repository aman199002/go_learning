package main
import (
	"errors"
	"fmt"
)

func f1(num int) (int,error){
	if num == 42{
		return -1,errors.New("Can't work with 42")
	}
	return num+3,nil;
}

type argError struct{
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(num int) (int, error){
	if num == 42{
		return -1,&argError{num,"Can't work with it"}
	}
	return num+3,nil
}


func main(){
	for _,num := range []int{7,42} {
		r,e := f1(num)
		if e != nil{
			fmt.Println("f1 failed:",e)
		} else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _,num := range []int{7,42} {
		r,e := f2(num)
		if e != nil{
			fmt.Println("f2 failed:",e)
		} else {
			fmt.Println("f2 worked:",r)
		}
	}

	fmt.Println(f1(42))
	fmt.Println(f1(40))

	res,e := f2(42)
	fmt.Println("res",res)
	fmt.Println("error",e)
}

