package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola desde go!")
	})


	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err)
	}
}
