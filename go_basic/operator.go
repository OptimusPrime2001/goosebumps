package main

import "fmt"

func testName(name *string) {
	*name = "Le Văn Trung"
	fmt.Println(*name)
}
func TestOperator(n int) {
	name := "Darwin"
	testName(&name)
	fmt.Println(name)

}
