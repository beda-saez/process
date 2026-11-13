package main

import (
	"fmt"
	"net/http"
	"time"
)

name := "Beda"
age := 24

func main(){
	fmt.Println("Hello World!")
	result := Hello(name, age)
	time.Sleep(1000*time.Milisecond)
	fmt.Println(result)

	mux := SetupRouter()
	fmt.Println("Servidor corriendo en el puerto 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}

}
