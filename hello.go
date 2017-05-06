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
  return &Human{ nam, age }
}

type SuperHuman struct {
  Human
  message string
}

func (s SuperHuman) sayHello() {
  s.Human.sayHello()
  fmt.Printf("WOHOA %s \n", s.message)
}

func NewSuperHuman(mes string) *SuperHuman {
  return &SuperHuman{
    Human{"Super", 123,}, mes,
  }
}

func main() {
  h := NewHuman("Marek", 21)
  h.sayHello()

  fmt.Printf("\n")

  s := NewSuperHuman("I'm the best")
  s.sayHello()
}
