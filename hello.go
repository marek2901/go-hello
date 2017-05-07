package main

import "fmt"

func main() {
  h := NewHuman("Marek", 22)
  h.sayHello()

  fmt.Printf("\n")

  s := NewSuperHuman("I'm the best")
  s.sayHello()
}
