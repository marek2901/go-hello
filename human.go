package main

import "fmt"

type Human struct {
  name string
  age int
}

func (h Human) sayHello() {
  fmt.Printf("Hello my name is %s I'm human and I'm %d old \n",
    h.name, h.age)
}

func NewHuman(nam string, age int) *Human {
  return &Human{ nam, age, }
}
