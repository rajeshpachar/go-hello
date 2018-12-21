package main;

import (
"github.com/rajeshpachar/hellomod"
"github.com/rajeshpachar/hellomod/child"
)

func main(){
	hellomod.SayHello("client calling")
	hellomod.DepHello()
	child.HelloChild()
}
