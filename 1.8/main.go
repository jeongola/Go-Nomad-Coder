package main

import "fmt"

func main() {
	a := 2
	b := &a     //&는 메모리 주소값을 가리킴
	*b = 202020 //*는 메모리값을 기리킴
	fmt.Println(a)
}
