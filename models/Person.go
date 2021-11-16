package models

import "time"

type Person struct {
	name string
	age  int
	addr string
	dob  time.Time
}

func NewPerson() Person {
	return Person{}
}

func (p Person) WithName(name string) Person {
	p.name = name
	return p
}

func (p Person) WithAge(age int) Person {
	p.age = age
	return p
}

func (p Person) WithAddr(addr string) Person {
	p.addr = addr
	return p
}

func (p Person) WithDob(dob time.Time) Person {
	p.dob = dob
	return p
}
