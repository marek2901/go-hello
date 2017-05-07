package main

import "fmt"

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
