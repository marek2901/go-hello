package main

import "fmt"

type Human struct {
  name string
  age int
}

type Humaner interface {
  sayHello()
}

func (h Human) sayHello() {
  fmt.Printf("Hello my name is %s I'm human and I'm %d old \n",
    h.name, h.age)
}

func NewHuman(nam string, age int) Humaner {
  return &Human{ nam, age }
}

type SuperHuman struct {
  Human
  message string
}

type SuperHumaner interface {
    sayHello()
}

func (s SuperHuman) sayHello() {
  s.Human.sayHello()
  fmt.Printf("WOHOA %s \n", s.message)
}

func NewSuperHuman(mes string) SuperHumaner{
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

