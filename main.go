package main

import (
	"fmt"
	"net/http"
)


func main(){
	fmt.Println("Hello World!")
	fmt.Println(Hello("Beda", 24))

	mux := SetupRouter()
	fmt.Println("Servidor corriendo en el puerto 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}

}
