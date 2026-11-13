package main

import(
	"fmt"
	"net/http"
)

func Hello(name string, age int) string{
	message := fmt.Sprintf("%v! Es un gusto saber que tiene %d de edad",name, age)
	return message

}