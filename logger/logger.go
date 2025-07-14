package main

import (
	"trung_go/greetings"

	"fmt"
)

func main(){
	message:= greetings.Greetings("Daniel Le");
	fmt.Println(message);
}