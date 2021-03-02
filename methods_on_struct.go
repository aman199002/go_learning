package main
import "fmt"

type rect struct{
	width, height int
}

func (r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int{
	return 2*r.width + 2*r.height
}

func main(){
	r := rect{width: 20, height: 10}
	fmt.Println("area:",r.area())
	fmt.Println("preimeter:",r.perim())

	rp := &r
	fmt.Println("area:",rp.area())
	fmt.Println("preimeter:",rp.perim())
}
