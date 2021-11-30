package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	p1 := person{"Sylvio", "Campos", 47, 173}
	fmt.Println(p1)

	s1 := student{p1, "Estágio de treinador", "SC Corinthians Paulista"}
	fmt.Println(s1)
}
