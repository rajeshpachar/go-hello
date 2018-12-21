package main;

import (
"github.com/rajeshpachar/hellomod/hellotest"
"github.com/rajeshpachar/hellomod/child"
//"github.com/rajeshpachar/go-hello/privatehello"
)

func main(){
	hellotest.SayHello("client calling")
	hellotest.DepHello()
	child.HelloChild()
}
