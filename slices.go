package main
import(
	"fmt"
)

func main() {
	s := make([]string,3)
	fmt.Println("s: ",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s: ",s)

	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("s: ",s)

	c := make([]string,len(s))
	fmt.Println("c: ",c)
	copy(c,s)
	fmt.Println("c: ",c)

	l := s[2:5]
	fmt.Println("l: ",l)

	l = s[:4]
	fmt.Println("l: ",l)

	l = s[2:]
	fmt.Println("l: ",l)

	t := []string{"g","h","i"}
	fmt.Println("t: ",t)

	twoD := make([][]int,3)
	fmt.Println("twoD: ",twoD)
	for i:=0; i<3; i++ {
		innerLen := i+1
		twoD[i] = make([]int,innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("twoD: ",twoD)
}